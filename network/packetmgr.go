// /////////////////////////////////////////////////////////////////////////////
// PacketManager 数据处理相关

package network

import (
	"errors"

	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
)

// /////////////////////////////////////////////////////////////////////////////
// 包 变量

var notPostter = errors.New("数据处理错误: 收发器 postter 为 nil ") // 收发器错误

// /////////////////////////////////////////////////////////////////////////////
// PacketManager 对象

// packet 数据管理
type PacketManager struct {
	postter model.IPacketPostter // packet 收发管理：符合 IPacketPostter 接口的对象
	//callback worldnet.EventCallback  // 事件回调函数
	// 钩子
}

// 读取 packet [IPacketPostter 接口]
func (this *PacketManager) RecvPacket(ses model.ISession, recoverPanic bool) (pkt interface{}, err error) {
	if nil != this.postter {
		return this.postter.RecvPacket(ses)
	}

	return nil, notPostter
}

// 发送 packet [IPacketPostter 接口]
func (this *PacketManager) SendPacket(pkt interface{}) {
	if nil != this.postter && nil != pkt {
		//this.postter.SendPacket(evt.GetSession(), evt.GetPacket())
	}
}
