// /////////////////////////////////////////////////////////////////////////////
// world 管理对象

package world

import (
	"github.com/zpab123/world/app"    // 1个通用服务器对象
	"github.com/zpab123/world/entity" // 实体库
	"github.com/zpab123/world/scene"  // 场景服务器
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

// 启动场景服务器
func RunScene() {
	scene.Run()
}

// 注册1个 space 类型实体
//
// spacePtr=符合 entity.ISpace 的指针对象
func RegisterSpace(spacePtr entity.ISpace) {
	entity.RegisterSpace(spacePtr)
}
