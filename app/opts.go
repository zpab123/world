// /////////////////////////////////////////////////////////////////////////////
// app 各种组件配置参数

package app

import (
	"time"

	"github.com/zpab123/world/network/connector" // 网络连接库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// ComponentOpt 对象

// 配置 app 各种组件参数
type ComponentOpt struct {
	connectorOpt *connector.ConnectorOpt // connector 组件配置参数
}

// 获取 connector 组件参数 [IConnectorOpt] 接口
func (this *ComponentOpt) GetConnectorOpt() *connector.ConnectorOpt {
	return this.connectorOpt
}

// 设置 connector 组件参数 [IConnectorOpt] 接口
func (this *ComponentOpt) SetConnectorOpt(opts *connector.ConnectorOpt) {
	this.connectorOpt = opts
}
