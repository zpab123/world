// /////////////////////////////////////////////////////////////////////////////
// cmd 信息相关

package cmd

// Application 启动参数
type CmdParam struct {
	ServerType string // 服务器类型 例如 gate connect area ...等
	Gid        uint   // 进程 id 也就是相同类型的有多个服务器，Gid=切片下标
	Name       string // 服务器的名字
	Frontend   bool   // 是否是前端服务器
	Host       string // 服务器的Ip地址
	Port       uint   // 服务器端口
	ClientHost string // 面向客户端的 IP地址
	CTcpPort   uint   // 客户端 tcp 连接端口
	CWsPort    uint   // 客户端 websocket 连接端口
}
