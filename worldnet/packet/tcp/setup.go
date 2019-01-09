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
	packet.RegisterPostter(consts.WORLDNET_PKT_TYPE_LTV, creator)
}

// /////////////////////////////////////////////////////////////////////////////
// private api

// 绑定函数
func creator(pm packet.IPacketManager) {
	// 设置 PacketPostter
	pm.SetPacketPostter(new(TcpPacketPostter))

	// 设置 Hooker

	// 设置回调函数
}
