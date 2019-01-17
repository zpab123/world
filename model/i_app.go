// /////////////////////////////////////////////////////////////////////////////
// 全局基础接口 -- app 包

package model

// App
type IApplication interface {
	GetAppDelegate() IAppDelegate // 获取 appDelegate
}

// App 代理
type IAppDelegate interface {
	ICilentPktHandler // 接口继承： 客户端 packet 消息处理
}
