// /////////////////////////////////////////////////////////////////////////////
// packet 接口汇总

package packet

import (
	"github.com/zpab123/world/worldnet" // 网络库
)

// packet 消息收发接口
type IDataManger interface {
	// 接收 1个 packet
	SetPacketManager(mgr worldnet.IPacketManager)
	// 设置 接收后，发送前的事件处理流程
	//SetHooker
	// 设置 接收后最终处理回调
	//SetCallback(v cellnet.EventCallback)
}
