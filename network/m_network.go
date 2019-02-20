// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package network

import (
	"net"
	"time"

	"github.com/zpab123/world/model" // 全局模型
	"golang.org/x/net/websocket"     // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// acceptor 类型
const (
	C_ACCEPTOR_TYPE_TCP       = "tcpAcceptor"  // 支持 tcp
	C_ACCEPTOR_TYPE_WEBSOCKET = "wsAcceptor"   // 支持 websocket
	C_ACCEPTOR_TYPE_MUL       = "multiformity" // 同时支持 tcp 和 websocket
	C_ACCEPTOR_TYPE_COM       = "composite"    // 同时支持 tcp 和 websocket
)

const (
	TCP_SERVER_RECONNECT_TIME = 3 * time.Second // tcp 网络服务 开启失败后，重新开启时间，单位秒
)

// socket 常量
const (
	C_BUFF_READ_SIZE  = 16384 // scoket 读取类 buff 长度
	C_BUFF_WRITE_SIZE = 16384 // scoket 写入类 buff 长度
)

// packet 常量
const (
	C_PACKET_HEAD_LEN     = 6                // 消息头大小:字节 type(2字节) + length(4字节)
	C_PACKET_MAX_LEN      = 25 * 1024 * 1024 // 最大单个 packet 数据，= head + body = 25M
	C_PACKET_DATA_TCP_TLV = "tcp.tlv"        // type-length-value 形式的 packet 数据
)

// packet ID
const (
	C_PACKET_ID_INVALID       uint16 = iota // 无效的消息类型
	C_PACKET_ID_HANDSHAKE                   // 握手消息
	C_PACKET_ID_HANDSHAKE_ACK               // 握手 ACK
	C_PACKET_ID_HEARTBEAT                   // 心跳消息
	C_PACKET_ID_WORLD                       // 分界线： 以上由 WorldConnection 处理的消息
)

// WorldConnection 状态
const (
	C_WCONN_STATE_INIT     uint32 = iota // 初始化状态
	C_WCONN_STATE_WAIT_ACK               // 等待客户端握手ACK
	C_WCONN_STATE_WORKING                // 工作中
	C_WCONN_STATE_CLOSED                 // 关闭状态
)

// WorldSocket 类型
const (
	C_NET_TYPE_TCP = "tcp" // tcp 连接对象
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// socket 组件
type ISocket interface {
	net.Conn // 接口继承： 符合 Conn 的对象
	Flush() error
}

// acceptor 接口
type IAcceptor interface {
	Run() bool  // 组件开始运行
	Stop() bool // 组件停止运行
}

// tcpSocket 连接管理
type ITcpConnManager interface {
	OnNewTcpConn(conn net.Conn) // 收到1个新的 Tcp 连接对象
}

// websocket 管理
type IWsConnManager interface {
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
}

// MulAcceptor 连接管理
type IMulConnManager interface {
	OnNewTcpConn(conn net.Conn)         // 收到1个新的 Tcp 连接对象
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
}

// ComAcceptor 连接管理
type IComConnManager interface {
	OnNewTcpConn(conn net.Conn)         // 收到1个新的 Tcp 连接对象
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
}

// /////////////////////////////////////////////////////////////////////////////
// Laddr 对象

// 监听地址集合
type TLaddr struct {
	TcpAddr string // Tcp 监听地址：格式 192.168.1.1:8600
	WsAddr  string // websocket 监听地址: 格式 192.168.1.1:8600
	UdpAddr string // udp 监听地址: 格式 192.168.1.1:8600
	KcpAddr string // kcp 监听地址: 格式 192.168.1.1:8600
}

// /////////////////////////////////////////////////////////////////////////////
// TBufferSocketOpt 对象

// BufferSocket 配置参数
type TBufferSocketOpt struct {
	ReadBufferSize  int // 读取 buffer 字节大小
	WriteBufferSize int // 写入 buffer 字节大小
}

// 新建1个 TBufferSocketOpt 对象
func NewTBufferSocketOpt() *TBufferSocketOpt {
	bs := &TBufferSocketOpt{
		ReadBufferSize:  C_BUFF_READ_SIZE,
		WriteBufferSize: C_BUFF_WRITE_SIZE,
	}

	return bs
}

// /////////////////////////////////////////////////////////////////////////////
// TWorldConnOpt 对象

// WorldConnection 配置参数
type TWorldConnOpt struct {
	ShakeKey       string            // 握手key
	Heartbeat      int64             // 心跳间隔，单位：秒。0=不设置心跳
	BuffSocketOpts *TBufferSocketOpt // BufferSocket 配置参数
}

// 新建1个 WorldConnection 对象
func NewTWorldConnOpt() *TWorldConnOpt {
	// 创建 buff opt
	buffOpt := NewTBufferSocketOpt()

	opt := &TWorldConnOpt{
		BuffSocketOpts: buffOpt,
	}

	return opt
}

// /////////////////////////////////////////////////////////////////////////////
// TWorldConnClientOpt 对象

// WorldConnClient 配置参数
type TWorldSocketOpt struct {
	NetType        string             // 类型
	ShakeKey       string             // 握手key
	TcpConnOpt     *model.TTcpConnOpt // tcpSocket 配置参数
	BuffSocketOpts *TBufferSocketOpt  // BufferSocket 配置参数
}

// 新建1个 TWorldConnClientOpt 对象
func NewTWorldSocketOpt() *TWorldSocketOpt {
	// 创建对象
	tcpOpt := model.NewTTcpConnOpt()
	buffOpt := NewTBufferSocketOpt()

	opt := &TWorldSocketOpt{
		NetType:        C_NET_TYPE_TCP,
		TcpConnOpt:     tcpOpt,
		BuffSocketOpts: buffOpt,
	}

	return opt
}
