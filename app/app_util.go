// /////////////////////////////////////////////////////////////////////////////
// Application 一些辅助函数

package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/zpab123/world/config"     // 配置读取工具
	"github.com/zpab123/world/connector"  // connector 组件
	"github.com/zpab123/world/dispatcher" // dispatcher 组件
	"github.com/zpab123/world/network"    // 网络库
	"github.com/zpab123/zaplog"           // log 库
)

// 完成 app 的默认设置
func defaultConfig(app *Application) {
	// 解析命令行参数
	parseArgs(app)

	// 获取服务器信息
	getServerInfo(app)

	// 设置 log 信息
	configLogger(app)
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
	appType := app.baseInfo.AppType

	// 获取服务器名字
	name := app.baseInfo.Name

	// 获取类型列表
	list := config.GetServerMap()[appType]
	if nil == list || len(list) <= 0 {
		zaplog.Error("Application 获取 appType 信息失败。 appType=%s", appType)

		os.Exit(1)
	}

	// 获取服务器信息
	for _, info := range list {
		if info.Name == name {
			app.serverInfo = info

			break
		}
	}

	if app.serverInfo == nil {
		zaplog.Errorf("Application 获取 server.json 信息失败。 appName=%s", app.baseInfo.Name)

		os.Exit(1)
	}
}

// 设置 log 信息
func configLogger(app *Application) {
	// 模块名字
	zaplog.SetSource(app.baseInfo.Name)

	// 输出等级
	lv := config.GetWorldConfig().LogLevel
	zaplog.SetLevel(zaplog.ParseLevel(lv))

	// 输出文件
	logFile := fmt.Sprintf("%s.log", app.baseInfo.Name)
	var outputs []string
	stdErr := config.GetWorldConfig().LogStderr
	if stdErr {
		outputs = append(outputs, "stderr")
	}
	outputs = append(outputs, logFile)
	zaplog.SetOutput(outputs)
}

// 创建默认组件
func createComponent(app *Application) {
	// master 服务器 - 注册 master 组件

	// 前端服务器
	if app.serverInfo.Frontend {

		// 创建1个 Connector 组件
		newConnector(app)

		// 创建分发服务器
		newDispatcherServer(app)
	}

	// 注册 backendSession 组件

	// 注册 channel 组件

	// 注册 server 组件

	// 注册 monitor 组件
}

// 创建 Connector 组件
func newConnector(app *Application) {
	// 获取地址信息
	serverInfo := app.GetServerInfo()

	var cTcpAddr string = ""
	if serverInfo.CTcpPort > 0 {
		cTcpAddr = fmt.Sprintf("%s:%d", serverInfo.ClientHost, serverInfo.CTcpPort) // 面向客户端的 tcp 地址
	}

	var cWsAddr string = ""
	if serverInfo.CWsPort > 0 {
		cWsAddr = fmt.Sprintf("%s:%d", serverInfo.ClientHost, serverInfo.CWsPort) // 面向客户端的 websocket 地址
	}

	laddr := &network.TLaddr{
		TcpAddr: cTcpAddr,
		WsAddr:  cWsAddr,
	}

	// connector 参数
	opts := app.componentMgr.GetConnectorOpt()
	if nil == opts {
		opts = getDefaultConnectorOpt(app)
	}

	// 创建 Connector
	contor := connector.NewConnector(laddr, opts)

	// 添加组件
	app.componentMgr.AddComponent(contor)
}

// 创建分发服务器
func newDispatcherServer(app *Application) {
	// 地址参数
	serverInfo := app.GetServerInfo()
	tcpAddr := fmt.Sprintf("%s:%d", serverInfo.Host, serverInfo.Port)
	laddr := &network.TLaddr{
		TcpAddr: tcpAddr,
	}

	// 配置参数
	opt := app.componentMgr.GetDisServerOpt()
	if nil == opt {
		opt = getDefaultDisServerOpt(app)
		app.componentMgr.SetDisServerOpt(opt)
	}

	// 创建 DispatcherServer
	dis := dispatcher.NewDispatcherServer(laddr, opt)

	// 添加组件
	app.componentMgr.AddComponent(dis)
}

// 获取默认 ConnectorOpt
func getDefaultConnectorOpt(app *Application) *connector.TConnectorOpt {
	// 创建默认
	opts := connector.NewTConnectorOpt(app.appDelegate)

	return opts
}

// 获取默认 DispatcherServerOpt
func getDefaultDisServerOpt(app *Application) *dispatcher.TDispatcherServerOpt {
	opt := dispatcher.NewTDispatcherServerOpt(app.appDelegate)

	return opt
}
