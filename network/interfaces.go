// /////////////////////////////////////////////////////////////////////////////
// network 包 接口汇总

package network

import (
	"net"
)

// /////////////////////////////////////////////////////////////////////////////
// 公用接口

// TcpServer 服务
type ITcpService interface {
	OnTcpConn(conn net.Conn) // 收到1个新的 Tcp 连接对象
}

// WsServer 消息服务
type IWsService interface {
	OnWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
}
