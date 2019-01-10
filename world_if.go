// /////////////////////////////////////////////////////////////////////////////
// world 接口

package world

import (
	"github.com/zpab123/world/network/connector" // 网络连接库
)

// /////////////////////////////////////////////////////////////////////////////
// app 接口

// app 接口
type IApplication interface {
	IConnectorOpt // 接口继承： 设置 app 的 Connector 组件参数
}

// 设置 app 的 Connector 组件参数
type IConnectorOpt interface {
	GetConnectorOpt() *connector.ConnectorOpt     // 获取 connecot 配置参数
	SetConnectorOpt(opts *connector.ConnectorOpt) // 设置 connecot 配置参数
}
