// /////////////////////////////////////////////////////////////////////////////
// world 管理对象

package world

import (
	"github.com/zpab123/world/app" // 1个通用服务器对象
)

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// 创建1个新的 Application 对象
//
// appType=server.json 中配置的类型
func CreateApp(appType string, appDelegate app.IAppDelegate) *app.Application {
	// 创建 app
	app := app.NewApplication(appType, appDelegate)
	app.Init()

	return app
}
