// /////////////////////////////////////////////////////////////////////////////
// app 组件管理

package app

import (
	"github.com/zpab123/world/component"         // 组件库
	"github.com/zpab123/world/network/connector" // 网络连接库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

import (
	"github.com/zpab123/syncutil" // 原子变量工具
)

// /////////////////////////////////////////////////////////////////////////////
// ComponentManager 对象

// app 组件管理
type componentManager struct {
	connectorOpt *connector.ConnectorOpt         // connector 组件配置参数
	componentMap map[string]component.IComponent // 名字-> 组件 集合
}

// 获取 connector 组件参数 [IComponentManager] 接口
func (this *componentManager) GetConnectorOpt() *connector.ConnectorOpt {
	return this.connectorOpt
}

// 设置 connector 组件参数 [IComponentManager] 接口
func (this *componentManager) SetConnectorOpt(opts *connector.ConnectorOpt) {
	this.connectorOpt = opts
}

// 注册1个 Component 组件
//
// com=符合 IComponent 接口的对象
func (this *componentManager) RegisterComponent(com component.IComponent) {
	// 获取名字
	name := com.Name()

	// 组件已经存在
	if this.componentMap[name] != nil {
		zplog.Warnf("组件[*s]重复注册，新组件将覆盖旧组件", name)
	}

	// 保存组件
	this.componentMap[name] = com
}
