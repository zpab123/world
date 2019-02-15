// /////////////////////////////////////////////////////////////////////////////
// 1个通用服务器对象

package app

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/zpab123/world/config"    // 配置文件库
	"github.com/zpab123/world/connector" // connector 组件
	"github.com/zpab123/world/state"     // 状态管理
	"github.com/zpab123/zaplog"          // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// Application 对象

// 1个通用服务器对象
type Application struct {
	stateMgr     *state.StateManager // 状态管理
	componentMgr *ComponentManager   // 组件管理
	baseInfo     *TBaseInfo          // 基础信息
	serverInfo   *config.TServerInfo // 配置信息
	appDelegate  IAppDelegate        // 代理对象
}

// 创建1个新的 Application 对象
//
// appType=server.json 中配置的类型
func NewApplication(appType string, delegate IAppDelegate) *Application {
	// 创建状态管理
	st := state.NewStateManager()

	// 创建组件管理
	cptMgr := NewComponentManager()

	// 创建 app
	app := &Application{
		stateMgr:     st,
		componentMgr: cptMgr,
		baseInfo:     &TBaseInfo{},
		appDelegate:  delegate,
	}

	// 设置类型
	app.baseInfo.AppType = appType

	// 设置为无效状态
	app.stateMgr.SetState(state.C_STATE_INVALID)

	return app
}

// 初始化 Application
func (this *Application) Init() {
	// 获取主程序路径
	dir, err := getMainPath()
	if err != nil {
		zaplog.Error("app Init 失败。读取 main 根目录失败")

		os.Exit(1)
	}
	this.baseInfo.MainPath = dir

	// 默认设置
	defaultConfig(this)

	// 创建组件
	createComponent(this)

	// 改变为初始化状态
	if !this.stateMgr.SwapState(state.C_STATE_INVALID, state.C_STATE_INIT) {
		zaplog.Errorf("app Init失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INVALID, this.stateMgr.GetState())

		os.Exit(1)
	}

	zaplog.Infof("app 状态：init完成 ...")
}

// 启动 app
func (this *Application) Run() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 记录启动时间
	this.baseInfo.RunTime = time.Now()

	// 改变状态为：启动中
	if !this.stateMgr.SwapState(state.C_STATE_INIT, state.C_STATE_RUNING) {
		zaplog.Errorf("app 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INIT, this.stateMgr.GetState())

		os.Exit(1)
	} else {
		zaplog.Infof("app 状态：正在启动中 ...")
	}

	// 启动所有组件
	for _, cpt := range this.componentMgr.componentMap {
		go cpt.Run()
	}

	// 结束信号侦听
	// setupSignals()

	// 启动 appDelegate
	go this.appDelegate.Run()

	// 改变为工作中
	if !this.stateMgr.SwapState(state.C_STATE_RUNING, state.C_STATE_WORKING) {
		zaplog.Errorf("app 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_RUNING, this.stateMgr.GetState())

		os.Exit(1)
	} else {
		zaplog.Infof("app 状态：启动成功，工作中 ...")
	}

	// 主循环
	select {}
}

// 停止 app
func (this *Application) Stop() error {
	// 停止所有组件
	for _, cpt := range this.componentMgr.componentMap {
		cpt.Stop()
	}

	return nil
}

// 获取 tcp 服务器监听地址(格式 -> 127.0.0.1:6532)
//
// 如果不存在，则返回 ""
func (this *Application) GetCTcpAddr() string {
	// tcp 地址
	var cTcpAddr string = ""
	if this.serverInfo.CTcpPort > 0 {
		cTcpAddr = fmt.Sprintf("%s:%d", this.serverInfo.ClientHost, this.serverInfo.CTcpPort) // 面向客户端的 tcp 地址
	}

	return cTcpAddr
}

// 获取 websocket 服务器监听地址(格式 -> 127.0.0.1:6532)
//
// 如果不存在，则返回 ""
func (this *Application) GetCWsAddr() string {
	// websocket 地址
	var cWsAddr string = ""
	if this.serverInfo.CWsPort > 0 {
		cWsAddr = fmt.Sprintf("%s:%d", this.serverInfo.ClientHost, this.serverInfo.CWsPort) // 面向客户端的 websocket 地址
	}

	return cWsAddr
}

// 获取服务器信息
func (this *Application) GetServerInfo() *config.TServerInfo {
	return this.serverInfo
}

// 获取组件管理对象
func (this *Application) GetComponentMgr() *ComponentManager {
	return this.componentMgr
}

// /////////////////////////////////////////////////////////////////////////////
// 包私有 api

// 获取 当前 Application 运行的绝对路径 例如：E:\code\go\go-project\src\test
func getMainPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		zaplog.Error("获取 App 绝对路径失败")

		return "", err
	}
	//strings.Replace(dir, "\\", "/", -1)
	return dir, nil
}
