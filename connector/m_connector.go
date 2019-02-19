// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package connector

import (
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络模型
	"github.com/zpab123/world/session" // session 库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// 组件名字
const (
	COMPONENT_NAME = "connector" // 组件名字
)

// 最大连接数
const (
	MAX_CONN = 100000 // connector 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// /////////////////////////////////////////////////////////////////////////////
// TConnectorOpt 对象

// connector 组件配置参数
type TConnectorOpt struct {
	AcceptorName       string                       // 接收器名字
	MaxConn            uint32                       // 最大连接数量，超过此数值后，不再接收新连接
	Frontend           bool                         // 是否面向前端
	TcpConnOpt         *model.TTcpConnOpt           // tcpSocket 配置参数
	FrontendSessionOpt *session.TFrontendSessionOpt // FrontendSession 配置参数
	BackendSessionOpt  *session.TBackendSessionOpt  // BackendSession 配置参数
}

// 创建1个新的 TConnectorOpt
func NewTConnectorOpt(handler session.IMsgHandler) *TConnectorOpt {
	// 创建对象
	tcpOpt := model.NewTTcpConnOpt()
	fSesOpt := session.NewTFrontendSessionOpt(handler)
	bSesOpt := session.NewTBackendSessionOpt(handler)

	// 创建 TConnectorOpt
	opts := &TConnectorOpt{
		AcceptorName:       network.C_ACCEPTOR_NAME_COM,
		MaxConn:            MAX_CONN,
		Frontend:           true,
		TcpConnOpt:         tcpOpt,
		FrontendSessionOpt: fSesOpt,
		BackendSessionOpt:  bSesOpt,
	}

	return opts
}
