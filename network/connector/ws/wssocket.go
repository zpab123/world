// /////////////////////////////////////////////////////////////////////////////
// websocket 客户端 <-> 服务器 通信

package ws

import (
	"net"

	"github.com/zpab123/world/ifs" // 全局接口库
)

// /////////////////////////////////////////////////////////////////////////////
// wsSocket 对象

// websocket 管理
type wsSocket struct {
	tcpConn   net.Conn       // Socket原始连接
	connector ifs.IConnector // connector 组件
}

// 创建1个新的 wsSocket 对象
func newWsSocket(conn net.Conn, cntor ifs.IConnector) ifs.ISocket {
	// 创建 socket
	socket := &wsSocket{
		tcpConn:   conn,
		connector: cntor,
	}

	return socket
}

// 刷新缓冲区
func (this *wsSocket) Flush() error {
	return nil
}
