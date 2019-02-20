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
}
