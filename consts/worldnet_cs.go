// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- worldnet 包

// 全局 接口 定义
package consts

import (
	"time"
)

// packet 数据包读取方式
const (
	WORLDNET_PKT_TYPE_LTV = "ltv" // ltv(Length-Type-Value), Length为封包大小，Type为消息ID，Value为消息内容
)
