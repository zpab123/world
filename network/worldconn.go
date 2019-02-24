// /////////////////////////////////////////////////////////////////////////////
// 对 PacketSocket 的封装，面向前端的连接对象

package network

import (
	"time"

	"github.com/gogo/protobuf/proto" // protobuf 库
	"github.com/pkg/errors"          // 错误库
	"github.com/zpab123/world/msg"   // world 内部通信消息
	"github.com/zpab123/world/state" // 状态管理
	"github.com/zpab123/zaplog"      // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// WorldConnection 对象

// world 框架内部需要用到的一些常用网络消息
type WorldConnection struct {
	stateMgr      *state.StateManager // 状态管理
	heartbeat     int64               // 心跳间隔，单位：秒。0=不设置心跳
	option        *TWorldConnOpt      // 配置参数
	packetSocket  *PacketSocket       // 接口继承： 符合 IPacketSocket 的对象
	timeOut       int64               // 心跳超时时间，单位：秒
	clientTimeOut int64               // 客户端心跳超时时间点，精确到秒
	serverTimeOut int64               // 服务器心跳超时时间点，精确到秒
}

// 新建1个 WorldConnection 对象
func NewWorldConnection(socket ISocket, opt *TWorldConnOpt) *WorldConnection {
	// 参数效验
	if nil == opt {
		opt = NewTWorldConnOpt()
	}

	// 创建状态管理
	st := state.NewStateManager()

	// 创建 packetSocket
	bufSocket := NewBufferSocket(socket, opt.BuffSocketOpts)
	pktSocket := NewPacketSocket(bufSocket)

	// 创建对象
	wc := &WorldConnection{
		stateMgr:     st,
		packetSocket: pktSocket,
		heartbeat:    opt.Heartbeat,
		option:       opt,
		timeOut:      opt.Heartbeat * 2,
	}

	// 设置为初始化状态
	wc.stateMgr.SetState(C_WCONN_STATE_INIT)

	return wc
}

// 接收1个 Packet 消息
func (this *WorldConnection) RecvPacket() (*Packet, error) {
	// 接收 packet
	pkt, err := this.packetSocket.RecvPacket()
	if nil == pkt || nil != err {
		return nil, err
	}

	// 记录时间
	if this.timeOut > 0 {
		this.clientTimeOut = time.Now().Unix() + this.timeOut
	}

	// 内部 packet
	if pkt.pktId < C_PACKET_ID_WORLD {
		this.handlePacket(pkt)

		return nil, err
	}

	// 状态效验
	s := this.stateMgr.GetState()
	if s != C_WCONN_STATE_WORKING {
		this.Close()

		err = errors.Errorf("WorldConnection 状态错误。正确状态=%d，当前状态=%d", C_WCONN_STATE_WORKING, s)

		return nil, err
	}

	return pkt, nil
}

// 发送1个 packet 消息
func (this *WorldConnection) SendPacket(pkt *Packet) error {
	// 状态效验

	// 记录超时
	if this.timeOut > 0 {
		this.serverTimeOut = time.Now().Unix() + this.timeOut
	}

	return this.packetSocket.SendPacket(pkt)
}

// 发送1个 packet 消息，然后将 packet 放回对象池
func (this *WorldConnection) SendPacketRelease(pkt *Packet) error {
	err := this.SendPacket(pkt)
	pkt.Release()

	return err
}

// 刷新缓冲区
func (this *WorldConnection) Flush() error {
	return this.packetSocket.Flush()
}

// 关闭 WorldConnection
func (this *WorldConnection) Close() error {
	var err error
	s := this.stateMgr.GetState()

	if s == C_WCONN_STATE_CLOSED {
		err = errors.New("WorldConnection 关闭失败：它已经处于关闭状态")

		return err
	}

	err = this.packetSocket.Close()

	this.stateMgr.SetState(C_WCONN_STATE_CLOSED)

	return err
}

// 检查客户端心跳
func (this *WorldConnection) CheckClientHeartbeat() {
	if this.timeOut > 0 {
		if time.Now().Unix() >= this.clientTimeOut {
			zaplog.Warnf("客户端心跳超时，断开连接")

			this.Close()
		}
	}
}

// 检查服务器心跳
func (this *WorldConnection) CheckServerHeartbeat() {
	if this.timeOut > 0 {
		if time.Now().Unix() >= this.serverTimeOut {

			this.sendHeartbeat()
		}
	}
}

// 处理 Packet 消息
func (this *WorldConnection) handlePacket(pkt *Packet) {
	// 根据类型处理数据
	switch pkt.pktId {
	case C_PACKET_ID_INVALID: // 无效类型
		zaplog.Error("WorldConnection 收到无效消息类型，关闭 WorldConnection")

		this.Close()

		break
	case C_PACKET_ID_HANDSHAKE: // 客户端握手请求
		this.handleHandshake(pkt.GetBody())

		break
	case C_PACKET_ID_HANDSHAKE_ACK: // 客户端握手 ACK
		this.handleHandshakeAck()

		break
	case C_PACKET_ID_HEARTBEAT: // 心跳数据
		zaplog.Debugf("收到 client 心跳消息")
		break
	default:
		break
	}
}

//  处理握手消息
func (this *WorldConnection) handleHandshake(data []byte) {
	// 状态效验
	if this.stateMgr.GetState() != C_WCONN_STATE_INIT {
		return
	}

	// 解码消息
	shakeInfo := &msg.HandshakeReq{}
	var err error

	err = proto.Unmarshal(data, shakeInfo)
	if nil != err {
		zaplog.Error("protobuf 解码握手消息出错。关闭连接")

		this.Close()

		return
	}

	zaplog.Debugf("收到握手消息。 key=%s", shakeInfo.Key)

	// 回复消息
	res := &msg.HandshakeRes{
		Code: msg.OK,
	}
	var buf []byte
	var sucess bool = true

	// 版本验证
	if this.option.ShakeKey != "" && shakeInfo.Key != this.option.ShakeKey {
		res.Code = msg.SHAKE_KEY_ERROR
		sucess = false
	}

	// 通信方式验证,后续添加

	// 回复处理结果
	buf, err = proto.Marshal(res)
	if nil != err {
		zaplog.Error("protobuf 编码握手消息出错。返回握手消息失败")
	} else {
		this.handshakeResponse(sucess, buf)
	}

	// 握手失败，关闭连接
	if sucess == false {
		this.Close()
	}
}

//  返回握手消息
func (this *WorldConnection) handshakeResponse(sucess bool, data []byte) {
	// 状态效验
	if this.stateMgr.GetState() != C_WCONN_STATE_INIT {
		return
	}

	// 返回数据
	pkt := NewPacket(C_PACKET_ID_HANDSHAKE)
	pkt.AppendBytes(data)
	this.SendPacketRelease(pkt)

	// 状态： 等待握手 ack
	if sucess {
		this.stateMgr.SetState(C_WCONN_STATE_WAIT_ACK)
	}
}

//  处理握手ACK
func (this *WorldConnection) handleHandshakeAck() {
	// 改变为工作状态
	if this.stateMgr.SwapState(C_WCONN_STATE_WAIT_ACK, C_WCONN_STATE_WORKING) {

		return
	}

	// 发送心跳数据
	this.sendHeartbeat()
}

//  发送心跳数据
func (this *WorldConnection) sendHeartbeat() {
	// 状态效验
	if this.stateMgr.GetState() != C_WCONN_STATE_WORKING {

		return
	}

	// 发送心跳数据
	pkt := NewPacket(C_PACKET_ID_HEARTBEAT)
	this.SendPacketRelease(pkt)
}
