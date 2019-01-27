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
	C_SES_STATE_CLOSING               // 正在关闭中
	C_SES_STATE_CLOSED                // 关闭完成
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// Session 组件
type ISession interface {
	GetId() int64  // 获取 session ID
	SetId(v int64) // 设置 session ID
	Close()        // 关闭 session
}

// session 管理
type ISessionManage interface {
	OnNewSession(ses ISession)   // 添加1个新的 session
	OnSessionClose(ses ISession) // 某个 session 关闭
}

// /////////////////////////////////////////////////////////////////////////////
// TSessionOpts 对象

// session 配置参数
type TSessionOpts struct {
	SessionType   string          // session 类型
	WorldConnOpts *TWorldConnOpts // WorldConnection 配置参数
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
