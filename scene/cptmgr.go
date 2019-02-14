// /////////////////////////////////////////////////////////////////////////////
// scene 组件管理

package scene

import (
	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/zaplog"      // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// /////////////////////////////////////////////////////////////////////////////
// ComponentManager 对象

// scene 组件管理
type ComponentManager struct {
	componentMap map[string]model.IComponent // 名字-> 组件 集合
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

// 注册1个 Component 组件
//
// com=符合 IComponent 接口的对象
func (this *ComponentManager) RegisterComponent(com model.IComponent) {
	// 获取名字
	name := com.Name()

	// 组件已经存在
	if this.componentMap[name] != nil {
		zaplog.Warnf("组件[*s]重复注册，新组件将覆盖旧组件", name)
	}

	// 保存组件
	this.componentMap[name] = com
}