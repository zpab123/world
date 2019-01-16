// /////////////////////////////////////////////////////////////////////////////
// 全局基础接口 -- app 包

package model

// App
type Application interface {
	GetAppDelegate() IAppDelegate // 获取 appDelegate
}

// App 代理
type IAppDelegate interface {
	OnClientMsg(ses ISession, msg interface{}) // 收到客户端消息
}
