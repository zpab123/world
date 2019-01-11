// /////////////////////////////////////////////////////////////////////////////
// 1个通用服务器对象

package app

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/zpab123/world/base"   // 基础信息
	"github.com/zpab123/world/consts" // 全局常量
	"github.com/zpab123/world/model"  // 全局 struct
	"github.com/zpab123/zplog"        // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// Application 对象

// 1个通用服务器对象
type Application struct {
	baseInfo         *base.BaseInfo   // 服务器基础信息
	serverInfo       model.ServerInfo // 服务器配置信息
	stateManager                      // 对象继承： app 状态管理
	componentManager                  // 对象继承： app 组件管理
}

// 创建1个新的 Application 对象
//
// appType=server.json 中配置的类型
func NewApplication(appType string) *Application {
	// 创建对象
	app := &Application{
		baseInfo:   &base.BaseInfo{},
		serverInfo: model.ServerInfo{},
	}

	// 设置类型
	app.baseInfo.ServerType = appType

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
		zplog.Error("app 初始化失败：读取根目录失败")

		return false
	}
	app.baseInfo.MainPath = dir

	// 组件管理初始化
	app.componentMgrInit()

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
	app.baseInfo.RunTime = time.Now()

	// 状态效验
	if app.state.Load() != consts.APP_STATE_INIT {
		//err := new error
		zplog.Error("app 非 init 状态，启动失败")
		return
	}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 注册默认组件
	regDefaultComponent(app)

	// 启动所有组件
	for _, cpt := range app.componentMap {
		go cpt.Run()
	}

	// 设置为启动中
	app.state.Store(consts.APP_STATE_RUNING)

	// 主循环
	for {

	}
}

// 停止 app
func (app *Application) Stop() error {

	return nil
}

// 设置 app 类型
func (app *Application) SetType(v string) {
	app.baseInfo.ServerType = v
}

// 获取 tcp 服务器监听地址(格式 -> 127.0.0.1:6532)
//
// 如果不存在，则返回 ""
func (app *Application) GetCTcpAddr() string {
	// tcp 地址
	var cTcpAddr string = ""
	if app.serverInfo.CTcpPort > 0 {
		cTcpAddr = fmt.Sprintf("%s:%d", app.serverInfo.ClientHost, app.serverInfo.CTcpPort) // 面向客户端的 tcp 地址
	}

	return cTcpAddr
}

// 获取 websocket 服务器监听地址(格式 -> 127.0.0.1:6532)
//
// 如果不存在，则返回 ""
func (app *Application) GetCWsAddr() string {
	// websocket 地址
	var cWsAddr string = ""
	if app.serverInfo.CWsPort > 0 {
		cWsAddr = fmt.Sprintf("%s:%d", app.serverInfo.ClientHost, app.serverInfo.CWsPort) // 面向客户端的 websocket 地址
	}

	return cWsAddr
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
