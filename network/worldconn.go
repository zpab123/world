// /////////////////////////////////////////////////////////////////////////////
// 对 PacketSocket 的封装，定义一些 world 内部常用的消息

package network

import (
	"time"

	"github.com/gogo/protobuf/proto"  // protobuf 库
	"github.com/zpab123/syncutil"     // 原子变量库
	"github.com/zpab123/world/config" // 配置文件读取
	"github.com/zpab123/world/model"  // 全局模型
	"github.com/zpab123/world/msg"    // world 内部通信消息
	"github.com/zpab123/world/utils"  // 工具库
	"github.com/zpab123/zplog"        // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// WorldConnection 对象

// world 框架内部需要用到的一些常用网络消息
type WorldConnection struct {
	state        syncutil.AtomicUint32 // conn 状态
	opts         *model.TWorldConnOpts // 配置参数
	packetSocket *PacketSocket         // 接口继承： 符合 IPacketSocket 的对象
	lastSendTime syncutil.AtomicInt64  // 上次给客户端发送消息的时间：单位秒
	lastRecvTime time.Time             // 上次接收到客户端消息的时间
	timeOut      time.Duration         // 心跳超时时间
}

// 新建1个 WorldConnection 对象
func NewWorldConnection(socket model.ISocket, opt *model.TWorldConnOpts) *WorldConnection {
	// 创建 packetSocket
	bufSocket := network.NewBufferSocket(socket)
	pktSocket := network.NewPacketSocket(bufSocket)

	// 创建参数
	if nil != opt {
		opt = model.NewTWorldConnOpts()
	}

	// 创建对象
	wc := &WorldConnection{
		packetSocket: pktSocket,
		opts:         opt,
		timeOut:      opt.Heartbeat * 2,
	}

	// 设置为初始化状态
	wc.state.Store(model.C_WCONN_STATE_INIT)

	return wc
}

// 接收1个 msg 消息
func (this *WorldConnection) RecvPacket() (*Packet, error) {
	// 接收 packet
	pkt, err := this.packetSocket.RecvPacket()
	if nil != err {
		return nil, err
	}

	// 处理 packet
	this.lastRecvTime = time.Now()
	pkt = this.handlePacket(pkt)

	return pkt, nil
}

// 发送1个 packet 消息
func (this *WorldConnection) SendPacket(pkt *Packet) error {
	this.lastSendTime.Store(time.Now().Unix())

	return this.packetSocket.SendPacket(pkt)
}

// 发送1个 packet 消息，然后将 packet 放回对象池
func (this *WorldConnection) SendPacketRelease(pkt *Packet) error {
	err := this.SendPacket(pkt)
	pkt.Release()

	return err
}

// 刷新缓冲区
func (this *WorldConnection) Flush() {
	this.packetSocket.Flush()
}

// 关闭 WorldConnection
func (this *WorldConnection) Close() {
	this.packetSocket.Close()
}

// 检查客户端心跳
func (this *WorldConnection) CheckClientHeartbeat() {
	if this.timeOut > 0 {
		outTime := this.lastRecvTime.Add(this.timeOut)
		if this.lastRecvTime.After(outTime) {
			zplog.Warnf("客户端心跳超时，断开连接")
			this.Close()
		}
	}
}

// 检查服务器心跳
func (this *WorldConnection) CheckServerHeartbeat() {
	if this.timeOut > 0 {
		passTime := time.Now().Unix() - this.lastSendTime.Load()
		if passTime >= this.timeOut {
			this.sendHeartbeat()
		}
	}
}

// 回应握手消息
func (this *WorldConnection) handlePacket(pkt *Packet) *Packet {
	// 根据类型处理数据
	switch pkt.GetId() {
	case model.C_PACKET_ID_INVALID: // 无效类型
		zplog.Error("WorldConnection 收到无效消息类型，关闭 WorldConnection ")
		this.Close()

		return nil
	case model.C_PACKET_ID_HANDSHAKE: // 客户端握手请求
		this.handleHandshake(pkt.GetBody())

		return nil
	case model.C_PACKET_ID_HANDSHAKE_ACK: // 客户端握手 ACK
		this.handleHandshakeAck()

		return nil
	case model.C_PACKET_ID_HEARTBEAT: // 心跳数据

		return nil
	default:
		return nil
	}
}

//  处理握手消息
func (this *WorldConnection) handleHandshake(body []byte) {
	// 状态效验
	if this.state.Load() != model.C_WCONN_STATE_INIT {
		return
	}

	// 解码消息
	shakeInfo := &msg.HandshakeReq{}
	err := proto.Unmarshal(body, shakeInfo)
	if nil != err {
		zplog.Error("protobuf 解码握手消息出错")
	}

	// 回复消息
	res := &msg.HandshakeRes{}

	// 版本验证
	if shakeInfo.Key != config.GetWorldIni().Key {
		res.Code = msg.SHAKE_KEY_ERROR
		body := proto.Marshal(res)
		this.handshakeResponse(false, body)
		this.Close()

		return
	}

	// 通信方式验证
	if shakeInfo.Acceptor != config.GetWorldIni().Acceptor {
		res.Code = msg.SHAKE_ACCEPTOR_ERROR
		body := proto.Marshal(res)
		this.handshakeResponse(false, body)
		this.Close()

		return
	}

	// 回复处理结果

}

//  返回握手消息
func (this *WorldConnection) handshakeResponse(sucess bool, body []byte) {
	// 状态效验
	if this.state.Load() != model.C_WCONN_STATE_INIT {
		return
	}

	// 返回数据
	pkt := NewPacket(model.C_PACKET_ID_HANDSHAKE)
	pkt.AppendBytes(body)
	this.SendPacketRelease(pkt)

	// 改变状态
	if sucess {
		this.state.Store(model.C_WCONN_STATE_WAIT_ACK)
	}
}

//  处理握手ACK
func (this *WorldConnection) handleHandshakeAck() {
	// 状态效验
	if this.state.Load() != model.C_WCONN_STATE_WAIT_ACK {
		return
	}

	// 改变为工作状态
	this.state.Store(model.C_WCONN_STATE_WORKING)

	// 发送心跳数据
	this.sendHeartbeat()
}

//  发送心跳数据
func (this *WorldConnection) sendHeartbeat() {
	// 状态效验
	if this.state.Load() != model.C_WCONN_STATE_WORKING {
		return
	}

	// 发送心跳数据
	pkt := NewPacket(model.C_PACKET_ID_HEARTBEAT)
	this.SendPacketRelease(pkt)
}
