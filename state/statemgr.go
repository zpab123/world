// /////////////////////////////////////////////////////////////////////////////
// 状态组件

package state

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

import (
	"sync"

	"github.com/zpab123/syncutil" // 原子变量工具
)

// /////////////////////////////////////////////////////////////////////////////
// StateManager 对象

// 状态管理，可以安全地被多个线程访问
type StateManager struct {
	state     syncutil.AtomicUint32 // app 当前状态
	runGroup  sync.WaitGroup        // 启动线程同步数组
	stopGroup sync.WaitGroup        // 结束线程同步数组
}

// 新建1个 StateManager 对象
func NewStateManager() *StateManager {
	sm := &StateManager{}

	return sm
}

// 设置状态
func (this *StateManager) SetState(v uint32) {
	this.state.Store(v)
}

// 获取状态
func (this *StateManager) GetState() uint32 {
	return this.state.Load()
}

// 添加 delta 个启动线程
func (this *StateManager) AddRunGo(delta int) {
	this.runGroup.Add(delta)
}

// 某个启动线程任务完成
func (this *StateManager) RunDone() {
	this.runGroup.Done()
}

// 等待启动任务完成 -- 阻塞线程
func (this *StateManager) RunWait() {
	this.runGroup.Wait()
}

// 添加 delta 个结束线程
func (this *StateManager) AddStopGo(delta int) {
	this.stopGroup.Add(delta)
}

// 某个结束线程任务完成
func (this *StateManager) StopDone() {
	this.stopGroup.Done()
}

// 等待结束任务完成 -- 阻塞线程
func (this *StateManager) StopWait() {
	this.stopGroup.Wait()
}
