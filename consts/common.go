// /////////////////////////////////////////////////////////////////////////////
// 全局常量 -- 通用

package consts

// /////////////////////////////////////////////////////////////////////////////
// 通用

// 运行环境
const (
	ENV_DEV = "development" // 默认：开发环境
	ENV_PRO = "production"  // 运营环境
)

// 配置文件路径
const (
	PATH_WORLD_INI = "/config/world.ini"    // world.ini 配置文件路径
	PATH_MASTER    = "/config/master.json"  // master 服务器配置文件路径
	PATH_SERVER    = "/config/servers.json" // servers 服务器配置文件路径
)

// 服务器类型
const (
	SERVER_TYPE_GATE      = "gate"      // gate 服务器类型
	SERVER_TYPE_CONNECTOR = "connector" // connector 服务器类型
)
