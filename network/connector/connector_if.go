// /////////////////////////////////////////////////////////////////////////////
// connector 接口汇总

package connector

//"github.com/zpab123/world/network" // 网络库

// /////////////////////////////////////////////////////////////////////////////
// acceptor 相关

// acceptor 接口
type IAcceptor interface {
	Run()            // 启动 acceptor
	Stop()           // 停止 acceptor
	GetType() string // 获取 acceptor 类型，例如 tcp.Connector/udp.Acceptor
	IAddress         // 接口继承： IAddress 接口
}

// 地址接口
type IAddress interface {
	SetAddr(addr *Laddr) // 设置监听地址
	GetAddr() *Laddr     // 获取监听地址
}
