// /////////////////////////////////////////////////////////////////////////////
// 场景服务器

package scene

import (
	"math/rand"
	"os"
	"time"

	"github.com/zpab123/world/config" // 配置工具库
	"github.com/zpab123/world/model"  // 全局模型
	"github.com/zpab123/world/state"  // 状态管理
	"github.com/zpab123/world/utils"  // 工具库
	"github.com/zpab123/zaplog"       // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	scene *Scene // Scene 实例
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 启动场景服务器
func Run() {
	// 初始化
	if nil == scene {
		scene = NewScene()
	}
	scene.Init()

	// 启动
	scene.Run()
}

// /////////////////////////////////////////////////////////////////////////////
// Scene

// 场景服务
type Scene struct {
	stateMgr     *state.StateManager // 状态管理
	baseInfo     *model.TBaseInfo    // 服务器启动基本信息
	serverInfo   *config.TServerInfo // 配置信息
	componentMgr *ComponentManager   // 组件管理
}

// 新建1个 Scene 对象
func NewScene() *Scene {
	// 创建状态管理
	st := state.NewStateManager()

	// 创建组件管理
	cptMgr := NewComponentManager()

	// 创建 Scene
	scene := &Scene{
		stateMgr:     st,
		baseInfo:     &model.TBaseInfo{},
		componentMgr: cptMgr,
	}

	// 设置类型
	scene.baseInfo.Type = C_SERVER_TYPE

	// 设置为无效状态
	scene.stateMgr.SetState(state.C_STATE_INVALID)

	return scene
}

// 场景初始化
func (this *Scene) Init() {
	// 获取主程序路径
	dir, err := utils.GetMainPath()
	if err != nil {
		zaplog.Error("Scene Init 失败。读取 main 根目录失败")

		os.Exit(1)
	}
	this.baseInfo.MainPath = dir

	// 设置默认参数
	defaultConfig(this)

	// 创建组件
	//createComponent(this)

	// 改变为初始化状态
	if !this.stateMgr.SwapState(state.C_STATE_INVALID, state.C_STATE_INIT) {
		zaplog.Errorf("Scene Init失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INVALID, this.stateMgr.GetState())

		os.Exit(1)
	}

	zaplog.Infof("scene 状态：init完成 ...")
}

// 启动 Scene
func (this *Scene) Run() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 记录启动时间
	this.baseInfo.RunTime = time.Now()

	// 改变状态为：启动中
	if !this.stateMgr.SwapState(state.C_STATE_INIT, state.C_STATE_RUNING) {
		zaplog.Errorf("Scene 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INIT, this.stateMgr.GetState())

		os.Exit(1)
	}

	// 启动所有组件
	for _, cpt := range this.componentMgr.componentMap {
		go cpt.Run()
	}

	// 消息分发

	// 改变为工作中
	if !this.stateMgr.SwapState(state.C_STATE_RUNING, state.C_STATE_WORKING) {
		zaplog.Errorf("Scene 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_RUNING, this.stateMgr.GetState())

		os.Exit(1)
	} else {
		zaplog.Infof("Scene 状态：启动成功，工作中 ...")
	}

	// 主循环
	select {}
}
