// /////////////////////////////////////////////////////////////////////////////
// 全局模型 -- connector 包

package model

// /////////////////////////////////////////////////////////////////////////////
// 常量

// 最大连接数
const (
	C_CNTOR_MAX_CONN = 100000 // connector 默认最大连接数
)

// tcp socket 默认参数
const (
	C_TCP_BUFFER_READ_SIZE  = 1024 * 1024 // 读 buffer 默认大小
	C_TCP_BUFFER_WRITE_SIZE = 1024 * 1024 // 写 buffer 默认大小
	C_TCP_NO_DELAY          = true        // net.tcpConn 对象写入数据后，是否立即发送
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// connector 组件
type IConnector interface {
	IComponent // 接口继承： 组件接口
}

// /////////////////////////////////////////////////////////////////////////////
// TConnectorOpt 对象

// connector 组件配置参数
type TConnectorOpt struct {
	AcceptorName string        // 接收器 名字
	MaxConn      uint32        // 最大连接数量，超过此数值后，不再接收新连接
	TcpConnOpts  *TTcpConnOpts // tcpSocket 配置参数
}

// 创建1个新的 TConnectorOpt
func NewTConnectorOpt() *TConnectorOpt {
	// 创建 tcp 配置参数
	tcpOpts := NewTTcpConnOpts()

	// 创建对象
	opts := &TConnectorOpt{
		TcpConnOpts: tcpOpts,
	}

	opts.SetDefaultOpts()

	return opts
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TConnectorOpt) Check() error {
	if this.MaxConn <= 0 {
		this.MaxConn = C_CNTOR_MAX_CONN
	}
	return nil
}

// 设置默认参数
func (this *TConnectorOpt) SetDefaultOpts() {
	this.AcceptorName = C_ACCEPTOR_NAME_COM
	this.MaxConn = C_CNTOR_MAX_CONN
}

// /////////////////////////////////////////////////////////////////////////////
// TTcpConnOpts 对象

// TcpSocket 配置参数
type TTcpConnOpts struct {
	ReadBufferSize  int           // 读取 buffer 字节大小
	WriteBufferSize int           // 写入 buffer 字节大小
	NoDelay         bool          // 写入数据后，是否立即发送
	MaxPacketSize   int           // 单个 packet 最大字节数
	ReadTimeout     time.Duration // 读数据超时时间
	WriteTimeout    time.Duration // 写数据超时时间
}

// 创建1个新的 TTcpConnOpts 对象
func NewTTcpConnOpts() *TTcpConnOpts {
	// 创建对象
	tcpOpts := &TTcpConnOpts{}
	tcpOpts.SetDefaultOpts()

	return tcpOpts
}

// 设置默认参数
func (this *TTcpConnOpts) SetDefaultOpts() {
	this.ReadBufferSize = C_TCP_BUFFER_READ_SIZE   // 张鹏：原先是-1，这里被修改了
	this.WriteBufferSize = C_TCP_BUFFER_WRITE_SIZE // 张鹏：原先是-1，这里被修改了
	this.NoDelay = C_TCP_NO_DELAY                  // 张鹏：原先没有这个设置项，这里被修改了
}
