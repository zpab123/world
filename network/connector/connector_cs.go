// /////////////////////////////////////////////////////////////////////////////
// connector 常量汇总

package connector

// connector 类型
const (
	TYPE_TCP_CONNECTOR    = "tcpConnector"   // tcpConnector 类型
	TYPE_HYBRID_CONNECTOR = "hybridconnecto" // hybridconnecto 类型
)

// tcp socket 默认参数
const (
	TCP_BUFFER_READ_SIZE  = 1024 * 1024 // 读 buffer 默认大小
	TCP_BUFFER_WRITE_SIZE = 1024 * 1024 // 写 buffer 默认大小
	TCP_NO_DELAY          = true        // net.tcpConn 对象写入数据后，是否立即发送
)

// session 关闭原因
const (
	SESSION_CLOSE_REASON_IO      = "IO"      // 普通 IO 断开
	SESSION_CLOSE_REASON_MANUAL  = "Manual"  // 关闭前，调用过Session.Close
	SESSION_CLOSE_REASON_UNKNOWN = "Unknown" // 未知错误
)
