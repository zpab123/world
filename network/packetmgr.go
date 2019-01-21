// /////////////////////////////////////////////////////////////////////////////
// PacketManager 数据处理相关

package network

import (
	"errors"

	"github.com/zpab123/world/model" // 全局模型
)

// /////////////////////////////////////////////////////////////////////////////
// 包 变量

var notPostter = errors.New("数据处理错误: 收发器 postter 为 nil ") // 收发器错误

// /////////////////////////////////////////////////////////////////////////////
// PacketManager 对象

// packet 数据管理
type PacketManager struct {
	dataManager model.IDataManager // 网络数据收发管理对象
	//callback worldnet.EventCallback  // 事件回调函数
	// 钩子
}

// 接收1个 packet [IDataManager 接口]
func (this *PacketManager) RecvPacket(socket model.IPacketSocket) (pkt interface{}, err error) {
	if nil != this.dataManager {
		return this.dataManager.RecvPacket(socket)
	}

	return nil, notPostter
}

// 发送1个 packet [IDataManager 接口]
func (this *PacketManager) SendPacket(pkt interface{}) {
	if nil != this.dataManager && nil != pkt {
		//this.postter.SendPacket(evt.GetSession(), evt.GetPacket())
	}
}
