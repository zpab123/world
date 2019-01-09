// /////////////////////////////////////////////////////////////////////////////
// connector 注册中心

package connector

import (
	"fmt"

	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/zplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	creatorMap = map[string]CreateFunc{} // typeName->CreateFunc 集合
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 注册1个 connector 创建函数
func RegisterCreator(f CreateFunc) {
	// 临时实例化一个，获取类型
	cntor := f()

	// 已经存在
	if _, ok := creatorMap[cntor.GetType()]; ok {
		panic("注册 connector 重复，类型=%s", cntor.GetType())
	}

	// 保存类型
	creatorMap[cntor.GetType()] = f
}

// 根据类型，创建1个 connector 对象
func NewConnector(typeName string, addr *Address) network.IConnector {
	// 类型检查
	creator := creatorMap[typeName]
	if nil == creator {
		zplog.Panicf("创建 connector 出错：找不到 %s 类型的 connector", typeName)
		panic(fmt.Sprintf("创建 connector 出错：找不到 %s 类型的 connector"), typeName)
	}

	// 地址检查
	_, err := addr.Check()
	if nil != err {
		zplog.Panicf(err.Error())
		panic(err.Error())
	}

	// 创建 connector
	cntor := creator()

	// 设置基本参数
	cntor.SetAddress(addr.GetAddrRange())

	return cntor
}

// /////////////////////////////////////////////////////////////////////////////
// CreateFunc 对象

type CreateFunc func() network.IConnector
