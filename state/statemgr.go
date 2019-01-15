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
// stateManager 对象

// app 状态管理
type stateManager struct {
	state        syncutil.AtomicUint32 // app 当前状态
	goGroup      sync.WaitGroup        // 线程同步组
	recoverPanic bool                  // 是否进行 Panic 捕获
}

// 设置状态
func (this *stateManager) SetState(v uint32) {
	this.state.Store(v)
}

// 获取状态
func (this *stateManager) GetState() uint32 {
	return this.state.Load()
}

// 添加 delta 个线程
func (this *stateManager) AddGo(delta int) {
	this.goGroup.Add(delta)
}

// 某个线程任务完成
func (this *stateManager) Done() {
	this.goGroup.Done()
}

// 等待当前任务完成 -- 阻塞线程
func (this *stateManager) Wait() {
	this.goGroup.Wait()
}
