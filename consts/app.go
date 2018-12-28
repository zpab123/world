// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- app 包

// 全局 接口 定义
package consts

// /////////////////////////////////////////////////////////////////////////////
// app 包

const (
	APP_STATE_INIT  = 1 + iota // 初始状态
	APP_STATE_READY            // 准备状态
	APP_STATE_RUN              // 运行状态
	APP_STATE_STOP             // 停止状态
)
