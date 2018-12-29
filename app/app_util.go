// /////////////////////////////////////////////////////////////////////////////
// Application 一些辅助函数

package app

import (
	"flag"

	"github.com/zpab123/world/config" // 配置读取工具
	"github.com/zpab123/world/consts" // 全局常量
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

	// 加载 world.ini 信息

	// 设置开发环境

	// 加载 server.json 信息

	// 保存 cmd 命令参数
}

// 设置默认组件
func setDefaultComponent(app *Application) {
	// 根据服务器类型，注册默认组件
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
	name := app.baseInfo.AppName
	if "" == name {
		return
	}
	serverList := config.GetServerMap()[name]
	if nil == serverList {
		return
	}

	// 获取第1个 配置数据
	app.serverInfo = serverList[0]
}
