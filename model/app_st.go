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

// 监听地址
type Laddr struct {
	TcpAddr string // Tcp 监听地址：格式 192.168.1.1:8600
	WsAddr  string // websocket 监听地址: 格式 192.168.1.1:8600
	UdpAddr string // udp 监听地址: 格式 192.168.1.1:8600
	KcpAddr string // kcp 监听地址: 格式 192.168.1.1:8600
}
