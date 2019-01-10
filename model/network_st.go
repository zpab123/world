// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- app 包

package model

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// Laddr 对象

// 监听地址
type Laddr struct {
	TcpAddr string // Tcp 监听地址：格式 192.168.1.1:8600
	WsAddr  string // websocket 监听地址: 格式 192.168.1.1:8600
	UdpAddr string // udp 监听地址: 格式 192.168.1.1:8600
	KcpAddr string // kcp 监听地址: 格式 192.168.1.1:8600
}

// /////////////////////////////////////////////////////////////////////////////
// ConnectorOpt 对象

// connector 组件配置参数
type ConnectorOpt struct {
	TypeName  string        // connector 类型
	Heartbeat time.Duration // 心跳间隔
	Handshake func()        // 自定义的握手处理函数
	MaxConn   uint32        // 最大连接数量，超过此数值后，不再接收新连接
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *ConnectorOpt) Check() error {

	return nil
}
