// /////////////////////////////////////////////////////////////////////////////
// 面向客户端的 session 组件

package session

import (
	"net"

	"github.com/zpab123/world/model"          // 全局模型
	"github.com/zpab123/world/network/packet" // packet 消息包
	"github.com/zpab123/world/network/socket" // socket
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// ClientSession 对象

// 面向客户端的 session 对象
type ClientSession struct {
	opts         *model.TSessionOpts  // session 配置参数
	packetSocket *socket.PacketSocket // 对象继承： 继承至 PacketSocket 对象
	sesssionMgr  model.ISessionManage // sessiong 管理对象
	// session_id
	// 用户id
	pktHandler model.ICilentPktHandler // 客户端 packet 消息处理器
}

func NewClientSession(opt *model.TSessionOpts) *ClientSession {
	// 创建 pktSocket
	pktSocket := socket.NewPacketSocket(opt.Socket)

	// 创建对象
	cs := &ClientSession{
		opts:         opt,
		packetSocket: pktSocket,
		sesssionMgr:  mgr,
		pktHandler:   handler,
	}

	return cs
}

// run
func (this *ClientSession) Run() {
	// 结束线程

	// 开启接收线程
	go this.recvLoop()

	// 开启发送线程
	go this.sendLoop()
}

// 接收线程
func (this *ClientSession) recvLoop() {
	for {
		// 接收消息
		var pkt *packet.Packet
		pkt = this.packetSocket.RecvPacket()
		if nil == pkt {
			continue
		}

		// 处理消息
		this.handlePacket(pkt)
	}
}

// 发送线程
func (this *ClientSession) sendLoop() {
	for {
		this.packetSocket.Flush()
	}
}

// 处理 packet
func (this *ClientSession) handlePacket(pkt *packet.Packet) {
	// 获取类型
	pktType := pkt.GetType()

	// 根据类型处理数据
	switch pktType {
	case model.C_PACKET_TYPE_HANDSHAKE: // 握手消息
		break
	case model.C_PACKET_TYPE_HEARTBEAT: // 心跳消息
		break
	default:
		break
	}
}
