// /////////////////////////////////////////////////////////////////////////////
// connector 接口汇总

package connector

//"github.com/zpab123/world/network" // 网络库

// /////////////////////////////////////////////////////////////////////////////
// connector 相关

// connector 接口
type IConnector interface {
	Run()            // 启动 connector
	Stop()           // 停止 connector
	GetType() string // 获取 connector 类型，例如 tcp.Connector/udp.Acceptor
	IAddress         // 接口继承： IAddress 接口
}

// 地址接口
type IAddress interface {
	SetAddr(addr *Laddr) // 设置监听地址
	GetAddr() *Laddr     // 获取监听地址
}
