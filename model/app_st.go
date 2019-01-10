// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- app 包

package model

// /////////////////////////////////////////////////////////////////////////////
// Application 基础属性

// Application 基础属性
type BaseInfo struct {
	Env      string // Application 当前运行环境 production= 开发环境 development = 运营环境
	AppType  string // 当前app的类型
	Name     string // 当前app的名字
	MainPath string // 程序根路径 例如 "E/server/gateserver.exe" 所在的目录 “E/server/”
}
