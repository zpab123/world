// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package dispatcher

import (
	"time"

	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

const (
	C_MAX_CONN                  = 10000                // DispatcherServer 默认最大连接数
	C_COMPONENT_NAME_SERVER     = "dispatcherServer"   // 服务器组件名字
	C_COMPONENT_NAME_CLIENT     = "dispatcherClient"   // 客户端组件名字
	C_DISPATCHER_FLUSH_INTERVAL = 5 * time.Millisecond // socket 数据刷新周期
)

// /////////////////////////////////////////////////////////////////////////////
// TDispatcherServerOpt 对象

// DispatcherServer 组件配置参数
type TDispatcherServerOpt struct {
	AcceptorName string                      // 接收器名字
	MaxConn      uint32                      // 最大连接数量，超过此数值后，不再接收新连接
	TcpConnOpt   *model.TTcpConnOpt          // tcpSocket 配置参数
	SessionOpt   *session.TBackendSessionOpt // session 配置参数
}

// 创建1个新的 TDispatcherServerOpts
func NewTDispatcherServerOpt(handler session.IServerMsgHandler) *TDispatcherServerOpt {
	// 创建组合对象
	tcpOpt := model.NewTTcpConnOpt()
	bsOpt := session.NewTBackendSessionOpt(handler)

	// 创建对象
	opt := &TDispatcherServerOpt{
		AcceptorName: network.C_ACCEPTOR_TYPE_TCP,
		MaxConn:      C_MAX_CONN,
		TcpConnOpt:   tcpOpt,
		SessionOpt:   bsOpt,
	}

	return opt
}

// /////////////////////////////////////////////////////////////////////////////
// TDispatcherClientOpt 对象

// DispatcherClient 组件配置参数
type TDispatcherClientOpt struct {
	Enable         bool                     // 是否启用 DispatcherClient
	FlushInterval  time.Duration            // socket 数据发送周期
	TcpConnOpt     *model.TTcpConnOpt       // tcpSocket 配置参数
	WorldSocketOpt *network.TWorldSocketOpt // WorldSocketOpt 配置参数
}

// 创建1个新的 TDispatcherServerOpts
func NewTDispatcherClientOpt(handler session.IServerMsgHandler) *TDispatcherClientOpt {
	// 创建组合对象
	tcpOpt := model.NewTTcpConnOpt()
	wsOpt := network.NewTWorldSocketOpt()

	// 创建对象
	opt := &TDispatcherClientOpt{
		Enable:         false,
		FlushInterval:  C_DISPATCHER_FLUSH_INTERVAL,
		TcpConnOpt:     tcpOpt,
		WorldSocketOpt: wsOpt,
	}

	return opt
}
