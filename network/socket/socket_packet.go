// /////////////////////////////////////////////////////////////////////////////
// 能够读写 packet 数据的 socket

package socket

import (
	"bufio"

	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
)

// /////////////////////////////////////////////////////////////////////////////
// PacketSocket 对象

// PacketSocket
type PacketSocket struct {
	model.ISocket // 接口继承： 符合 ISocket 的对象
}

// 创建1个新的 PacketSocket 对象
func NewPacketSocket(socket model.ISocket) *PacketSocket {
	// 创建对象
	bufsocket := &PacketSocket{
		ISocket: socket,
	}

	return bufsocket
}
