// /////////////////////////////////////////////////////////////////////////////
// 全局基础接口 -- app 包

package model

// 组件基础
type IComponent interface {
	Name() // 获取组件名字
	Run()  // 组件开始运行
	Stop() // 组件停止运行
}

// App
type IApplication interface {
	GetAppDelegate() IAppDelegate // 获取 appDelegate
}

// App 代理
type IAppDelegate interface {
	ICilentPktHandler // 接口继承： 客户端 packet 消息处理
}
