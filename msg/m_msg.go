// /////////////////////////////////////////////////////////////////////////////
// msg 包模型

package msg

// /////////////////////////////////////////////////////////////////////////////
// 常量

// 消息 id
const (
	PKT_ID_SHAKE uint16 = iota + 1 // 握手消息 ID
	CODE_SUCCESS                   // 成功类消息 1
)

// 通用消息码(1-1000)
const (
	CODE_ERROR   uint32 = iota // 误类消息 0
	CODE_SUCCESS               // 成功类消息 1
)

// 其他消息(1001-)
const (
	SHAKE_KEY_ERROR      uint32 = iota + 1001 // 握手 key 消息错误 1001
	SHAKE_ACCEPTOR_ERROR                      // 网络方式错误 1002
)
