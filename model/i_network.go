// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- 组件接口

package model

import (
	"net"
)

// /////////////////////////////////////////////////////////////////////////////
// acceptor 相关

// connector 组件
type IConnector interface {
	IComponent                            // 接口继承： 组件接口
	IRecoverIoPanic                       // 接口继承： 设置是否捕获 io 异常
	GetAddr() *Laddr                      // 获取地址信息集合
	GetConnectorOpt() *ConnectorOpt       // 网络连接配置
	OnNewSocket(socket IPacketSocket)     // 收到1个新的 socket 连接
	OnSocketClose(socket IPacketSocket)   // 某个 socket  断开
	OnSocketMessage(socket IPacketSocket) // 某个 socket  收到数据
}

// 开启 IO 层异常捕获, 在生产版本对外端口应该打开此设置
type IRecoverIoPanic interface {
	SetRecoverIoPanic(v bool) // 设置是否捕获 Io 异常 [IRecoverIoPanic 接口]
	GetRecoverIoPanic() bool  // 获取是否捕获 Io 异常 [IRecoverIoPanic 接口]
}

// /////////////////////////////////////////////////////////////////////////////
// acceptor 相关

// acceptor 接口
type IAcceptor interface {
	Run()    // 启动 acceptor
	Stop()   // 停止 acceptor
	IAddress // 接口继承： IAddress 接口
}

// 地址接口
type IAddress interface {
	SetAddr(addr *Laddr) // 设置监听地址
	GetAddr() *Laddr     // 获取监听地址
}

// socket 组件
type ISocket interface {
	net.Conn // 接口继承： 符合 Conn 的对象
	Flush() error
}

// 可以收发 packet 的 socket
type IPacketSocket interface {
}

// packet 消息收发接口
type IPacketPostter interface {
	// 接收 1个 packet
	RecvPacket(socket IPacketSocket) (msg interface{}, err error)
	// 发送 1个 packet
	SendPacket(socket IPacketSocket, msg interface{}) error
}
