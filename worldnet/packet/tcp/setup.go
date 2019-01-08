// /////////////////////////////////////////////////////////////////////////////
// tcp 数据处理器注册

package tcp

import (
	"github.com/zpab123/worldnet"        // 网络库
	"github.com/zpab123/worldnet/consts" // 全局常量
	"github.com/zpab123/worldnet/proc"   // 111
)

// /////////////////////////////////////////////////////////////////////////////
// 包 初始化

func init() {
	proc.RegisterProcessor(consts.PROC_NAME_TCP_LTV, Binder)
}

// 绑定函数
func Binder(bundle proc.IProcessorBundle, userCallback worldnet.EventCallback) {
	// 设置消息发送器
	bundle.SetTransmitter(new(TCPMessageTransmitter))

	// 设置
	bundle.SetHooker(new(MsgHooker))
}
