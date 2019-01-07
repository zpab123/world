// /////////////////////////////////////////////////////////////////////////////
// connector 接口汇总

package connector

import (
	"github.com/zpab123/world/worldnet" // worldnet 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// connector 相关

// Session 管理接口
type ISessionManager interface {
	worldnet.ISessionAccessor // 接口继承： Session 存取器接口
	Add(worldnet.ISession)    // 添加1个符合 ISession 接口的对象
	Remove(worldnet.ISession) // 移除1个符合 ISession 接口的对象
	GetCount() int            // 获取当前 ISession 数量
	SetIDStart(start int64)   // 设置ID开始的号
}

// 数据处理接口
type IDataMananger interface {
	GetDataMananger() *DataManager // 获取 *DataManager 对象
}
