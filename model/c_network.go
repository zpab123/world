// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- network 包

// 全局 接口 定义
package model

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// network 包

// tcp socket 默认参数
const (
	C_TCP_BUFFER_READ_SIZE  = 1024 * 1024 // 读 buffer 默认大小
	C_TCP_BUFFER_WRITE_SIZE = 1024 * 1024 // 写 buffer 默认大小
	C_TCP_NO_DELAY          = true        // net.tcpConn 对象写入数据后，是否立即发送
)

// acceptor 类型
const (
	C_ACCEPTOR_NAME_TCP       = "tcpAcceptor"  // 支持 tcp
	C_ACCEPTOR_TYPE_WEBSOCKET = "wsAcceptor"   // 支持 websocket
	C_ACCEPTOR_TYPE_MUL       = "multiformity" // 同时支持 tcp 和 websocket
	C_ACCEPTOR_TYPE_COM       = "composite"    // 同时支持 tcp 和 websocket
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
	C_PACKET_HEAD_LEN                  = 4         // 消息头大小:字节 type(2字节) + length(2字节)
	C_PACKET_MAX_LEN                   = 64 * 1024 // 最大单个 packet 数据，= head + body （64K = 65536）
	C_PACKET_DATA_TCP_TLV              = "tcp.tlv" // type-length-value 形式的 packet 数据
	C_PACKET_TYPE_INVALID       uint16 = iota      // 无效的消息类型
	C_PACKET_TYPE_HANDSHAKE                        // 握手消息
	C_PACKET_TYPE_HANDSHAKE_ACK                    // 握手 ACK
	C_PACKET_TYPE_HEARTBEAT                        // 心跳消息
)

// 消息类型常量
