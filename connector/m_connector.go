// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package connector

import (
	"time"

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
	AcceptorName string                // 接收器名字
	MaxConn      uint32                // 最大连接数量，超过此数值后，不再接收新连接
	TcpConnOpts  *model.TTcpConnOpts   // tcpSocket 配置参数
	SessiobOpts  *session.TSessionOpts // session 配置参数
}

// 创建1个新的 TConnectorOpt
func NewTConnectorOpt(handler session.IMsgHandler) *TConnectorOpt {
	// 创建 tcp 配置参数
	tcpOpts := model.NewTTcpConnOpts()

	// 创建 session 配置参数
	sesOpts := session.NewTSessionOpts(handler)

	// 创建对象
	opts := &TConnectorOpt{
		AcceptorName: network.C_ACCEPTOR_NAME_COM,
		MaxConn:      MAX_CONN,
		TcpConnOpts:  tcpOpts,
		SessiobOpts:  sesOpts,
	}

	return opts
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TConnectorOpt) Check() error {
	if this.MaxConn <= 0 {
		this.MaxConn = MAX_CONN
	}

	return nil
}
