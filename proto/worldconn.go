// /////////////////////////////////////////////////////////////////////////////
// 对 PacketSocket 的封装，定义一些 world 内部常用的消息

package proto

import (
	"github.com/zpab123/syncutil"      // 原子变量库
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/utils"   // 工具库
	"github.com/zpab123/zplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// WorldConnection 对象

// world 框架内部需要用到的一些常用网络消息
type WorldConnection struct {
	state        syncutil.AtomicUint32 // conn 状态
	opts         *model.TWorldConnOpts // 配置参数
	packetSocket *network.PacketSocket // 接口继承： 符合 IPacketSocket 的对象
}

// 新建1个 WorldConnection 对象
func NewWorldConnection(socket model.ISocket, opt *model.TWorldConnOpts) *WorldConnection {
	// 创建 packetSocket
	bufSocket := network.NewBufferSocket(socket)
	pktSocket := network.NewPacketSocket(bufSocket)

	// 创建对象
	wc := &WorldConnection{
		packetSocket: pktSocket,
	}

	// 设置为初始化状态
	wc.state.Store(model.C_WCONN_STATE_INIT)

	return wc
}

// 接收1个 msg 消息
func (this *WorldConnection) RecvPacket() {
	// 接收 packet
	pkt, err := this.packetSocket.RecvPacket()
	if nil != err {
		return nil, err
	}

	// 处理 packet
	this.HandlePacket(pkt)

}

// 发送1个 msg 消息
func (this *WorldConnection) SendMsg() {

}

// 回应握手消息
func (this *WorldConnection) HandlePacket(pkt *network.Packet) {
	// 根据类型处理数据
	switch pkt.GetId() {
	case model.C_PACKET_ID_INVALID: // 无效类型
		break
	case model.C_PACKET_ID_HANDSHAKE: // 客户端握手请求
		break
	case model.C_PACKET_ID_HANDSHAKE_ACK: // 客户端握手 ACK
		break
	case model.C_PACKET_ID_HEARTBEAT: // 心跳数据
		break
	default:
		break
	}
}

//  处理握手消息
func (this *WorldConnection) HandleHandshake(pkt *network.Packet) {
	// 状态效验
	if this.state.Load() != model.C_WCONN_STATE_INIT {
		return
	}

	// 解码消息

	// 处理消息

	// 回复处理结果
}

//  处理握手ACK
func (this *WorldConnection) HandleHandshakeAck() {
	// 状态效验
	if this.state.Load() != model.C_WCONN_STATE_WAIT_ACK {
		return
	}

	// 改变为工作状态
	this.state.Store(model.C_WCONN_STATE_WORKING)

	// 发送心跳数据
}

//  处理心跳消息
func (this *WorldConnection) HandleHeartbeat() {
	// 状态效验
	if this.state.Load() != model.C_WCONN_STATE_WORKING {
		return
	}

	// 发送心跳数据
}

// 游戏常用内部消息
