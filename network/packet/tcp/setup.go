// /////////////////////////////////////////////////////////////////////////////
// tcp 数据处理器注册

package tcp

import (
	"github.com/zpab123/world/consts"         // 全局常量
	"github.com/zpab123/world/model"          // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/world/network/packet" // 数据包管理
	"github.com/zpab123/world/worldnet"       // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// 包 初始化函数
func init() {
	packet.RegDataManager(model.C_PACKET_TYPE_TCP_TLV, creator)
}

// /////////////////////////////////////////////////////////////////////////////
// private api

// 数据处理创建函数
func creator(pm model.IPacketManager, handler model.IPacketHandler) {
	// 设置 DataManager
	dm := NewTcpDataManager(handler)
	pm.SetDataManager(dm)

	// 设置 Hooker

	// 设置回调函数
}
