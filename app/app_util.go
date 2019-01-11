// /////////////////////////////////////////////////////////////////////////////
// Application 一些辅助函数

package app

import (
	"flag"

	"github.com/zpab123/world/config"            // 配置读取工具
	"github.com/zpab123/world/model"             // 全局数据类型库
	"github.com/zpab123/world/network/connector" // 网络连接库
)

// 完成 app 的默认设置
func defaultConfiguration(app *Application) {
	// 解析命令行参数
	parseArgs(app)

	// 获取服务器信息
	getServerInfo(app)
}

// 解析 命令行参数
func parseArgs(app *Application) {
	// 参数定义
	//serverType := flag.String("type", "serverType", "server type") // 服务器类型，例如 gate connect area ...等
	//gid := flag.Uint("gid", 0, "gid")                                       // 服务器进程id
	name := flag.String("name", "master", "server name") // 服务器名字
	//frontend := flag.Bool("frontend", false, "is frontend server")          // 是否是前端服务器
	//host := flag.String("host", "127.0.0.1", "server host")                 // 服务器IP地址
	//port := flag.Uint("port", 0, "server port")                             // 服务器端口
	//clientHost := flag.String("clientHost", "127.0.0.1", "for client host") // 面向客户端的 IP地址
	//cTcpPort := flag.Uint("cTcpPort", 0, "tcp port")                        // tcp 连接端口
	//cWsPort := flag.Uint("cWsPort", 0, "websocket port")                    // websocket 连接端口

	// 解析参数
	flag.Parse()

	// 赋值
	//cmdParam := &cmd.CmdParam{
	//ServerType: *serverType,
	//Gid:        *gid,
	//Name: *name,
	//Frontend:   *frontend,
	//Host:       *host,
	//Port:       *port,
	//ClientHost: *clientHost,
	//CTcpPort:   *cTcpPort,
	//CWsPort:    *cWsPort,
	//}

	// 设置 app 名字
	app.baseInfo.Name = *name
}

// 获取服务器信息
func getServerInfo(app *Application) {
	// 获取服务器类型
	serverType := app.baseInfo.ServerType

	// 获取服务器名字
	name := app.baseInfo.Name

	// 获取类型列表
	list := config.GetServerMap()[serverType]
	if len(list) <= 0 {
		return
	}

	// 获取服务器信息
	for _, info := range list {
		if info.Name == name {
			app.serverInfo = info
		}
	}
}

// 设置默认组件
func regDefaultComponent(app *Application) {
	// master 服务器 - 注册 master 组件

	// 前端服务器
	if app.serverInfo.Frontend {

		// 注册1个 Connector 组件
		regConnector(app)

		// 注册 session 组件
	}

	// 注册 backendSession 组件

	// 注册 channel 组件

	// 注册 server 组件

	// 注册 monitor 组件
}

// 注册1个 Connector 组件
func regConnector(app *Application) {
	// 地址参数
	laddr := &model.Laddr{
		TcpAddr: app.GetCTcpAddr(),
		WsAddr:  app.GetCWsAddr(),
	}

	// connector 参数
	opts := app.GetConnectorOpt()
	if nil == opts {
		opts = getDefaultConnectorOpt()
	}

	// 创建 Connector
	contor := connector.NewConnector(laddr, opts)

	// 注册组件
	app.RegisterComponent(contor)
}

// 获取默认 ConnectorOpt
func getDefaultConnectorOpt() *model.ConnectorOpt {
	// 创建默认
	opts := &model.ConnectorOpt{
		TypeName: connector.CONNECTOR_TYPE_MUL, // 默认支持多种
	}

	return opts
}
