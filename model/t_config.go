// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- config 包

package model

// /////////////////////////////////////////////////////////////////////////////
// world.ini

// world.ini 配置信息
type WorldIni struct {
	Env string // 当前运行环境，production= 开发环境；development = 运营环境
}

// /////////////////////////////////////////////////////////////////////////////
// servers.json 配置文件

// servers.json 的 server 服务器信息
type ServerInfo struct {
	Name       string // 服务器的名字
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

// /////////////////////////////////////////////////////////////////////////////
// servers.json 配置文件

// servers.json 的 server 服务器信息
type TServerInfo struct {
	Name       string // 服务器的名字
	Frontend   bool   // 是否是前端服务器
	Host       string // 服务器的ID地址
	Port       uint   // 服务器端口
	ClientHost string // 面向客户端的 IP地址
	CTcpPort   uint   // 面向客户端的 tcp端口
	CWsPort    uint   // 面向客户端的 websocket端口
}

// 服务器 type -> *[]ServerInfo 信息集合
type TServerMap map[string][]*ServerInfo

// server.json 配置表
type TServerConfig struct {
	Development TServerMap // 开发环境 配置信息
	Production  TServerMap // 运营环境 配置信息
}
