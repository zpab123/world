// /////////////////////////////////////////////////////////////////////////////
// 1个通用服务器对象

package app

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/zplog"       // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// Application 对象

// 1个通用服务器对象
type Application struct {
	baseInfo         *model.TBaseInfo      // 服务器基础信息
	serverInfo       *model.TServerInfo    // 服务器配置信息
	state            syncutil.AtomicUint32 // app 当前状态
	goGroup          sync.WaitGroup        // 线程同步组
	componentManager                       // 对象继承： app 组件管理
	appDelegate      model.IAppDelegate    // app 代理对象
}

// 创建1个新的 Application 对象
//
// appType=server.json 中配置的类型
func NewApplication(appType string, appDelegate model.IAppDelegate) *Application {
	// 创建对象
	app := &Application{
		baseInfo:    &model.TBaseInfo{},
		serverInfo:  &model.TServerConfig{},
		appDelegate: appDelegate,
	}

	// 设置类型
	app.baseInfo.AppType = appType

	// 设置为无效状态
	app.state.Store(model.C_APP_STATE_INVALID)

	return app
}

// 初始化 Application
func (this *Application) Init() bool {
	// 错误变量
	//var initErr error = nil

	// 路径信息
	dir, err := getMainPath()
	if err != nil {
		zplog.Error("app 初始化失败：读取根目录失败")

		return false
	}
	this.baseInfo.MainPath = dir

	// 组件管理初始化
	this.componentMgrInit()

	// 设置基础配置
	defaultConfiguration(this)

	// 设置为初始化状态
	this.state.Store(model.C_APP_STATE_INIT)
	zplog.Infof("app 状态：初始化完成")

	return true
}

// 启动 app
func (this *Application) Run() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 记录启动时间
	this.baseInfo.RunTime = time.Now()

	// 状态效验
	if this.state.Load() != model.C_APP_STATE_INIT {
		//err := new error
		zplog.Error("app 非 init 状态，启动失败")

		return
	}

	// 设置默认组件
	setDefaultComponent(this)

	// 设置为启动中
	this.state.Store(model.C_APP_STATE_RUNING)
	zplog.Infof("app 状态：正在启动 ...")

	// 启动所有组件
	for _, cpt := range this.componentMap {
		this.goGroup.Add(1)
		go cpt.Run()
	}

	// 阻塞 - 等待启动完成
	this.goGroup.Wait()

	// 启动完毕 - 设置为工作中
	this.state.Store(model.C_APP_STATE_WORKING)
	zplog.Infof("app 状态：启动成功，工作中 ...")

	// 主循环
	for {

	}
}

// 停止 this
func (this *Application) Stop() error {
	// 停止所有组件
	for _, cpt := range this.componentMap {
		cpt.Stop()
	}

	return nil
}

// 设置 this 类型
func (this *Application) SetType(v string) {
	this.baseInfo.ServerType = v
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

// 获取 appDelegate
func (this *Application) GetAppDelegate() model.IAppDelegate {
	return this.appDelegate
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
