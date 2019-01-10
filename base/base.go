// /////////////////////////////////////////////////////////////////////////////
// 服务器公用的一些基础 信息

package base

import (
	"time"

	"github.com/zpab123/world/consts" // 全局常量
)

// /////////////////////////////////////////////////////////////////////////////
// Cmd 对象

// 启动信息
type BaseInfo struct {
	ServerType string    // 服务器类型
	MainPath   string    // main 程序所在路径
	Env        string    // 运行环境 production= 开发环境 development = 运营环境
	Name       string    // 服务器名字
	RunTime    time.Time // 启动时间
}
