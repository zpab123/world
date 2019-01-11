// /////////////////////////////////////////////////////////////////////////////
// tcp 客户端 <-> 服务器 通信

package tcp

import (
	"net"

	"github.com/zpab123/world/network/connector" // 连接器
)

// /////////////////////////////////////////////////////////////////////////////
// tcpSocket 对象
type tcpSocket struct {
	closeNotify func() // tcpSession 关闭成功后的回调函数
}

// 创建1个新的 tcpSocket 对象
func newtcpSocket(conn net.Conn, aptor connector.IAcceptor) {

}
