// /////////////////////////////////////////////////////////////////////////////
// 配置文件读取工具

package config

import (
	"path/filepath"
	"sync"

	"github.com/zpab123/world/utils" // 工具库
	"github.com/zpab123/zaplog"      // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// 变量
var (
	configMutex      sync.Mutex                        // 进程互斥锁
	mainPath         string                            // 程序启动目录
	iniFilePath                     = C_PATH_WORLD_INI // world.ini 默认路径
	serverConfigPath                = C_PATH_SERVER    // servers.json 配置文件路径
	worldConfig      *TWorld        = &TWorld{}        // world 引擎配置信息
	serverConfig     *TServerConfig                    // server.json 配置表
	serverMap        TServerMap                        // servers.json 中// 服务器 type -> *[]ServerInfo 信息集合
)

// 初始化
func init() {
	// 获取路径
	dir, err := utils.GetMainPath()
	if nil == err {
		mainPath = dir
	}

	// 读取 world.ini 配置
	readWorldIni()

	// 读取 servers.json 配置信息
	readServerJson()
}

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// 获取 world.ini 配置对象
func GetWorldConfig() *TWorld {
	return worldConfig
}

// 获取 servers.json 配置信息
//
// 返回： map[string][]ServerInfo 数据集合
func GetServerConfig() *TServerConfig {
	return serverConfig
}

// 获取 当前环境的 服务器信息集合
func GetServerMap() TServerMap {
	return serverMap
}

// /////////////////////////////////////////////////////////////////////////////
// 私有 api

// 读取 servers.json 配置信息
func readServerJson() {
	// 锁住线程
	configMutex.Lock()

	// retun 后，解锁
	defer configMutex.Unlock()

	// 读取文件
	if nil == serverConfig {
		// 获取 main 路径
		if "" == mainPath {
			zaplog.Fatal("读取 main 路径失败")
			return
		}

		// 创建对象
		serverConfig = &TServerConfig{
			Development: TServerMap{},
			Production:  TServerMap{},
		}

		// 加载文件
		fPath := filepath.Join(mainPath, serverConfigPath)
		LoadJsonToMap(fPath, serverConfig)

		// 根据运行环境赋值
		if C_ENV_DEV == worldConfig.Env {
			serverMap = serverConfig.Development
		} else {
			serverMap = serverConfig.Production
		}
	}
}
