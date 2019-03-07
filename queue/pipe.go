// /////////////////////////////////////////////////////////////////////////////
// 管道 [代码完整]

package queue

import (
	"sync"
)

// /////////////////////////////////////////////////////////////////////////////
// Pipe 对象

// 通用队列对象，并发安全
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
func (this *Pipe) Add(msg interface{}) {
	// 添加消息
	this.listMutex.Lock()
	this.list = append(this.list, msg)
	this.listMutex.Unlock()

	// 发送信号
	this.listCond.Signal()
}

// 初始化 list
func (this *Pipe) Reset() {
	this.list = this.list[0:0]
}

// 如果没有数据，发生阻塞
//
// list 中如果出现 nil 数据， exit = true
func (this *Pipe) Pick(retList *[]interface{}) (exit bool) {
	// 没有数据 - 阻塞
	this.listMutex.Lock()
	for len(this.list) == 0 {
		this.listCond.Wait()
	}
	this.listMutex.Unlock()

	// 复制数据
	this.listMutex.Lock()
	for _, data := range this.list {
		if nil == data {
			exit = true
			break
		} else {
			*retList = append(*retList, data)
		}
	}

	// 初始化 list
	this.Reset()
	this.listMutex.Unlock()

	return
}
