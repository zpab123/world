// /////////////////////////////////////////////////////////////////////////////
// Application 代理

package app

import (
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session库
)

// app 服务
type Appdelegate struct {
	clientPacketQueue chan *session.Message
}

// 创建1个新的 AppProxy
func NewAppdelegate() *Appdelegate {
	s := &Appdelegate{
		clientPacketQueue: make(chan *session.Message, CLIENT_PKT_QUEUE_SIZE),
	}

	return s
}
