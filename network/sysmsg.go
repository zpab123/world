// /////////////////////////////////////////////////////////////////////////////
// 系统消息

package worldnet

import (
	"github.com/zpab123/world/consts" // 全局常量
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 常量
const (
	CloseReason_IO     CloseReason = iota // 普通IO断开
	CloseReason_Manual                    // 关闭前，调用过Session.Close
)

// /////////////////////////////////////////////////////////////////////////////
// SessionClosed 对象

// session 断开
type SessionClosed struct {
	Reason CloseReason // 断开原因
}

// 字符串化
func (self *SessionClosed) String() string {
	return fmt.Sprintf("%+v", *self)
}

// 使用类型断言判断是否为系统消息
func (self *SessionClosed) SystemMessage() {}

// /////////////////////////////////////////////////////////////////////////////
// SessionAccepted 对象

// 接入新的 session
type SessionAccepted struct {
}

// 字符串化
func (self *SessionAccepted) String() string {
	return fmt.Sprintf("%+v", *self)
}

// 使用类型断言判断是否为系统消息
func (self *SessionAccepted) SystemMessage() {}

// /////////////////////////////////////////////////////////////////////////////
// CloseReason 对象

// 关闭原因
type CloseReason int32

// fmt 打印接口
func (self CloseReason) String() string {
	switch self {
	case CloseReason_IO:
		return consts.SESSION_CLOSE_REASON_IO
	case CloseReason_Manual:
		return consts.SESSION_CLOSE_REASON_MANUAL
	}

	return consts.SESSION_CLOSE_REASON_UNKNOWN
}
