// /////////////////////////////////////////////////////////////////////////////
// 1个通用服务器对象

package app

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/zpab123/world/config" // 配置文件库
	"github.com/zpab123/world/state"  // 状态管理
	"github.com/zpab123/zplog"        // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// Application 对象

// 1个通用服务器对象
type Application struct {
	stateMgr     *state.StateManager // 状态管理
	componentMgr *ComponentManager   // app 组件管理
	baseInfo     *TBaseInfo          // 服务器基础信息
	serverInfo   *config.TServerInfo // 服务器配置信息
	appDelegate  IAppDelegate        // app 代理对象
}

// 创建1个新的 Application 对象
//
// appType=server.json 中配置的类型
func NewApplication(appType string, delegate IAppDelegate) *Application {
	// 创建状态管理
	st := state.NewStateManager()

	// 创建组件管理
	cptMgr := NewComponentManager()

	// 创建对象
	app := &Application{
		stateMgr:     st,
		componentMgr: cptMgr,
		baseInfo:     &TBaseInfo{},
		serverInfo:   &config.TServerInfo{},
		appDelegate:  delegate,
	}

	// 设置类型
	app.baseInfo.AppType = appType

	// 设置为无效状态
	app.stateMgr.SetState(state.C_STATE_INVALID)

	return app
}

// 初始化 Application
func (this *Application) Init() bool {
	// 获取主程序路径
	dir, err := getMainPath()
	if err != nil {
		zplog.Error("app Init 失败。读取根目录失败")

		return false
	}
	this.baseInfo.MainPath = dir

	// 默认设置
	defaultConfiguration(this)

	// 改变为初始化状态
	if !this.stateMgr.SwapState(state.C_STATE_INVALID, state.C_STATE_INIT) {
		zplog.Errorf("app Init失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INVALID, this.stateMgr.GetState())

		return false
	}

	zplog.Infof("app 状态：init完成 ...")

	return true
}

// 启动 app
func (this *Application) Run() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 记录启动时间
	this.baseInfo.RunTime = time.Now()

	// 改变为启动中
	if !this.stateMgr.SwapState(state.C_STATE_INIT, state.C_STATE_RUNING) {
		zplog.Errorf("app 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INIT, this.stateMgr.GetState())

		return
	} else {
		zplog.Infof("app 状态：正在启动中 ...")
	}

	// 设置默认组件
	setDefaultComponent(this)

	// 启动所有组件
	for _, cpt := range this.componentMgr.componentMap {
		go cpt.Run()
	}

	// 改变为工作中
	if !this.stateMgr.SwapState(state.C_STATE_RUNING, state.C_STATE_WORKING) {
		zplog.Errorf("app 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_RUNING, this.stateMgr.GetState())

		return
	} else {
		zplog.Infof("app 状态：启动成功，工作中 ...")
	}

	// 主循环
	for {

	}
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
