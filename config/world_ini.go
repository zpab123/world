// /////////////////////////////////////////////////////////////////////////////
// world.ini 配置文件

package config

// World 引擎配置
type TWorld struct {
	Env      string // 当前运行环境，production= 开发环境；development = 运营环境
	LogLevel string // log 输出等级
}
