// /////////////////////////////////////////////////////////////////////////////
// packet 接口汇总

package packet

import (
	"github.com/zpab123/world/worldnet" // 网络库
)

// packet 消息管理接口
type IPacketManager interface {
	// 设置 packet 接受/发送对象
	SetPacketPostter(mgr worldnet.IPacketPostter)
	// 设置 接收后，发送前的事件处理流程
	//SetHooker
	// 设置 接收后最终处理回调
	//SetCallback(v cellnet.EventCallback)
}
