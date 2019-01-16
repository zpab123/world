// /////////////////////////////////////////////////////////////////////////////
// 工具包

package connector

import (
	"net"

	"github.com/zpab123/world/model"          // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/world/network/socket" // socket 库
)

// 创建1个可以收发 packet 的 socket
func CreatePacketSocket(conn net.Conn, contor model.IConnector) *socket.PacketSocket {
	// 创建1个带 读写 buffer 的 socket
	buffSocket := CreateBufferSocket(conn)

	// 创建 pacetSocket
	pacetSocket := socket.NewPacketSocket(buffSocket, contor)

	return pacetSocket
}

// 创建1个自定义 buffer 的 socket
func CreateBufferSocket(conn net.Conn) *socket.BufferSocket {
	// 创建1个基础 socket
	st := socket.Socket{
		Conn: conn,
	}

	// 创建1个带 读写 buffer 的 socket
	buffSocket := socket.NewBufferSocket(st)

	return buffSocket
}
