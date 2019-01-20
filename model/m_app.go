// /////////////////////////////////////////////////////////////////////////////
// 全局模型 -- app 包

package model

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// APP 状态
const (
	C_APP_STATE_INVALID = iota // 无效状态
	C_APP_STATE_INIT           // 初始状态
	C_APP_STATE_RUNING         // 正在启动
	C_APP_STATE_WORKING        // 启动完成，运行状态
	C_APP_STATE_STOPING        // 正在停止
	C_APP_STATE_STOP           // 停止状态
)

// APP 组件名字
const (
	C_CPT_NAME_CONNECTOR = "connector.connector" // connector 组件
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// 组件基础
type IComponent interface {
	Name() string // 获取组件名字
	Run()         // 组件开始运行
	Stop()        // 组件停止运行
}

// App
type IApplication interface {
	GetAppDelegate() IAppDelegate // 获取 appDelegate
}

// App 代理
type IAppDelegate interface {
	//ICilentPktHandler // 接口继承： 客户端 packet 消息处理
}

// /////////////////////////////////////////////////////////////////////////////
// TBaseInfo 对象

// app 启动信息
type TBaseInfo struct {
	AppType  string    // 服务器类型
	MainPath string    // main 程序所在路径
	Env      string    // 运行环境 production= 开发环境 development = 运营环境
	Name     string    // 服务器名字
	RunTime  time.Time // 启动时间
}
