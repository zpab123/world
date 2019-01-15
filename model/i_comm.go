// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- 通用接口

package model

// /////////////////////////////////////////////////////////////////////////////
// 组件相关

// 组件基础
type IComponent interface {
	Name() string // 获取组件名字
	Run()         // 组件开始运行
	Stop()        // 组件停止运行
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
