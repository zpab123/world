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
	MAX_CONN = 10000 // DispatcherServer 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// TDispatcherServerOpts 对象

// DispatcherServer 组件配置参数
type TDispatcherServerOpts struct {
	MaxConn     uint32                       // 最大连接数量，超过此数值后，不再接收新连接
	TcpConnOpts *model.TTcpConnOpts          // tcpSocket 配置参数
	SessiobOpts *session.TBackendSessionOpts // session 配置参数
}

// 创建1个新的 TDispatcherServerOpts
func NewTDispatcherServerOpts() *TDispatcherServerOpts {
	// 创建 tcp 配置参数
	tcpOpts := model.NewTTcpConnOpts()

	// 创建对象
	opt := TDispatcherServerOpts{
		MaxConn:     MAX_CONN,
		TcpConnOpts: tcpOpts,
	}

	return opt
}

// /////////////////////////////////////////////////////////////////////////////
// TDispatcherServerOpts 对象
