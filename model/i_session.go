// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- session 包

package model

// Session 组件
type ISession interface {
	SendMessage(msg interface{})
	IState                  // 接口继承： 状态接口
	GetOpts() *TSessionOpts // 获取配置参数
}

// 客户端 packet 处理
type ICilentPktHandler interface {
	ISessionManage
	OnClientPkt(ses ISession, pkt interface{}) // 收到客户端 packet 消息包
}

// session 管理
type ISessionManage interface {
	OnNewSession(ISession)   // 1个新的 session 创建成功
	OnSessionClose(ISession) // 某个 session 关闭
}
