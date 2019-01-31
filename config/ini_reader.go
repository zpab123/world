// /////////////////////////////////////////////////////////////////////////////
// ini 配置文件读取对象

package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"     // ini 库
	"github.com/zpab123/zaplog" // log 工具
)

// /////////////////////////////////////////////////////////////////////////////
// 私有 api

// 读取 world.ini 配置文件
func readWorldIni() {
	// 读取配置文件
	fPath := filepath.Join(mainPath, iniFilePath)
	zaplog.Infof("读取 world.ini 配置文件，路径=%s", fPath)
	iniFile, err := ini.Load(fPath)

	// 错误检查
	checkConfigError(err, "")

	// 获取配置
	for _, sec := range iniFile.Sections() {
		// 配置名字
		secName := sec.Name()
		secName = strings.ToLower(secName) // 转化成小写

		// 读取配置
		if secName == "world" {
			readWorld(sec, worldConfig) // world
		} else if "network" == secName {
			readNetwork(sec, config) // 握手信息配置
		} else if "world" == secName {
			//readWorld(sec, worldConfig) // world 服务器配置
		}
	}
}

// 检查 ini 读取错误
func checkConfigError(err error, msg string) {
	if err != nil {
		if msg == "" {
			msg = err.Error()
		}

		zaplog.Fatalf("读取 world.ini 出现错误。 err=%s", msg)

		os.Exit(1)
	}
}

// 读取 world 配置数据
func readWorld(sec *ini.Section, config *TWorld) {
	// 设置默认
	config.Env = C_ENV_DEV

	// 读取属性
	for _, key := range sec.Keys() {
		name := strings.ToLower(key.Name()) //转化成小写
		if "env" == name {
			config.Env = key.MustString(config.Env)
			if config.Env != C_ENV_DEV && config.Env != C_ENV_PRO {
				config.Env = C_ENV_DEV
				zaplog.Fatal("world.ini 中 [world]-env 参数配置错误。修改为 development")
			}
		} else if "log_level" == name {
			config.LogLevel = key.MustString(config.LogLevel)
		}
	}
}

// 读取 network 配置信息
func readNetwork(sec *ini.Section, config *TWorldIni) {
	// 读取属性
	for _, key := range sec.Keys() {
		name := strings.ToLower(key.Name()) //转化成小写
		if "shake_key" == name {
			config.Key = key.MustString(config.Key)
		} else if "acceptor" == name {
			var aptor int = 0
			aptor = key.MustInt(aptor)
			config.Acceptor = uint32(aptor)
		}
	}
}
