// /////////////////////////////////////////////////////////////////////////////
// 消息库

package session

import (
	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// FrontendMsg

// FrontendSession 消息
type FrontendMsg struct {
	session *FrontendSession // session 对象
	packet  *network.Packet  // packet 数据包
}

// 创建1个 FrontendMsg
func NewFrontendMsg(ses *FrontendSession, pkt *network.Packet) *FrontendMsg {
	msg := &FrontendMsg{
		session: ses,
		packet:  pkt,
	}

	return msg
}

// 获取 session 对象
func (this *FrontendMsg) GetSession() *FrontendSession {
	return this.session
}

// 获取 packet 对象
func (this *FrontendMsg) GetPacket() *network.Packet {
	return this.packet
}

// /////////////////////////////////////////////////////////////////////////////
// BackendSesMsg

// BackendSession 消息
type BackendMsg struct {
	session *BackendSession // session 对象
	packet  *network.Packet // packet 数据包
}

// 创建1个 FrontendMsg
func NewBackendMsg(ses *BackendSession, pkt *network.Packet) *BackendMsg {
	msg := &BackendMsg{
		session: ses,
		packet:  pkt,
	}

	return msg
}

// 获取 session 对象
func (this *BackendMsg) GetSession() *BackendSession {
	return this.session
}

// 获取 packet 对象
func (this *BackendMsg) GetPacket() *network.Packet {
	return this.packet
}
