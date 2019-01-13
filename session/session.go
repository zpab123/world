// /////////////////////////////////////////////////////////////////////////////
// 工具包

package session

import (
	"net"

	"github.com/zpab123/world/network/socket" // socket 库
)

// 创建1个可以收发 packet 的 socket
func CreatePacketSocket(conn net.Conn) *socket.PacketSocket {
	// 创建1个基础 socket
	st := socket.Socket{
		Conn: conn,
	}

	// 创建1个带 读写 buffer 的 socket
	buffSocket := socket.NewBufferSocket(st)

	// 创建 pacetSocket
	pacetSocket := socket.NewPacketSocket(buffSocket)

	return pacetSocket
}
