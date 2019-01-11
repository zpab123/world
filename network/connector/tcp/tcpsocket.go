// /////////////////////////////////////////////////////////////////////////////
// tcp 客户端 <-> 服务器 通信

package tcp

import (
	"net"

	"github.com/zpab123/world/ifs" // 全局接口库
)

// /////////////////////////////////////////////////////////////////////////////
// tcpSocket 对象
type tcpSocket struct {
	tcpConn   net.Conn       // Socket原始连接
	connector ifs.IConnector // connector 组件
}

// 创建1个新的 tcpSocket 对象
func newtcpSocket(conn net.Conn, cntor ifs.IConnector) ifs.ISocket {
	// 创建 socket
	socket := &tcpSocket{
		tcpConn:   conn,
		connector: cntor,
	}

	return socket
}

// 刷新缓冲区
func (this *tcpSocket) Flush() error {
	return nil
}
