// /////////////////////////////////////////////////////////////////////////////
// ini 配置文件读取对象

package config

import (
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"           // ini 库
	"github.com/zpab123/world/consts" // 全局常量
	"github.com/zpab123/world/model"  // 全局 struct
	"github.com/zpab123/zplog"        // log 工具
)

// /////////////////////////////////////////////////////////////////////////////
// 私有 api

// 读取 world.ini 配置文件
func readWorldIni() *model.WorldIni {
	// 创建 WorldIni 对象
	config := &model.WorldIni{
		Env: consts.ENV_DEV, // 默认开发环境
	}

	// 读取配置文件
	fPath := filepath.Join(mainPath, iniFilePath)
	zplog.Infof("读取 world.ini 配置文件，路径=%s", fPath)
	iniFile, err := ini.Load(fPath)

	// 错误检查
	checkConfigError(err, "")

	// 获取配置
	for _, sec := range iniFile.Sections() {
		// 配置名字
		secName := sec.Name()
		secName = strings.ToLower(secName) // 转化成小写

		// 读取配置
		if secName == "env" {
			readEnv(sec, config) // 开发环境
		} else if "gate" == secName {
			//readGate(sec, config) // gate 服务器配置
		} else if "world" == secName {
			//readWorld(sec, worldConfig) // world 服务器配置
		}
	}

	return config
}

// 检查 ini 读取错误
func checkConfigError(err error, msg string) {
	if err != nil {
		if msg == "" {
			msg = err.Error()
		}
		zplog.Fatalf("在读取 world.ini 中出现错误 error: %s", msg)
	}
}

// 读取 env 开发环境
func readEnv(sec *ini.Section, config *model.WorldIni) {
	// 设置成默认
	config.Env = consts.ENV_DEV

	// 读取属性
	for _, key := range sec.Keys() {
		name := strings.ToLower(key.Name()) //转化成小写
		if name == "env" {
			config.Env = key.MustString(config.Env)
			if config.Env != consts.ENV_DEV && config.Env != consts.ENV_PRO {
				config.Env = consts.ENV_DEV
				zplog.Fatal("world.ini 中 [env] 参数配置错误，设置成默认 development")
			}
		} else {
			zplog.Fatal("读取 world.ini [env] 失败")
		}
	}
}
