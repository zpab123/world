// /////////////////////////////////////////////////////////////////////////////
// DataManager 数据处理相关

package connector

import (
	"errors"

	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 变量

var notPacketManager = errors.New("数据处理错误: 收发器 packetManager 为 nil ") // 收发器错误

// /////////////////////////////////////////////////////////////////////////////
// DataManager 对象

// 网络数据管理
type DataManager struct {
	packetManager worldnet.IPacketManager // 网络消息管理：符合 IPacketManager 接口的对象
	callback      worldnet.EventCallback  // 事件回调函数
	// 钩子
}

// 获取 网络消息管理对象 [IDataMananger 接口]
func (self *DataManager) GetDataMananger() *DataManager {
	return self
}

// 读取 packet [IPacketManager 接口]
func (self *DataManager) ReadPacket(ses worldnet.ISession) (pkt interface{}, err error) {
	if nil != self.packetManager {
		return self.packetManager.RecvPacket(ses)
	}

	return nil, notPacketManager
}

// 发送 packet [IPacketManager 接口]
func (self *DataManager) SendPacket(evt worldnet.IEvent) {
	if nil != self.packetManager && nil != evt {
		self.packetManager.SendPacket(evt.GetSession(), evt.GetMessage())
	}
}

// 发送事件
func (self *DataManager) SendEvent(evt worldnet.IEvent) {
	// 消息钩子

	// 回调函数
	if nil != self.callback && nil != evt {
		self.callback(evt)
	}
}
