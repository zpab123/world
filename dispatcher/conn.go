// /////////////////////////////////////////////////////////////////////////////
// dispatcher 连接对象

package dispatcher

import (
	"net"

	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
)

// dispatcher 连接对象
type DispatcherConn struct {
	worldConn *network.WorldConnection // world 引擎连接对象
}

// 新建1个 DispatcherConn
func NewDispatcherConn(netconn net.Conn, opt *network.TWorldConnOpts) *DispatcherConn {
	// 创建组件
	socket := &network.Socket{
		Conn: netconn,
	}
	wc := network.NewWorldConnection(socket, opt)

	// 创建 DispatcherConn
	cp := &DispatcherConn{
		worldConn: ws,
	}

	return cp
}

// 启动 DispatcherConn
func (this *DispatcherConn) Run() {
	// 接收线程

	// 发送线程
}
