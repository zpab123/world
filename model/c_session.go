// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- session 包

package model

// Session 类型
const (
	C_SES_TYPE_CLINET = "client" // 客户端类型的 session
)

// Session 状态
const (
	C_SES_STATE_INITED   uint32 = iota // 初始化状态
	C_SES_STATE_WAIT_ACK               // 等待握手ACK
	C_SES_STATE_WORKING                // 工作中
	C_SES_STATE_CLOSED                 // 关闭状态
)
