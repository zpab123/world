// /////////////////////////////////////////////////////////////////////////////
// acceptor 运行状态操作 [代码完整]

package acceptor

import (
	"sync"
	"sync/atomic"

	"github.com/zpab123/syncutil" // 原子变量工具
)

// /////////////////////////////////////////////////////////////////////////////
// State 对象

// 运行状态
type State struct {
	running       syncutil.AtomicBool // 是否正在运行状态
	stopWaitGroup sync.WaitGroup      // 结束线程组
	stopping      syncutil.AtomicBool // 是否正在停止
}

// 是否正在运行状态
func (this *State) IsRuning() bool {
	return this.running.Load()
}

// 设置运行状态
func (this *State) SetRunning(v bool) {
	this.running.Store(v)
}

// 等待所有线程结束
func (this *State) WaitAllStop() {
	// 如果正在停止时, 等待停止完成
	this.stopWaitGroup.Wait()
}

// 是否处于停止状态
func (this *State) IsStopping() bool {
	return this.stopping.Load()
}

// 开始停止
func (this *State) StartStop() {
	this.stopWaitGroup.Add(1)
	this.stopping.Store(true)
}

// 停止结束
func (this *State) EndStop() {
	if this.IsRuning() {
		this.stopWaitGroup.Done()
		this.stopping.Store(false)
	}
}
