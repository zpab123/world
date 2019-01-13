// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- network 包

// 全局 接口 定义
package model

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
	C_SOCKET_READ_BUFFSIZE  = 16384 // scoket 读取类 buff 长度
	C_SOCKET_WRITE_BUFFSIZE = 16384 // scoket 写入类 buff 长度
)

// packet 常量
const (
	C_PACKET_HEAD_LEN = 4         // 消息头大小:字节 type(2字节) + length(2字节)
	C_PACKET_MAX_LEN  = 64 * 1024 // 最大单个 packet 数据，= head + body （64K = 65536）
)
