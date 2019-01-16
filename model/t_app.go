// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- app 包

package model

// /////////////////////////////////////////////////////////////////////////////
// 服务器基础信息

// 启动信息
type TBaseInfo struct {
	ServerType string    // 服务器类型
	MainPath   string    // main 程序所在路径
	Env        string    // 运行环境 production= 开发环境 development = 运营环境
	Name       string    // 服务器名字
	RunTime    time.Time // 启动时间
}
