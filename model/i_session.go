// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- session 包

package model

// Session 组件
type ISession interface {
}

// 客户端 packet 处理
type ICilentPktHandler interface {
	OnClientPkt(ses ISession, pkt interface{}) // 收到客户端 packet 消息包
}
