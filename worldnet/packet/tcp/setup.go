// /////////////////////////////////////////////////////////////////////////////
// tcp 数据处理器注册

package tcp

import (
	"github.com/zpab123/world/consts"          // 全局常量
	"github.com/zpab123/world/worldnet"        // 网络库
	"github.com/zpab123/world/worldnet/packet" // 数据包管理
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

// 包 初始化函数
func init() {
	packet.RegisterPktMgr(consts.WORLDNET_PKT_TYPE_LTV, creator)
}

// /////////////////////////////////////////////////////////////////////////////
// private api

// 绑定函数
func creator(dm packet.IDataManger) {
	// 设置 PacketManger
	dm.SetPacketManager(new(TcpPacketManager))

	// 设置 Hooker

	// 设置回调函数
}
