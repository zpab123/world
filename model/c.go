// /////////////////////////////////////////////////////////////////////////////
// 全局常量

package model

// 服务器类型
const (
	C_SERVER_TYPE_GATE = "gate" // gate 服务器
)

// tcp socket 默认参数
const (
	C_TCP_BUFFER_READ_SIZE  = 1024 * 1024 // 读 buffer 默认大小
	C_TCP_BUFFER_WRITE_SIZE = 1024 * 1024 // 写 buffer 默认大小
	C_TCP_NO_DELAY          = true        // net.tcpConn 对象写入数据后，是否立即发送
)
