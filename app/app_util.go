// /////////////////////////////////////////////////////////////////////////////
// Application 一些辅助函数

package app

// 完成 app 的默认设置
func defaultConfiguration(app *Application) error {
	// 解析命令行参数
	parseArgs()

	// 有启动参数 -- 则启动服务器

	// 无启动参数 -- 则读取配置表

	// 加载 world.ini 信息

	// 设置开发环境

	// 加载 server.json 信息

	// 保存 cmd 命令参数
}

// 设置默认组件
func setDefaultComponent(app *Application) {
	// 根据服务器类型，注册默认组件
}

// 根据服务器类型 启动服务器
func runByType(app *Application) {
	// 服务器类型

	// 本机启动

	// ssh 启动
}

// 解析 命令行参数
func parseArgs() {
	// 参数定义
	env := flag.String("env", "production", "development type")             // 服务器运行环境 production= 开发环境 development = 运营环境
	serverType := flag.String("type", "serverType", "server type")          // 服务器类型，例如 gate connect area ...等
	name := flag.String("name", "master", "server name")                    // 服务器名字
	host := flag.String("host", "127.0.0.1", "server host")                 // 服务器IP地址
	port := flag.Uint("port", 0, "server port")                             // 服务器端口
	frontend := flag.Bool("frontend", false, "is frontend server")          // 是否是前端服务器
	clientHost := flag.String("clientHost", "127.0.0.1", "for client host") // 面向客户端的 IP地址
	clientPort := flag.Uint("clientPort", 0, "for client port")             // 面向客户端的 端口

	// 解析参数
	flag.Parse()
}
