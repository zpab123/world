// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- proto 包

package model

import (
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// 常量

// WorldConnection 状态
const (
	C_WCONN_STATE_INIT     = iota // 初始化状态
	C_WCONN_STATE_WAIT_ACK        // 等待客户端握手ACK
	C_WCONN_STATE_WORKING         // 工作中
	C_WCONN_STATE_CLOSED          // 关闭状态
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
