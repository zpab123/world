// /////////////////////////////////////////////////////////////////////////////
// 全局模型 -- connector 包

package model

// /////////////////////////////////////////////////////////////////////////////
// 常量

// 最大连接数
const (
	C_CNTOR_MAX_CONN = 100000 // connector 默认最大连接数
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
	return nil
}

// 设置默认参数
func (this *TConnectorOpt) SetDefaultOpts() {
	this.AcceptorName = C_ACCEPTOR_NAME_COM
	this.MaxConn = C_CNTOR_MAX_CONN

}
