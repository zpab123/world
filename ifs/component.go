// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- 组件接口

package ifs

// component 组件
type IComponent interface {
	Name() string // 获取组件名字
	Run()         // 组件开始运行
	Stop()        // 组件停止运行
}
