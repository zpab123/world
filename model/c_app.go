// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- app 包

package model

// /////////////////////////////////////////////////////////////////////////////
// app 包

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
