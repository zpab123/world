// /////////////////////////////////////////////////////////////////////////////
// 包 通用 struct

// 全局 struct 数据定义
package app

// /////////////////////////////////////////////////////////////////////////////
// Application

// Application 基础属性
type BaseInfo struct {
	Env       string // Application 当前运行环境 production= 开发环境 development = 运营环境
	StartTime int64  // 当前服务器的启动时间
	State     int    // app当前状态
	AppName   string // 当前app的名字
	BasePath  string // 程序根路径 例如 "E/server/gateserver.exe" 所在的目录 “E/server/”
}
