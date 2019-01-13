// /////////////////////////////////////////////////////////////////////////////
// app 组件管理

package app

import (
	"github.com/zpab123/world/model" // 全局[常量-数据类型-接口]汇总
	"github.com/zpab123/zplog"       // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// ComponentManager 对象

// app 组件管理
type componentManager struct {
	connectorOpt *model.ConnectorOpt         // connector 组件配置参数
	componentMap map[string]model.IComponent // 名字-> 组件 集合
}

// 获取 connector 组件参数 [IComponentManager] 接口
func (this *componentManager) GetConnectorOpt() *model.ConnectorOpt {
	return this.connectorOpt
}

// 设置 connector 组件参数 [IComponentManager] 接口
func (this *componentManager) SetConnectorOpt(opts *model.ConnectorOpt) {
	this.connectorOpt = opts
}

// 注册1个 Component 组件
//
// com=符合 IComponent 接口的对象
func (this *componentManager) RegisterComponent(com model.IComponent) {
	// 获取名字
	name := com.Name()

	// 组件已经存在
	if this.componentMap[name] != nil {
		zplog.Warnf("组件[*s]重复注册，新组件将覆盖旧组件", name)
	}

	// 保存组件
	this.componentMap[name] = com
}

// 初始化 componentManager
func (this *componentManager) componentMgrInit() {
	// 创建 map
	this.componentMap = map[string]model.IComponent{}
}
