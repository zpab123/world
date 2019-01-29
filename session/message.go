// /////////////////////////////////////////////////////////////////////////////
// Session 消息

package session

import (
	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// Message

// session 消息对象
type Message struct {
	session *ClientSession  // session 对象
	packet  *network.Packet // 数据包
}

// 获取 session
func (this *Message) GetSession() *ClientSession {
	return this.session
}

// 获取 packet
func (this *Message) GetPacket() *network.Packet {
	return this.packet
}
