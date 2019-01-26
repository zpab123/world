// /////////////////////////////////////////////////////////////////////////////
// 全局基础模型

package model

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/net/websocket" // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// acceptor 名字
const (
	C_ACCEPTOR_NAME_TCP       = "tcpAcceptor"  // 支持 tcp
	C_ACCEPTOR_NAME_WEBSOCKET = "wsAcceptor"   // 支持 websocket
	C_ACCEPTOR_NAME_MUL       = "multiformity" // 同时支持 tcp 和 websocket
	C_ACCEPTOR_NAME_COM       = "composite"    // 同时支持 tcp 和 websocket
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
	C_PACKET_HEAD_LEN                = 6                // 消息头大小:字节 type(2字节) + length(4字节)
	C_PACKET_MAX_LEN                 = 25 * 1024 * 1024 // 最大单个 packet 数据，= head + body = 25M
	C_PACKET_DATA_TCP_TLV            = "tcp.tlv"        // type-length-value 形式的 packet 数据
	C_PACKET_ID_INVALID       uint16 = iota             // 无效的消息类型
	C_PACKET_ID_HANDSHAKE                               // 握手消息
	C_PACKET_ID_HANDSHAKE_ACK                           // 握手 ACK
	C_PACKET_ID_HEARTBEAT                               // 心跳消息
)

// WorldConnection 状态
const (
	C_WCONN_STATE_INIT     = iota // 初始化状态
	C_WCONN_STATE_WAIT_ACK        // 等待客户端握手ACK
	C_WCONN_STATE_WORKING         // 工作中
	C_WCONN_STATE_CLOSED          // 关闭状态
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// tcpSocket 连接管理
type ITcpSocketManager interface {
	OnNewTcpConn(conn net.Conn) // 收到1个新的 Tcp 连接对象
	CloseAllConn()              // 关闭所有连接
}

// websocket 连接管理
type IWebsocketManager interface {
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
	CloseAllConn()                      // 关闭所有连接
}

// MulAcceptor 连接管理
type IMulConnManager interface {
	OnNewTcpConn(conn net.Conn)         // 收到1个新的 Tcp 连接对象
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
	CloseAllConn()                      // 关闭所有连接
}

// ComAcceptor 连接管理
type IComConnManager interface {
	OnNewTcpConn(conn net.Conn)         // 收到1个新的 Tcp 连接对象
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
	CloseAllConn()                      // 关闭所有连接
}

// acceptor 接口
type IAcceptor interface {
	Run()  // 组件开始运行
	Stop() // 组件停止运行
}

// socket 组件
type ISocket interface {
	net.Conn // 接口继承： 符合 Conn 的对象
	Flush() error
}

// 可以收发 packet 的 socket
type IPacketSocket interface {
	ISocket                   // 接口继承： 符合 ISocket 的对象
	GetSocket() ISocket       // 获取 ISocket 对象
	GetConnector() IConnector // 获取 PacketSocket 对象所属 Connector
}

// 网络数据收发管理
type IDataManager interface {
	// 接收 1个 packet
	RecvPacket(socket IPacketSocket) (pkt interface{}, err error)
	// 发送 1个 packet
	SendPacket(socket IPacketSocket, pkt interface{}) error
}

// packet 消息管理接口
type IPacketManager interface {
	// 设置网络数据 接受/发送对象
	SetDataManager(dm IDataManager)
	// 设置 接收后，发送前的事件处理流程
	//SetHooker
	// 设置 接收后最终处理回调
	//SetCallback(v cellnet.EventCallback)
}

// packet 处理接口
type IPacketHandler interface {
	OnMessage(ses ISession, msg interface{})
}

// /////////////////////////////////////////////////////////////////////////////
// TAddress 对象

// 支持地址范围的格式
type TAddress struct {
	Scheme  string
	Host    string
	MinPort int
	MaxPort int
	Path    string
}

// 参数检查,正确返回 true 错误 返回 fasle
func (this *TAddress) Check() (bool, error) {
	// 地址效验 -- 正则是否是ip 地址

	// 端口效验

	// 最大端口效验
	if this.MaxPort < this.MinPort {
		this.MaxPort = this.MinPort
	}

	return true, nil
}

// 获取带范围的 addr 格式
//
// 例如：scheme://host:minPort~maxPort/path
func (this *TAddress) GetAddrRange() (string, error) {
	// 参数检查
	ok, err := this.Check()
	if !ok {
		return "", err
	}

	// 获取地址
	var addr string
	if "" == this.Scheme {
		addr = fmt.Sprintf("%s:%d~%d", this.Host, this.MinPort, this.MaxPort)
	} else {
		addr = fmt.Sprintf("%s://%s:%d~%d/%s", this.Scheme, this.Host, this.MinPort, this.MaxPort, this.Path)
	}

	return addr, nil
}

// 根据 port 参数，与 TAddress 对象的 host 组成1个 addr 字符
//
// 返回格式： 192.168.1.1:6002
func (this *TAddress) HostPortString(port int) string {
	return fmt.Sprintf("%s:%d", this.Host, port)
}

// 根据 port 参数，与 TAddress 对象的 Scheme host Path 组成1个完整 addr 字符
//
// 返回格式： http://192.168.1.1:6002/romte
func (this *TAddress) String(port int) string {
	if this.Scheme == "" {
		return this.HostPortString(port)
	}

	return fmt.Sprintf("%s://%s:%d%s", this.Scheme, this.Host, port, this.Path)
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

// BufferSocket 配置参数
type TBufferSocketOpts struct {
	ReadBufferSize  int // 读取 buffer 字节大小
	WriteBufferSize int // 写入 buffer 字节大小
}

// 新建1个 TBufferSocketOpts 对象
func NewTBufferSocketOpts() *TBufferSocketOpts {
	bs := &TBufferSocketOpts{
		ReadBufferSize:  C_SOCKET_READ_BUFFSIZE,
		WriteBufferSize: C_SOCKET_WRITE_BUFFSIZE,
	}

	return bs
}

// /////////////////////////////////////////////////////////////////////////////
// TWorldConnOpts 对象

// WorldConnection 配置参数
type TWorldConnOpts struct {
	Heartbeat time.Duration // 心跳间隔
	Handshake func()        // 自定义的握手处理函数
}

// 新建1个 WorldConnection 对象
func NewTWorldConnOpts() *TWorldConnOpts {
	wc := &TWorldConnOpts{
		Heartbeat: time.Second,
	}
}