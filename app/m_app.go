// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package app

import (
	"time"

	"github.com/zpab123/world/session" // session 库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// app 代理
const (
	CLIENT_PKT_QUEUE_SIZE = 10000 // 客户端消息队列默认长度
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// App 代理
type IAppDelegate interface {
	OnCreat(app *Application) // app 创建成功
	OnInit(app *Application)  // app 初始化成功
	OnRun(app *Application)   // app 开始运行
	OnStop(app *Application)  // app 停止运行
	session.IMsgHandler       // 接口继承：消息管理
}

// /////////////////////////////////////////////////////////////////////////////
// TBaseInfo 对象

// app 启动信息
type TBaseInfo struct {
	AppType  string    // App 类型
	MainPath string    // main 程序所在路径
	Env      string    // 运行环境 production= 开发环境 development = 运营环境
	Name     string    // App 名字
	RunTime  time.Time // 启动时间
}
