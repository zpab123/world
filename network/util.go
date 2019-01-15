// /////////////////////////////////////////////////////////////////////////////
// network 通用工具

package network

import (
	"io"

	"github.com/zpab123/world/model"          // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/world/network/packet" // tcp 消息包
	"github.com/zpab123/world/utils"          // 工具
)

// 发送Type-Length-Value格式的封包流程
func SendTLVPacket(writer io.Writer, ctx model.IContextSet, pkt *packet.Packet) error {
	// 将数据写入Socket
	err := utils.WriteAll(writer, pkt.Data())
}
