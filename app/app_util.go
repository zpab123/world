// /////////////////////////////////////////////////////////////////////////////
// Application 一些辅助函数

package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/zpab123/world/acceptor"   // acceptor 组件
	"github.com/zpab123/world/config"     // 配置读取工具
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

	// 组件默认参数
	defaultComponentOpt(app)
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
	// 根据 AppType 和 Name 获取 服务器配置参数
	appType := app.baseInfo.AppType
	name := app.baseInfo.Name
	list, ok := config.GetServerMap()[appType]

	if nil == list || len(list) <= 0 || !ok {
		zaplog.Fatal("Application 获取 appType 信息失败。 appType=%s", appType)

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
		zaplog.Fatal("Application 获取 server.json 信息失败。 appName=%s", app.baseInfo.Name)

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
	logFile := fmt.Sprintf("./logs/%s.log", app.baseInfo.Name)
	var outputs []string
	stdErr := config.GetWorldConfig().LogStderr
	if stdErr {
		outputs = append(outputs, "stderr")
	}
	outputs = append(outputs, logFile)
	zaplog.SetOutput(outputs)
}

// 设置组件默认参数
func defaultComponentOpt(app *Application) {
	// Acceptor 组件
	if nil == app.componentMgr.GetAcceptorOpt() {
		opt := getDefaultAcceptorOpt(app)
		app.componentMgr.SetAcceptorOpt(opt)
	}

	// DispatcherClient 组件
	if nil == app.componentMgr.GetDisClientOpt() {
		opt := getDefaultDisClientOpt(app)
		app.componentMgr.SetDisClientOpt(opt)
	}
}

// 获取默认 AcceptorOpt
func getDefaultAcceptorOpt(app *Application) *acceptor.TAcceptorOpt {
	opt := acceptor.NewTAcceptorOpt(app.appDelegate)

	return opt
}

// 获取默认 DispatcherClientOpt
func getDefaultDisClientOpt(app *Application) *dispatcher.TDispatcherClientOpt {
	opt := dispatcher.NewTDispatcherClientOpt(app.appDelegate)

	return opt
}

// 创建默认组件
func createComponent(app *Application) {
	// master 服务器 - 注册 master 组件

	// 网络连接组件
	acceptorOpt := app.componentMgr.GetAcceptorOpt()
	if nil != acceptorOpt && acceptorOpt.Enable {
		newAcceptor(app)
	}

	// 消息分发客户端
	disClientOpt := app.componentMgr.GetDisClientOpt()
	if nil != disClientOpt && disClientOpt.Enable {
		newDisClient(app)
	}

	// 注册 backendSession 组件

	// 注册 channel 组件

	// 注册 server 组件

	// 注册 monitor 组件
}

// 创建 Acceptor 组件
func newAcceptor(app *Application) {
	// 创建地址
	serverInfo := app.GetServerInfo()
	opt := app.GetComponentMgr().GetAcceptorOpt()

	var tcpAddr string = ""
	if opt.ForClient && serverInfo.CTcpPort > 0 {
		tcpAddr = fmt.Sprintf("%s:%d", serverInfo.ClientHost, serverInfo.CTcpPort) // 面向客户端的 tcp 地址
	} else if serverInfo.Port > 0 {
		tcpAddr = fmt.Sprintf("%s:%d", serverInfo.Host, serverInfo.Port) // 面向服务器的 tcp 地址
	}

	var wsAddr string = ""
	if opt.ForClient && serverInfo.CWsPort > 0 {
		wsAddr = fmt.Sprintf("%s:%d", serverInfo.ClientHost, serverInfo.CWsPort) // 面向客户端的 websocket 地址
	} else if serverInfo.Port > 0 {
		wsAddr = fmt.Sprintf("%s:%d", serverInfo.Host, serverInfo.Port) // 面向服务器的 websocket 地址
	}

	laddr := &network.TLaddr{
		TcpAddr: tcpAddr,
		WsAddr:  wsAddr,
	}

	// 创建 Connector
	actor, err := acceptor.NewAcceptor(laddr, opt)
	if nil != err {
		return
	}

	if nil != actor {
		app.componentMgr.AddComponent(actor)
	}

}

// 创建分发客户端
func newDisClient(app *Application) {
	opt := app.componentMgr.GetDisClientOpt()

	disClient, err := dispatcher.NewDispatcherClient(opt)
	if nil != err {
		return
	}

	if nil != disClient {
		app.componentMgr.AddComponent(disClient)
	}
}
