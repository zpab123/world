// /////////////////////////////////////////////////////////////////////////////
// PacketManager 数据处理相关

package connector

import (
	"errors"

	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 变量

var notPostter = errors.New("数据处理错误: 收发器 postter 为 nil ") // 收发器错误

// /////////////////////////////////////////////////////////////////////////////
// PacketManager 对象

// packet 数据管理
type PacketManager struct {
	postter worldnet.IPacketPostter // packet 收发管理：符合 IPacketPostter 接口的对象
	//callback worldnet.EventCallback  // 事件回调函数
	// 钩子
}

// 获取 网络消息管理对象 [IDataMananger 接口]
func (self *PacketManager) GetDataMananger() *PacketManager {
	return self
}

// 读取 packet [IPacketPostter 接口]
func (self *PacketManager) ReadPacket(ses worldnet.ISession) (pkt interface{}, err error) {
	if nil != self.postter {
		return self.postter.RecvPacket(ses)
	}

	return nil, notPostter
}

// 发送 packet [IPacketPostter 接口]
func (self *PacketManager) SendPacket(evt worldnet.IEvent) {
	if nil != self.postter && nil != evt {
		self.postter.SendPacket(evt.GetSession(), evt.GetPacket())
	}
}

// 发送事件
//
// PacketManager 将各种消息分发给需要处理数据的对象
func (self *PacketManager) SendEvent(evt worldnet.IEvent) {
	// 消息钩子

	// 回调函数
	if nil != self.callback && nil != evt {
		self.callback(evt)
	}
}
