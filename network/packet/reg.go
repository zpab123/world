// /////////////////////////////////////////////////////////////////////////////
// packet 处理器 注册中心

package packet

import (
	"fmt"

	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 包 变量

// 变量
var (
	postCreatorMap = map[string]PostCreator{} // pktType->Creator 对象集合
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 注册1个 packet 投递对象
func RegisterPostter(pktType string, f Creator) {
	postCreatorMap[pktType] = f
}

// 设置 packet 收发方法
func SetPacketPostter(pm IPacketManager, pktType string) {
	// 创建收发对象
	if postCreator, ok := postCreatorMap[pktType]; ok {
		postCreator(pm)
	} else {
		panic(fmt.Sprintf("设置 packet 收发对象错误： pktType=%s；不存在", pktType))
	}
}

// /////////////////////////////////////////////////////////////////////////////
// Creator 对象

// 创建函数
type PostCreator func(pm IPacketManager)
