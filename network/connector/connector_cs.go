// /////////////////////////////////////////////////////////////////////////////
// connector 常量汇总

package connector

// connector 类型
const (
	CONNECTOR_TYPE_TCP       = "tcp"          // tcp
	CONNECTOR_TYPE_WEBSOCKET = "websocket"    // websocket
	CONNECTOR_TYPE_MUL       = "multiformity" // 同时支持 tcp 和 websocket
	CONNECTOR_TYPE_COM       = "composite"    // 同时支持 tcp 和 websocket
)

// tcp socket 默认参数
const (
	TCP_BUFFER_READ_SIZE  = 1024 * 1024 // 读 buffer 默认大小
	TCP_BUFFER_WRITE_SIZE = 1024 * 1024 // 写 buffer 默认大小
	TCP_NO_DELAY          = true        // net.tcpConn 对象写入数据后，是否立即发送
)
