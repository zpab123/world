// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- 基础通用

package model

// 运行环境
const (
	C_ENV_DEV = "development" // 默认：开发环境
	C_ENV_PRO = "production"  // 运营环境
)

// 配置文件路径
const (
	C_PATH_WORLD_INI = "/config/world.ini"    // world.ini 配置文件路径
	C_PATH_MASTER    = "/config/master.json"  // master 服务器配置文件路径
	C_PATH_SERVER    = "/config/servers.json" // servers 服务器配置文件路径
)

// 服务器类型
const (
	C_SERVER_TYPE_MASTER    = "master"    // master 服务器
	C_SERVER_TYPE_GATE      = "gate"      // gate 服务器类型
	C_SERVER_TYPE_CONNECTOR = "connector" // connector 服务器类型
)
