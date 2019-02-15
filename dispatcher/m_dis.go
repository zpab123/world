// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package dispatcher

import (
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/session" // session 库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

const (
	MAX_CONN              = 10000              // DispatcherServer 默认最大连接数
	COMPONENT_NAME_SERVER = "dispatcherServer" // 服务器组件名字
	COMPONENT_NAME_CLIENT = "dispatcherClient" // 客户端组件名字
)

// /////////////////////////////////////////////////////////////////////////////
// TDispatcherServerOpt 对象

// DispatcherServer 组件配置参数
type TDispatcherServerOpt struct {
	MaxConn     uint32                       // 最大连接数量，超过此数值后，不再接收新连接
	TcpConnOpts *model.TTcpConnOpt           // tcpSocket 配置参数
	SessionOpts *session.TBackendSessionOpts // session 配置参数
}

// 创建1个新的 TDispatcherServerOpts
func NewTDispatcherServerOpt() *TDispatcherServerOpt {
	//
	tcpOpt := model.NewTTcpConnOpt()
	bsOpt := session.NewTBackendSessionOpt()

	opt := TDispatcherServerOpts{
		MaxConn:     MAX_CONN,
		TcpConnOpts: tcpOpt,
		SessionOpts: bsOpt,
	}

	return opt
}

// /////////////////////////////////////////////////////////////////////////////
// TDispatcherServerOpts 对象
