// /////////////////////////////////////////////////////////////////////////////
// 能够读写 packet 数据的 socket

package socket

import (
	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
)

// /////////////////////////////////////////////////////////////////////////////
// PacketSocket 对象

// PacketSocket
type PacketSocket struct {
	socket model.ISocket // 接口继承： 符合 ISocket 的对象
}

// 创建1个新的 PacketSocket 对象
func NewPacketSocket(st model.ISocket) *PacketSocket {
	// 创建对象
	pktSocket := &PacketSocket{
		socket: st,
	}

	return pktSocket
}

// 接收下1个 packet 数据
//
// 返回, nil=没收到完整的 packet 数据; packet=完整的 packet 数据包
