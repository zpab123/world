// /////////////////////////////////////////////////////////////////////////////
// tcp 接口汇总

package tcp

import (
	"net"
)

// socket 基础参数接口
type iSocketOpt interface {
	GetMaxPacketSize() int
	SetReadTimeout(conn net.Conn, callback func())
	SetWriteTimeout(conn net.Conn, callback func())
}
