// /////////////////////////////////////////////////////////////////////////////
// connector 运行状态操作

package connector

import (
	"sync"
	"sync/atomic"

	"github.com/zpab123/syncutil" // 原子变量工具
)

// /////////////////////////////////////////////////////////////////////////////
// State 对象

// 运行状态
type State struct {
	running    syncutil.AtomicBool // 是否正在运行状态
	stopWaitor sync.WaitGroup      // 结束线程组
	stopping   syncutil.AtomicBool // 是否正在停止
}

// 是否正在运行状态
func (self *State) IsRuning() bool {
	return self.running.Load()
}

// 设置运行状态
func (self *State) SetRunning(v bool) {
	self.running.Store(v)
}

// 等待所有线程结束
func (self *State) WaitAllStop() {
	// 如果正在停止时, 等待停止完成
	self.stopWaitor.Wait()
}

// 是否处于停止状态
func (self *State) IsStopping() bool {
	return self.stopping.Load()
}

// 开始停止
func (self *State) StartStop() {
	self.stopWaitor.Add(1)
	self.stopping.Store(true)
}

// 结束停止
func (self *State) EndStop() {
	if self.IsRuning() {
		self.stopWaitor.Done()
		self.running.Store(false)
	}
}
