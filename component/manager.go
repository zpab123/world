// /////////////////////////////////////////////////////////////////////////////
// 组件库

package component

import (
	"github.com/zpab123/zplog" // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// 包变量
var (
	componentMap map[string]IComponent = map[string]IComponent{} // 组件 名字-> 组件 集合
)

// /////////////////////////////////////////////////////////////////////////////
// 组件库

// 注册组件
//
// com=符合 IComponent 接口的对象
func RegisterComponent(com IComponent) {
	// 获取名字
	name := com.Name()

	// 组件已经存在
	component := componentMap[name]
	if component != nil {
		zplog.Warnf("组件[*s]重复注册，新组件覆盖旧组件", name)
	}

	// 保存组件
	componentMap[name] = com
}
