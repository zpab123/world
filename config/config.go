// /////////////////////////////////////////////////////////////////////////////
// 配置文件读取工具

package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/zpab123/world/consts" // 全局常量
	"github.com/zpab123/zplog"        // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// 变量
var (
	configMutex      sync.Mutex              // 进程互斥锁
	mainPath         string                  // 程序启动目录
	iniConfig        *WorldIni               // world.ini 配置信息
	iniFilePath      = consts.PATH_WORLD_INI // world.ini 默认路径
	serverConfig     *ServerConfig           // server.json 配置表
	serverConfigPath = consts.PATH_SERVER    // servers.json 配置文件路径
)

// 初始化
func init() {
	// 获取路径
	mainPath = getMainPath()

	// 读取 world.ini 配置
	getWorldIni()

	// 读取 server.json 配置
	getServerConfig()
}

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// /////////////////////////////////////////////////////////////////////////////
// 私有 api

// 获取 当前 main 程序运行的绝对路径 例如：E:\code\go\go-project\src\test
func getMainPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		zplog.Error("获取 main 程序 绝对路径失败")
		return ""
	}
	//strings.Replace(dir, "\\", "/", -1)
	return dir
}

// 获取 zpworld.ini 配置对象
func getWorldIni() {
	// 锁住线程
	configMutex.Lock()

	// retun 后，解锁
	defer configMutex.Unlock()

	// 读取文件
	if nil == iniConfig {
		// 获取 main 程序路径
		if "" == mainPath {
			zplog.Fatal("读取 main 程序路径失败")
			return
		}

		// 读取 ini 文件
		iniConfig = readWorldIni()
	}
}

// 获取 servers.json 配置信息
func getServerConfig() {
	// 锁住线程
	configMutex.Lock()

	// retun 后，解锁
	defer configMutex.Unlock()

	// 读取文件
	if nil == serverConfig {
		// 获取 main 路径
		if "" == mainPath {
			zplog.Fatal("读取 main 路径失败")
			return
		}

		// 创建对象
		serverConfig = &ServerConfig{
			Development: map[string][]ServerInfo{},
			Production:  map[string][]ServerInfo{},
		}

		// 加载文件
		fPath := filepath.Join(mainPath, serverConfigPath)
		LoadJsonToMap(fPath, serverConfig)
	}
}
