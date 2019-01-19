// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- 通用接口

package model

// 状态管理
type IState interface {
	SetState(v uint32) // 设置状态
	GetState() uint32  // 获取状态
}

// 多线程同步
type ISyncGroup interface {
	Add(delta int) // 添加 delta 个 go 线程
	Done()         // 完成1个 go  线程
	Wait()         // 阻塞： 等待所有 go 线程结束
}

// 多线程任务状态同步
type IStateGroup interface {
	IState     // 接口继承： 状态基础
	ISyncGroup // 接口继承： 多线程同步
}

// 设置和获取自定义属性
type IContextSet interface {
	// 为对象设置一个自定义属性
	SetContext(key interface{}, v interface{})

	// 从对象上根据key获取一个自定义属性
	GetContext(key interface{}) (interface{}, bool)

	// 给定一个值指针, 自动根据值的类型GetContext后设置到值
	FetchContext(key, valuePtr interface{}) bool
}
