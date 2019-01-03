// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- component 包

// 全局 接口 定义
package consts

// /////////////////////////////////////////////////////////////////////////////
// 通用

// 组件名字
const (
	COMPONENT_NAME_CONNECTOR = "connector" // connector 组件
)

// /////////////////////////////////////////////////////////////////////////////
// connector 组件

// socket buffer 大小

const (
	CONNECTOR_SOCKET_WRITE_BUF_SIZE = 1024 * 1024 // connector <-> client 之间 socket 的写入类buffer 大小
	CONNECTOR_SOCKET_READ_BUF_SIZE  = 1024 * 1024 // connector <-> client 之间 socket 的读取类buffer 大小
	CONNECTOR_SOCKET_NO_DELAY       = true        // connector <-> client 之间 socket 写入数据后，是否立即发送
)
