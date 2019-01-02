// /////////////////////////////////////////////////////////////////////////////
// Application 一些辅助函数

package app

import (
	"flag"

	"github.com/zpab123/world/component" // 组件库
	"github.com/zpab123/world/config"    // 配置读取工具
	"github.com/zpab123/world/consts"    // 全局常量
	"github.com/zpab123/world/model"     // 全局结构体
)

// 完成 app 的默认设置
func defaultConfiguration(app *Application) {
	// 解析命令行参数
	parseArgs(app)

	// 获取启动参数
	if app.runer == consts.APP_RUNER_CMD {
		// 从配置文件中获取服务器信息
		getInfoFromConfig(app)
	}
}

// 解析 命令行参数
func parseArgs(app *Application) {
	// 参数定义
	runer := flag.String("runer", "cmd", "runer") // 服务器启动者

	// 解析参数
	flag.Parse()

	//赋值
	if consts.APP_RUNER_MASTER == *runer {
		// 保存 runer
		app.runer = consts.APP_RUNER_MASTER
	} else {
		// 保存 runer
		app.runer = consts.APP_RUNER_CMD
	}
}

// 从配置文件中获取服务器信息
func getInfoFromConfig(app *Application) {
	// 获取运行环境
	env := config.GetWorldIni().Env // 当前环境
	app.baseInfo.Env = env

	// 获取 server.json 中关于 当前服务器的配置信息
	appType := app.baseInfo.AppType
	if "" == appType {
		return
	}
	serverList := config.GetServerMap()[appType]
	if nil == serverList {
		return
	}

	// 获取第1个 配置数据
	app.serverInfo = serverList[0]
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
	// 创建 Connector
	con := component.NewConnector(app.connectorConfig)
	con.TcpAddr = app.GetCTcpAddr()
	con.WsAddr = app.GetCWsAddr()

	// 注册组件
	app.RegisterComponent(con)
}
