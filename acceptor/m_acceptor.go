// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package acceptor

import (
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络模型
	"github.com/zpab123/world/session" // session 库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

const (
	C_COMPONENT_NAME = "acceptor" // 组件名字
	C_MAX_CONN       = 100000     // acceptor 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// /////////////////////////////////////////////////////////////////////////////
// TAcceptorOpt 对象

// Acceptor 组件配置参数
type TAcceptorOpt struct {
	Enable       bool                       // 是否启动 connector
	AcceptorName string                     // 接收器名字
	MaxConn      uint32                     // 最大连接数量，超过此数值后，不再接收新连接
	ForClient    bool                       // 是否面向客户端
	TcpConnOpt   *model.TTcpConnOpt         // tcpSocket 配置参数
	ClientSesOpt *session.TClientSessionOpt // ClientSession 配置参数
	ServerSesOpt *session.TServerSessionOpt // ServerSession 配置参数
}

// 创建1个新的 TAcceptorOpt
func NewTAcceptorOpt(handler session.IMsgHandler) *TAcceptorOpt {
	// 创建对象
	tcpOpt := model.NewTTcpConnOpt()

	csOpt := session.NewTClientSessionOpt(handler)
	ssOpt := session.NewTServerSessionOpt(handler)

	// 创建 TAcceptorOpt
	opt := &TAcceptorOpt{
		Enable:       true,
		AcceptorName: network.C_ACCEPTOR_NAME_COM,
		MaxConn:      C_MAX_CONN,
		ForClient:    true,
		TcpConnOpt:   tcpOpt,
		ClientSesOpt: csOpt,
		ServerSesOpt: ssOpt,
	}

	return opt
}
