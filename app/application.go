// /////////////////////////////////////////////////////////////////////////////
// 1个通用服务器对象

package app

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/zpab123/syncutil"        // 同步工具
	"github.com/zpab123/world/component" // 组件库
	"github.com/zpab123/world/consts"    // 全局常量
	"github.com/zpab123/world/model"     // 全局 struct
	"github.com/zpab123/zplog"           // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// Application 对象

// 1个通用服务器对象
type Application struct {
	baseInfo     *model.BaseInfo                 // 基础属性
	runer        string                          // 服务器启动者 (master=master 命令启动 cmd=cmd 启动)
	serverInfo   model.ServerInfo                // 服务器配置信息
	state        syncutil.AtomicUint32           // app 当前状态
	componentMap map[string]component.IComponent // 名字-> 组件 集合
	runTime      time.Time                       // 启动时间
}

// 创建1个新的 Application 对象
func NewApplication() *Application {
	// 创建对象
	app := &Application{
		baseInfo:     &model.BaseInfo{},
		serverInfo:   model.ServerInfo{},
		componentMap: map[string]component.IComponent{},
	}

	// 设置为无效状态
	app.state.Store(consts.APP_STATE_INVALID)

	return app
}

// 初始化 Application
func (app *Application) Init() bool {
	// 错误变量
	//var initErr error = nil

	// 路径信息
	dir, err := getMainPath()
	if err != nil {
		zplog.Error("app 初始化失败->读取根目录失败")
		return false
	}
	app.baseInfo.MainPath = dir

	// 设置基础配置
	defaultConfiguration(app)

	// 设置为初始化状态
	app.state.Store(consts.APP_STATE_INIT)
	zplog.Infof("app 初始化完成")

	return true
}

// 启动 app
func (app *Application) Run() {
	// 记录启动时间
	app.runTime = time.Now()

	// 状态效验
	if app.state.Load() != consts.APP_STATE_INIT {
		//err := new error
		zplog.Error("app 非 init 状态，启动失败")
		return
	}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 主循环
	for {

	}
}

// 停止 app
func (app *Application) Stop() error {

	return nil
}

// 设置 app 名字
func (app *Application) SetType(v string) {
	app.baseInfo.AppType = v
}

// 注册组件
//
// com=符合 IComponent 接口的对象
func (app *Application) RegisterComponent(com component.IComponent) {
	// 获取名字
	name := com.Name()

	// 组件已经存在
	if app.componentMap[name] != nil {
		zplog.Warnf("组件[*s]重复注册，新组件覆盖旧组件", name)
	}

	// 保存组件
	app.componentMap[name] = com
}

// /////////////////////////////////////////////////////////////////////////////
// 包私有 api

// 获取 当前 Application 运行的绝对路径 例如：E:\code\go\go-project\src\test
func getMainPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		zplog.Warnf("获取 App 绝对路径失败")
		return "", err
	}
	//strings.Replace(dir, "\\", "/", -1)
	return dir, nil
}
