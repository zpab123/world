// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- app 包

package model

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// Application 基础属性

// Application 基础属性
type BaseInfo struct {
	Env      string // Application 当前运行环境 production= 开发环境 development = 运营环境
	AppType  string // 当前app的类型
	Name     string // 当前app的名字
	MainPath string // 程序根路径 例如 "E/server/gateserver.exe" 所在的目录 “E/server/”
}

// /////////////////////////////////////////////////////////////////////////////
// connector 组件

// connector 组件配置参数
type ConnectorConfig struct {
	Heartbeat time.Duration // 心跳间隔
	Handshake func()        // 自定义的握手处理函数
	MaxConn   uint32        // 最大连接数量，超过此数值后，不再接收新连接
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *ConnectorConfig) Check() error {

	return nil
}

// 监听地址
type Laddr struct {
	TcpAddr string // Tcp 监听地址：格式 192.168.1.1:8600
	WsAddr  string // websocket 监听地址: 格式 192.168.1.1:8600
	UdpAddr string // udp 监听地址: 格式 192.168.1.1:8600
	KcpAddr string // kcp 监听地址: 格式 192.168.1.1:8600
}
