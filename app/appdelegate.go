// /////////////////////////////////////////////////////////////////////////////
// Application 代理

package app

import (
	"github.com/zpab123/world/session" // session库
)

// app 服务
type Appdelegate struct {
	clientPacketQueue chan *session.Message
}

// 创建1个新的 Appdelegate
func NewAppdelegate() *Appdelegate {
	s := &Appdelegate{
		clientPacketQueue: make(chan *session.Message, CLIENT_PKT_QUEUE_SIZE),
	}

	return s
}
