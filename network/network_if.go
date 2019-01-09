// /////////////////////////////////////////////////////////////////////////////
// worldnet 接口汇总

package network

import (
	"github.com/zpab123/world/network/connector" // 连接器
)

// /////////////////////////////////////////////////////////////////////////////
// connector 接口

// connector 接口
type IConnector interface {
	Run()                   // 启动 connector
	Stop()                  // 停止 connector
	GetType() string        // 获取 connector 类型，例如 tcp.Connector/udp.Acceptor
	SetAddress(addr string) // 设置监听地址
	GetAddress() string     // 获取监听地址
}
