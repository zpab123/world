// /////////////////////////////////////////////////////////////////////////////
// tcp 客户端 <-> 服务器 通信。能够接收和发送 packet 数据

package tcp

import (
	"net"

	"github.com/zpab123/world/model" // 全局接口库
)

// /////////////////////////////////////////////////////////////////////////////
// tcpSocket 对象
type tcpSocket struct {
	tcpConn   net.Conn         // Socket原始连接
	connector model.IConnector // connector 组件
}

// 创建1个新的 tcpSocket 对象
func newtcpSocket(conn net.Conn, cntor model.IConnector) interface{} {
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
