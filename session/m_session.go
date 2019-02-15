// /////////////////////////////////////////////////////////////////////////////
// 模型 -- session 包

package session

import (
	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// /////////////////////////////////////////////////////////////////////////////
// 接口

// session 接口
type ISession interface {
	Run()
	Close()
	GetId() int64
	SetId(v int64)
}

// session 管理
type ISessionManage interface {
	OnNewSession(ses ISession)   // 添加1个新的 session
	OnSessionClose(ses ISession) // 某个 session 关闭
}

// 客户端消息管理
type IClientMsgHandler interface {
	OnClientMessage(ses *FrontendSession, msg *Message) // 收到1个新的客户端消息
}

// 服务端消息管理
type IServerMsgHandler interface {
	OnServerMessage(ses *BackendSession, msg *Message) // 收到1个新的服务器消息
}

// /////////////////////////////////////////////////////////////////////////////
// TSessionOpts 对象

// session 配置参数
type TSessionOpts struct {
	MsgHandler    IMsgHandler             // 消息处理对象
	WorldConnOpts *network.TWorldConnOpts // WorldConnection 配置参数
}

// 创建1个新的 TSessionOpts
func NewTSessionOpts(handler IMsgHandler) *TSessionOpts {
	// 创建 WorldConnection
	wc := network.NewTWorldConnOpts()

	// 创建 TSessionOpts
	opts := &TSessionOpts{
		MsgHandler:    handler,
		WorldConnOpts: wc,
	}

	return opts
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TSessionOpts) Check() error {
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// TBackendSessionOpt 对象

// BackendSession 配置参数
type TBackendSessionOpt struct {
	ServerMsgHandler IServerMsgHandler       // 消息处理对象
	WorldConnOpts    *network.TWorldConnOpts // WorldConnection 配置参数
}

// 创建1个新的 TSessionOpts
func NewTBackendSessionOpt(handler IServerMsgHandler) *TBackendSessionOpt {
	// 创建 WorldConnection
	wc := network.NewTWorldConnOpts()

	// 创建 TSessionOpts
	opts := &TBackendSessionOpt{
		ServerMsgHandler: handler,
		WorldConnOpts:    wc,
	}

	return opts
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TBackendSessionOpt) Check() error {
	return nil
}
