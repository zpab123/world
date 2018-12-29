// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- app 包

// 全局 接口 定义
package consts

// /////////////////////////////////////////////////////////////////////////////
// app 包

// APP 状态
const (
	APP_STATE_INVALID = iota // 无效状态
	APP_STATE_INIT           // 初始状态
	APP_STATE_READY          // 准备状态
	APP_STATE_RUN            // 运行状态
	APP_STATE_STOP           // 停止状态
)

// 启动方式
const (
	APP_RUNER_CMD    = "cmd"    // 由 cmd 启动
	APP_RUNER_MASTER = "master" // 由 master 服务器启动
)
