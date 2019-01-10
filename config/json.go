// /////////////////////////////////////////////////////////////////////////////
// world.ini 配置信息

package config

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

// 服务器配置信息列表
type ServetList []ServerInfo

// 服务器 type -> ServetList 信息集合
type ServerMap map[string]ServetList

// server.json 配置表
type ServerConfig struct {
	Development ServerMap // 开发环境 配置信息
	Production  ServerMap // 运营环境 配置信息
}
