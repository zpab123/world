// /////////////////////////////////////////////////////////////////////////////
// dispatcher 客户端代理

package dispatcher

import (
	"net"

	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
)

// 分发客户端
type ClientProxy struct {
	worldConn *network.WorldConnection // world 引擎连接对象
}

// 新建1个 ClientProxy
func NewClientProxy(netconn net.Conn) *ClientProxy {
	// 创建 socket
	socket := &network.Socket{
		Conn: netconn,
	}

	// 创建 WorldConnection
	wc := network.NewWorldConnection(socket)
}
