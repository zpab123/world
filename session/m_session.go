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
	C_SES_FLUSH_INTERVAL = 5 * time.Millisecond // socket 数据刷新周期
)

// session 状态
const (
	C_SES_STATE_INIT    uint32 = iota // 初始化状态
	C_SES_STATE_RUNING                // 正在启动中
	C_SES_STATE_WORKING               // 工作状态
	C_SES_STATE_STOPING               // 正在停止中
	C_SES_STATE_STOPED                // 停止完成
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// session 接口
type ISession interface {
	Run() error
	Stop() error
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
	OnClientMessage(ses *ClientSession, packet *network.Packet) // 收到1个新的客户端消息
}

// 服务端消息管理
type IServerMsgHandler interface {
	OnServerMessage(ses *ServerSession, packet *network.Packet) // 收到1个新的服务器消息
}

// 消息管理
type IMsgHandler interface {
	IClientMsgHandler // 客户端消息管理
	IServerMsgHandler // 服务端消息管理
}

// /////////////////////////////////////////////////////////////////////////////
// TClientSessionOpt 对象

// ClientSession 配置参数
type TClientSessionOpt struct {
	FlushInterval time.Duration          // socket 数据发送周期
	MsgHandler    IClientMsgHandler      // 消息处理对象
	WorldConnOpt  *network.TWorldConnOpt // WorldConnection 配置参数
}

// 创建1个新的 TFrontendSessionOpt
func NewTClientSessionOpt(handler IClientMsgHandler) *TClientSessionOpt {
	// 创建 WorldConnection
	wc := network.NewTWorldConnOpt()

	// 创建 TFrontendSessionOpt
	opt := &TClientSessionOpt{
		FlushInterval: C_SES_FLUSH_INTERVAL,
		MsgHandler:    handler,
		WorldConnOpt:  wc,
	}

	return opt
}

// /////////////////////////////////////////////////////////////////////////////
// TServerSessionOpt 对象

// ServerSession 配置参数
type TServerSessionOpt struct {
	FlushInterval    time.Duration          // socket 数据发送周期
	ServerMsgHandler IServerMsgHandler      // 消息处理对象
	WorldConnOpt     *network.TWorldConnOpt // WorldConnection 配置参数
}

// 创建1个新的 TServerSessionOpt
func NewTServerSessionOpt(handler IServerMsgHandler) *TServerSessionOpt {
	// 创建 WorldConnection
	wc := network.NewTWorldConnOpt()

	// 创建 TServerSessionOpt
	opt := &TServerSessionOpt{
		FlushInterval:    C_SES_FLUSH_INTERVAL,
		ServerMsgHandler: handler,
		WorldConnOpt:     wc,
	}

	return opt
}
