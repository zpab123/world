// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- proto 包

package model

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// TWorldConnOpts 对象

// WorldConnection 配置参数
type TWorldConnOpts struct {
	pakcetSocket IPacketSocket // socket 对象
	DataType     string        // packet 数据结构类型
	Heartbeat    time.Duration // 心跳间隔
	Handshake    func()        // 自定义的握手处理函数
}
