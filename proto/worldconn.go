// /////////////////////////////////////////////////////////////////////////////
// 对 PacketSocket 的封装，定义一些 world 内部常用的消息

package proto

import (
	"encoding/binary"
	"net"
	"sync"
	"time"

	"github.com/zpab123/world/model"          // 全局模型
	"github.com/zpab123/world/network"        // 网络库
	"github.com/zpab123/world/network/packet" // 消息包
	"github.com/zpab123/world/queue"          // 消息队列
	"github.com/zpab123/world/utils"          // 工具库
	"github.com/zpab123/zplog"                // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// WorldConnection 对象

// world 框架内部需要用到的一些常用网络消息
type WorldConnection struct {
	opts         *model.TWorldSocketOpts // 配置参数
	packetSocket model.IPacketSocket     // 接口继承： 符合 IPacketSocket 的对象
}

// 新建1个 WorldConnection 对象
func NewWorldConnection(opt *model.TWorldSocketOpts) *WorldConnection {
	// 创建对象
	wc := &WorldConnection{}

	return wc
}
