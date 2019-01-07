// /////////////////////////////////////////////////////////////////////////////
// processor 数据处理相关 [代码完整]

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
	// 消息传输
	packetManager worldnet.IPacketManager // 符合 IPacketManager 接口的对象
	// 钩子
	// 回调函数
}

// 读取消息
func (self *DataManager) ReadPacket(ses worldnet.ISession) (pkt interface{}, err error) {
	if nil != self.packetManager {
		return self.packetManager.RecvPacket(ses)
	}

	return nil, notPacketManager
}
