// /////////////////////////////////////////////////////////////////////////////
// 全局基础模型 -- session 包

package model

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// Session 类型
const (
	C_SES_TYPE_CLINET = "client" // 客户端类型的 session
)

// Session 状态
const (
	C_SES_STATE_INITED  uint32 = iota // 初始化状态
	C_SES_STATE_RUNING                // 正在运行
	C_SES_STATE_CLOSING               // 关闭中
	C_SES_STATE_CLOSED                // 关闭完成
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

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

// /////////////////////////////////////////////////////////////////////////////
// TSessionOpts 对象

// session 配置参数
type TSessionOpts struct {
	SessionType    string          // session 类型
	SessionManager ISessionManage  // session 管理对象
	WorldConnOpts  *TWorldConnOpts // WorldConnection 配置参数
}

// 创建1个新的 TSessionOpts
func NewTSessionOpts() *TSessionOpts {
	// 创建 WorldConnection
	wc := NewTWorldConnOpts()

	// 创建 TSessionOpts
	opts := &TSessionOpts{
		WorldConnOpts: wc,
	}

	return opts
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TSessionOpts) Check() error {
	return nil
}
