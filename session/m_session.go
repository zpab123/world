// /////////////////////////////////////////////////////////////////////////////
// 模型 -- session 包

package session

import (
	"time"

	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

const (
	C_SESSION_FLUSH_INTERVAL = 5 * time.Millisecond // socket 数据刷新周期
)

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
	OnClientMessage(ses *FrontendSession, packet *network.Packet) // 收到1个新的客户端消息
}

// 服务端消息管理
type IServerMsgHandler interface {
	OnServerMessage(ses *BackendSession, packet *network.Packet) // 收到1个新的服务器消息
}

// 消息管理
type IMsgHandler interface {
	IClientMsgHandler // 客户端消息管理
	IServerMsgHandler // 服务端消息管理
}

// /////////////////////////////////////////////////////////////////////////////
// TFrontendSessionOpt 对象

// session 配置参数
type TFrontendSessionOpt struct {
	FlushInterval time.Duration          // socket 数据发送周期
	MsgHandler    IClientMsgHandler      // 消息处理对象
	WorldConnOpt  *network.TWorldConnOpt // WorldConnection 配置参数
}

// 创建1个新的 TFrontendSessionOpt
func NewTFrontendSessionOpt(handler IClientMsgHandler) *TFrontendSessionOpt {
	// 创建 WorldConnection
	wc := network.NewTWorldConnOpt()

	// 创建 TFrontendSessionOpt
	opts := &TFrontendSessionOpt{
		FlushInterval: C_SESSION_FLUSH_INTERVAL,
		MsgHandler:    handler,
		WorldConnOpt:  wc,
	}

	return opts
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TFrontendSessionOpt) Check() error {
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// TBackendSessionOpt 对象

// BackendSession 配置参数
type TBackendSessionOpt struct {
	FlushInterval    time.Duration          // socket 数据发送周期
	ServerMsgHandler IServerMsgHandler      // 消息处理对象
	WorldConnOpts    *network.TWorldConnOpt // WorldConnection 配置参数
}

// 创建1个新的 TFrontendSessionOpt
func NewTBackendSessionOpt(handler IServerMsgHandler) *TBackendSessionOpt {
	// 创建 WorldConnection
	wc := network.NewTWorldConnOpt()

	// 创建 TFrontendSessionOpt
	opts := &TBackendSessionOpt{
		FlushInterval:    C_SESSION_FLUSH_INTERVAL,
		ServerMsgHandler: handler,
		WorldConnOpts:    wc,
	}

	return opts
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TBackendSessionOpt) Check() error {
	return nil
}
