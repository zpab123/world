// /////////////////////////////////////////////////////////////////////////////
// 管道 [代码完整]

package network

import (
	"sync"
)

// /////////////////////////////////////////////////////////////////////////////
// Pipe 对象

// 通用队列对象
//
// 不限制大小，添加不发生阻塞，接收阻塞等待
type Pipe struct {
	list      []interface{} // 消息数组
	listMutex sync.Mutex    // 互斥锁
	listCond  *sync.Cond    // 条件变量
}

// 新建1个 Pipe 对象
func NewPipe() *Pipe {
	// 创建对象
	p := &Pipe{}
	p.listCond = sync.NewCond(&p.listMutex)

	return p
}

// 添加1个消息
//
// 添加时不会发送阻塞
func (self *Pipe) Add(msg interface{}) {
	// 添加消息
	self.listMutex.Lock()
	self.list = append(self.list, msg)
	self.listMutex.Unlock()

	// 发送信号
	self.listCond.Signal()
}

// 初始化 list
func (self *Pipe) Reset() {
	self.list = self.list[0:0]
}

// 如果没有数据，发生阻塞
//
// list 中如果出现 nil 数据， exit = true
func (self *Pipe) Pick(retList *[]interface{}) (exit bool) {
	// 没有数据 - 阻塞
	self.listMutex.Lock()
	for len(self.list) == 0 {
		self.listCond.Wait()
	}
	self.listMutex.Unlock()

	// 复制数据
	self.listMutex.Lock()
	for _, data := range self.list {
		if nil == data {
			exit = true
			break
		} else {
			*retList = append(*retList, data)
		}
	}

	// 初始化 list
	self.Reset()
	self.listMutex.Unlock()

	return
}
