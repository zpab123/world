// /////////////////////////////////////////////////////////////////////////////
// 场景服务器

package scene

import (
	"math/rand"
	"os"
	"time"

	"github.com/zpab123/world/state" // 状态管理
	"github.com/zpab123/zaplog"      // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	scene *Scene = NewScene() // Scene 实例
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 启动场景服务器
func Run() {

}

// /////////////////////////////////////////////////////////////////////////////
// Scene

// 场景服务
type Scene struct {
	stateMgr *state.StateManager // 状态管理
}

// 新建1个 Scene 对象
func NewScene() *Scene {
	// 创建状态管理
	st := state.NewStateManager()

	// 创建 Scene
	scene := &Scene{
		stateMgr: st,
	}

	scene.stateMgr.SetState(state.C_STATE_INIT)

	return scene
}

// 场景初始化
func (this *Scene) Init() {

}

// 启动 Scene
func (this *Scene) Run() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 改变状态为：启动中
	if !this.stateMgr.SwapState(state.C_STATE_INIT, state.C_STATE_RUNING) {
		zaplog.Errorf("Scene 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INIT, this.stateMgr.GetState())

		os.Exit(1)
	}

}
