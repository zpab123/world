// /////////////////////////////////////////////////////////////////////////////
// connector 注册中心

package connector

import (
	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	creatorMap = map[string]CreateFunc{} // name->CreateFunc 集合
)

// /////////////////////////////////////////////////////////////////////////////
// public api

func RegisterCreator(f CreateFunc) {
	// 临时实例化一个，获取类型
	cntor := f()

	// 已经存在
	if _, ok := creatorMap[cntor.TypeName()]; ok {
		panic("注册 connector 重复，类型=%s", cntor.TypeName())
	}

	// 保存类型
	creatorMap[cntor.TypeName()] = f
}

// /////////////////////////////////////////////////////////////////////////////
// CreateFunc 对象

type CreateFunc func() worldnet.IConnector