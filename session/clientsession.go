// /////////////////////////////////////////////////////////////////////////////
// 面向客户端的 session 组件

package session

import (
	"net"

	"github.com/zpab123/world/model"          // 全局模型
	"github.com/zpab123/world/network/packet" // packet 消息包
	"github.com/zpab123/world/network/socket" // socket
)

type ClientSession struct {
	packetSocket *socket.PacketSocket // 对象继承： 继承至 PacketSocket 对象
	connector    model.IConnector     // connector 组件
	// session_id
	// 用户id
	pktHandler model.ICilentPktHandler // 客户端 packet 消息处理器
}

func NewClientSession(st model.ISocket, cntor model.IConnector, handler model.ICilentPktHandler) *ClientSession {
	// 创建 pktSocket
	pktSocket := socket.NewPacketSocket(st, cntor)

	// 创建对象
	cs := &ClientSession{
		packetSocket: pktSocket,
		connector:    cntor,
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
		this.pktHandler.OnClientPkt(this, pkt)
	}
}

// 发送线程
func (this *ClientSession) sendLoop() {
	for {
		this.packetSocket.Flush()
	}
}
