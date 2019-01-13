// /////////////////////////////////////////////////////////////////////////////
// socket 创建管理

package connector

import (
	"net"

	"github.com/zpab123/world/network/socket" // socket 库
)

// /////////////////////////////////////////////////////////////////////////////
// SocketCreator 对象

// socket 创建管理
type SocketCreator struct {
	pacetSocket *socket.PacketSocket // 可以收发 packet 的 socket
}

// 创建1个可以收发 packet 的 socket
func (this *SocketCreator) CreatePacketSocket(conn net.Conn) {
	// 创建1个基础 socket
	st := socket.Socket{
		Conn: conn,
	}

	// 创建1个带 读写 buffer 的 socket
	buffSocket := socket.NewBufferSocket(st)

	// 创建 pacetSocket
	this.pacetSocket = socket.NewPacketSocket(buffSocket)
}

// 获取 PacketSocket
func (this *SocketCreator) GetPacketSocket() *socket.PacketSocket {
	return this.pacetSocket
}
