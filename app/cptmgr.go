// /////////////////////////////////////////////////////////////////////////////
// app 组件管理

package app

import (
	"github.com/zpab123/world/connector"  // connector 组件
	"github.com/zpab123/world/dispatcher" // dispatcher 组件
	"github.com/zpab123/world/model"      // 全局模型
	"github.com/zpab123/zaplog"           // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// ComponentManager 对象

// app 组件管理
type ComponentManager struct {
	componentMap map[string]model.IComponent      // 名字-> 组件 集合
	connectorOpt *connector.TConnectorOpt         // connector 组件配置参数
	disClientOpt *dispatcher.TDispatcherClientOpt // DispatcherClient 组件配置参数
}

// 新建1个 ComponentManager
func NewComponentManager() *ComponentManager {
	// 组件
	cptMgr := &ComponentManager{
		componentMap: map[string]model.IComponent{},
	}

	// 返回
	return cptMgr
}

// 添加1个 Component 组件
//
// com=符合 IComponent 接口的对象
func (this *ComponentManager) AddComponent(com model.IComponent) {
	// 获取名字
	name := com.Name()

	// 组件已经存在
	if _, ok := this.componentMap[name]; ok {
		zaplog.Warnf("组件[*s]重复添加，新组件将覆盖旧组件", name)
	}

	// 保存组件
	this.componentMap[name] = com
}

// 根据名字获取组件
func (this *ComponentManager) GetComponentByName(name string) model.IComponent {
	cpt, ok := this.componentMap[name]

	if ok {
		return cpt
	} else {
		return nil
	}
}

// 获取 connector 组件参数
func (this *ComponentManager) GetConnectorOpt() *connector.TConnectorOpt {
	return this.connectorOpt
}

// 设置 connector 组件参数
func (this *ComponentManager) SetConnectorOpt(opt *connector.TConnectorOpt) {
	this.connectorOpt = opt
}

// 获取 DispatcherClient 组件参数
func (this *ComponentManager) GetDisClientOpt() *dispatcher.TDispatcherClientOpt {
	return this.disClientOpt
}

// 设置 DispatcherClient 组件参数
func (this *ComponentManager) SetDisClientOpt(opt *dispatcher.TDispatcherClientOpt) {
	this.disClientOpt = opt
}
