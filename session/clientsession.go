// /////////////////////////////////////////////////////////////////////////////
// 面向客户端的 session 组件

package session

import (
	"net"

	"github.com/zpab123/world/model" // 全局模型
)

type ClientSession struct {
	// socket
	// session_id
	// 用户id
	// 发送队列
	pktHandler model.ICilentPktHandler // 客户端 packet 消息处理器
}

func NewClientSession(handler model.ICilentPktHandler) *ClientSession {
	// 创建对象
	cs := &ClientSession{
		pktHandler: handler,
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

		var pkt interface{}

		// 处理消息
		this.pktHandler.OnClientPkt(this, pkt)
	}
}

// 发送线程
func (this *ClientSession) sendLoop() {

}
