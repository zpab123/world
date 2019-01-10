// /////////////////////////////////////////////////////////////////////////////
// 配置文件读取工具

package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/zpab123/world/consts" // 全局常量
	"github.com/zpab123/world/model"  // 全局 struct
	"github.com/zpab123/zplog"        // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// 变量
var (
	configMutex      sync.Mutex              // 进程互斥锁
	mainPath         string                  // 程序启动目录
	iniConfig        *model.WorldIni         // world.ini 配置信息
	iniFilePath      = consts.PATH_WORLD_INI // world.ini 默认路径
	serverConfig     *model.ServerConfig     // server.json 配置表
	serverConfigPath = consts.PATH_SERVER    // servers.json 配置文件路径
	serverMap        model.ServerMap         // servers.json 中// 服务器 type -> *[]ServerInfo 信息集合
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

// 获取 world.ini 配置对象
func GetWorldIni() *model.WorldIni {
	return iniConfig
}

// 获取 servers.json 配置信息
//
// 返回： map[string][]ServerInfo 数据集合
func GetServerConfig() *model.ServerConfig {
	return serverConfig
}

// 获取 当前环境的 服务器信息集合
func GetServerMap() model.ServerMap {
	return serverMap
}

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
		serverConfig = &model.ServerConfig{
			Development: map[string][]model.ServerInfo{},
			Production:  map[string][]model.ServerInfo{},
		}

		// 加载文件
		fPath := filepath.Join(mainPath, serverConfigPath)
		LoadJsonToMap(fPath, serverConfig)

		// 根据运行环境赋值
		if consts.ENV_DEV == iniConfig.Env {
			serverMap = serverConfig.Development
		} else {
			serverMap = serverConfig.Production
		}
	}
}

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// 配置信息
type Config struct {
	inicfg  *WorldIni    // world.ini 配置信息
	jsoncfg ServerConfig // server.json 配置信息
}

// 获取当前开发环境下的服务器集合
func (this *Config) GetServerMap() ServerMap {
	// 获取环境
	env := this.inicfg.Env

	// 获取 serverMap
	var sMap ServerMap
	if env == consts.ENV_PRO {
		sMap = this.jsoncfg.Production
	} else {
		sMap = this.jsoncfg.Development
	}

	return sMap
}

// 根据服务器类型，获取当前环境下服务配置列表
func (this *Config) GetServerLsit(serverType string) ServetList {
	// 获取 serverMap
	sMap := this.GetServerMap()

	// 获取服务器列表
	return sMap[serverType]
}

// 获取某台服务器的具体配置信息
func (this *Config) GetServerInfo(gid uint16) ServerInfo {
	// 获取 app 类型
	var serverType string

	// 获取服务器列表
	list := this.GetServerLsit(serverType)

	// 进程 id 错误
	if gid >= len(list) {
		return nil
	}

	// 返回服务器信息
	return list[gid]
}

// 获取 tcp 监听地址
func (this *Config) GetCTcpAddr() string {
	// tcp 地址
	var cTcpAddr string

	//if this

	return cTcpAddr
}
