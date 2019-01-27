// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- 通用接口

package model

// /////////////////////////////////////////////////////////////////////////////
// 常量

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

// 状态通用
const (
	C_STATE_INVALID  uint32 = iota // 无效状态
	C_STATE_INIT                   // 初始化状态
	C_STATE_RUNING                 // 正在启动中
	C_STATE_WORKING                // 工作状态
	C_STATE_CLOSEING               // 正在关闭中
	C_STATE_CLOSED                 // 关闭完成
	C_STATE_STOPING                // 正在停止中
	C_STATE_STOP                   // 停止完成
)

// /////////////////////////////////////////////////////////////////////////////
// 接口

// 设置和获取自定义属性
type IContextSet interface {
	// 为对象设置一个自定义属性
	SetContext(key interface{}, v interface{})

	// 从对象上根据key获取一个自定义属性
	GetContext(key interface{}) (interface{}, bool)

	// 给定一个值指针, 自动根据值的类型GetContext后设置到值
	FetchContext(key, valuePtr interface{}) bool
}
