// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- worldnet 包

// 全局 接口 定义
package consts

import (
	"time"
)

// connector 类型
const (
	CONNECTOR_TYPE_TCP_ACCEPTOR = "tcpConnector" // tcpConnector 类型
)

// session 关闭原因
const (
	SESSION_CLOSE_REASON_IO      = "IO"      // 普通 IO 断开
	SESSION_CLOSE_REASON_MANUAL  = "Manual"  // 关闭前，调用过Session.Close
	SESSION_CLOSE_REASON_UNKNOWN = "Unknown" // 未知错误
)
