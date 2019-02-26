// /////////////////////////////////////////////////////////////////////////////
// 1个通用服务器对象

package app

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zpab123/world/config" // 配置文件库
	"github.com/zpab123/world/state"  // 状态管理
	"github.com/zpab123/world/utils"  // 工具库
	"github.com/zpab123/zaplog"       // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// Application 对象

// 1个通用服务器对象
type Application struct {
	baseInfo     *TBaseInfo          // 基础信息
	stateMgr     *state.StateManager // 状态管理
	componentMgr *ComponentManager   // 组件管理
	serverInfo   *config.TServerInfo // 配置信息
	appDelegate  IAppDelegate        // 代理对象
	signalChan   chan os.Signal      // 操作系统信号
}

// 创建1个新的 Application 对象
//
// appType=server.json 中配置的类型
func NewApplication(appType string, delegate IAppDelegate) *Application {
	// 参数验证
	if nil == delegate {
		zaplog.Error("Application 创建失败。 delegate=nil")

		os.Exit(1)
	}

	if "" == appType {
		zaplog.Error("Application 创建失败。 appType为空")

		os.Exit(1)
	}

	// 创建对象
	base := &TBaseInfo{}
	st := state.NewStateManager()
	cptMgr := NewComponentManager()
	signal := make(chan os.Signal, 1)

	// 创建 app
	app := &Application{
		baseInfo:     base,
		stateMgr:     st,
		componentMgr: cptMgr,
		appDelegate:  delegate,
		signalChan:   signal,
	}

	// 设置类型
	app.baseInfo.AppType = appType

	// 设置为无效状态
	app.stateMgr.SetState(state.C_STATE_INVALID)

	// 通知代理
	app.appDelegate.OnCreat(app)

	return app
}

// 初始化 Application
func (this *Application) Init() {
	st := this.stateMgr.GetState()
	if st != state.C_STATE_INVALID {
		zaplog.Fatal("app Init 失败，状态错误。当前状态=%d，正确状态=%d", st, state.C_STATE_INVALID)

		os.Exit(1)
	}

	// 获取主程序路径
	dir, err := utils.GetMainPath()
	if err != nil {
		zaplog.Fatal("app Init 失败。读取 main 根目录失败")

		os.Exit(1)
	}
	this.baseInfo.MainPath = dir

	// 默认设置
	defaultConfig(this)

	// 通知代理
	this.appDelegate.OnInit(this)

	// 状态： 初始化
	this.stateMgr.SetState(state.C_STATE_INIT)

	zaplog.Debugf("app 状态：init完成 ...")
}

// 启动 app
func (this *Application) Run() {
	// 状态：启动中
	if !this.stateMgr.SwapState(state.C_STATE_INIT, state.C_STATE_RUNING) {
		zaplog.Fatalf("app 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INIT, this.stateMgr.GetState())

		os.Exit(1)
	} else {
		zaplog.Infof("app 状态：正在启动中 ...")
	}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 记录启动时间
	this.baseInfo.RunTime = time.Now()

	// 创建组件
	createComponent(this)

	// 启动所有组件
	for _, cpt := range this.componentMgr.componentMap {
		go cpt.Run()
	}

	// 操作系统信号
	this.listenSignal()

	// 状态：工作中
	if !this.stateMgr.SwapState(state.C_STATE_RUNING, state.C_STATE_WORKING) {
		zaplog.Errorf("app 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_RUNING, this.stateMgr.GetState())

		os.Exit(1)
	} else {
		zaplog.Infof("app 状态：启动成功，工作中 ...")
	}

	// 主循环
	this.appDelegate.OnRun(this)
}

// 停止 app
func (this *Application) Stop() error {
	// 停止所有组件
	for _, cpt := range this.componentMgr.componentMap {
		cpt.Stop()
	}

	return nil
}

// 获取服务器信息
func (this *Application) GetServerInfo() *config.TServerInfo {
	return this.serverInfo
}

// 获取组件管理对象
func (this *Application) GetComponentMgr() *ComponentManager {
	return this.componentMgr
}

// 侦听操作系统信号
func (this *Application) listenSignal() {
	// 排除信号
	signal.Ignore(syscall.Signal(10), syscall.Signal(12), syscall.SIGPIPE, syscall.SIGHUP)
	signal.Notify(this.signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			sig := <-this.signalChan
			if syscall.SIGINT == sig || syscall.SIGTERM == sig {
				// this.Stop()

				zaplog.Infof("%s 服务器，优雅地退出", this.baseInfo.Name)

				os.Exit(0)
			} else {
				zaplog.Errorf("异常的操作系统信号=%s", sig)
			}
		}
	}()
}
