// /////////////////////////////////////////////////////////////////////////////
// packet 处理器 注册中心

package packet

import (
	"fmt"

	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/zplog"       // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 变量

// 变量
var (
	handler    model.IPacketHandler                 // 消息处理器，每个应用只能有1个消息处理器
	dataMgrMap = map[string]model.TDataMgrCreator{} // dataType->TDataMgrCreator 集合
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 注册个1个 dataManager 对象
func RegDataManager(dataType string, f model.TDataMgrCreator) {
	dataMgrMap[dataType] = f
}

// 注册个1个 handler 处理对象
func RegHandler(pm model.IPacketManager, pktType string, handler model.IPacketHandler) {
	// 创建数据管理器
	if dmf, ok := dataMgrMap[pktType]; ok {
		dmf(pm)
	} else {
		zplog.Panicf("注册 Handler 失败。不存在 %s 类型的数据管理器", pktType)
		panic(fmt.Sprintf("注册 Handler 失败。不存在 %s 类型的数据管理器", pktType))
	}
}
