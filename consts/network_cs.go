// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- network 包

// 全局 接口 定义
package consts

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// network 包

// connector 类型
const (
	NETWORK_CONNECTOR_TYPE_TCP = "tcpConnector" // tcp 连接器
)

const (
	TCP_SERVER_RECONNECT_TIME = 3 * time.Second // tcp 网络服务 开启失败后，重新开启时间，单位秒
)

// socket 常量
const (
	SOCKET_READ_BUFFSIZE  = 16384 // scoket 读取类 buff 长度
	SOCKET_WRITE_BUFFSIZE = 16384 // scoket 写入类 buff 长度
)
