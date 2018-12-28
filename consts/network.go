// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- network 包

// 全局 接口 定义
package consts

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// network 包

const (
	TCP_SERVER_RECONNECT_TIME = 3 * time.Second // tcp 网络服务 开启失败后，重新开启时间，单位秒
)
