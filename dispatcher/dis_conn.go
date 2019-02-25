// /////////////////////////////////////////////////////////////////////////////
// DispatcherClient 连接参数

package dispatcher

import (
	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// DispatcherConn

// 连接对象
type DispatcherConn struct {
	addr        *network.TLaddr       // 连接地址集合
	option      *TDispatcherClientOpt // 配置参数
	worldSocket *network.WorldSocket  // world 框架内部通信 socket 客户端
}

// 新建1个 DispatcherConn
func NewDispatcherConn(addr *network.TLaddr, opt *TDispatcherClientOpt) *DispatcherConn {
	// 创建对象
	ws := network.NewWorldSocket(addr, opt.WorldSocketOpt)

	// 创建 DispatcherConn
	dc := &DispatcherConn{
		addr:        addr,
		option:      opt,
		worldSocket: ws,
	}

	return dc
}

// 连接服务器
func (this *DispatcherConn) Connect() {
	this.worldSocket.Connect()

	go this.recvLoop()

	go this.sendLoop()
}

// 接收线程
func (this *DispatcherConn) recvLoop() {
	for {
		// 心跳检查
		this.worldSocket.CheckRemoteHeartbeat()

		// 接收消息
		pkt, _ := this.worldSocket.RecvPacket()
		if nil == pkt {
			continue
		}

		// 消息处理

	}
}

// 发送线程
func (this *DispatcherConn) sendLoop() {
	for {
		// 心跳检查
		this.worldSocket.CheckLocalHeartbeat()

		// 刷新缓冲区
		this.worldSocket.Flush()
	}
}
