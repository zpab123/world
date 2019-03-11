// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package connector

import (
	"time"

	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// /////////////////////////////////////////////////////////////////////////////
// TConnectorOpts 对象

// Connector 配置参数
type TConnectorOpts struct {
	TcpConnOpt   *model.TTcpConnOpt     // tcpSocket 配置参数
	ConnectorOpt *network.TConnectorOpt // Connector 配置参数
}

// 创建1个新的 TConnectorOpts
func NewTConnectorOpts() *TConnectorOpts {
	// 创建组合对象
	tcpOpt := model.NewTTcpConnOpt()
	ctOpt := network.NewTConnectorOpt()

	// 创建对象
	opts := &TDispatcherClientOpt{
		TcpConnOpt:   tcpOpt,
		ConnectorOpt: ctOpt,
	}

	return opts
}
