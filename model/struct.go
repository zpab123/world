// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct

package model

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// app 包

// Application 基础属性
type BaseInfo struct {
	Env      string // Application 当前运行环境 production= 开发环境 development = 运营环境
	AppType  string // 当前app的类型
	Name     string // 当前app的名字
	MainPath string // 程序根路径 例如 "E/server/gateserver.exe" 所在的目录 “E/server/”
}

// connector 组件配置参数
type ConnectorConfig struct {
	Heartbeat time.Duration // 心跳间隔
	Handshake func()        // 自定义的握手处理函数
	MaxConn   uint32        // 最大连接数量，超过此数值后，不在接收新连接
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *ConnectorConfig) Check() error {

	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// config 包

// world.ini 配置信息
type WorldIni struct {
	Env string // 当前运行环境，production= 开发环境；development = 运营环境
}

// /////////////////////////////////////////////////////////////////////////////
// servers.json 配置文件

// servers.json 的 server 服务器信息
type ServerInfo struct {
	Name string // 服务器的名字
	// ServiceId  uint8  // 服务ID（用于gate服务器消息转发）
	Frontend   bool   // 是否是前端服务器
	Host       string // 服务器的ID地址
	Port       uint   // 服务器端口
	ClientHost string // 面向客户端的 IP地址
	CTcpPort   uint   // 面向客户端的 tcp端口
	CWsPort    uint   // 面向客户端的 websocket端口
}

// 服务器 type -> *[]ServerInfo 信息集合
type ServerMap map[string][]ServerInfo

// server.json 配置表
type ServerConfig struct {
	Development map[string][]ServerInfo // 开发环境 配置信息
	Production  map[string][]ServerInfo // 运营环境 配置信息
}
