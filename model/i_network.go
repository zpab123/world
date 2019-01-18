// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- network 网络消息包

package model

import (
	"net"
)

// /////////////////////////////////////////////////////////////////////////////
// acceptor 相关

// tcpSocket 连接管理
type ITcpSocketManager interface {
	OnNewTcpConn(conn net.Conn) // 收到1个新的 Tcp 连接对象
	CloseAllConn()              // 关闭所有连接
}

// acceptor 接口
type IAcceptor interface {
	Run()    // 启动 acceptor
	Stop()   // 停止 acceptor
	IAddress // 接口继承： IAddress 接口
}

// 地址接口
type IAddress interface {
	SetAddr(addr *TLaddr) // 设置监听地址
	GetAddr() *TLaddr     // 获取监听地址
}

// socket 组件
type ISocket interface {
	net.Conn // 接口继承： 符合 Conn 的对象
	Flush() error
}

// 可以收发 packet 的 socket
type IPacketSocket interface {
	ISocket                   // 接口继承： 符合 ISocket 的对象
	GetSocket() ISocket       // 获取 ISocket 对象
	GetConnector() IConnector // 获取 PacketSocket 对象所属 Connector
}

// 网络数据收发管理
type IDataManager interface {
	// 接收 1个 packet
	RecvPacket(socket IPacketSocket) (pkt interface{}, err error)
	// 发送 1个 packet
	SendPacket(socket IPacketSocket, pkt interface{}) error
}

// packet 消息管理接口
type IPacketManager interface {
	// 设置网络数据 接受/发送对象
	SetDataManager(dm IDataManager)
	// 设置 接收后，发送前的事件处理流程
	//SetHooker
	// 设置 接收后最终处理回调
	//SetCallback(v cellnet.EventCallback)
}

// packet 处理接口
type IPacketHandler interface {
	OnMessage(ses ISession, msg interface{})
}

// /////////////////////////////////////////////////////////////////////////////
// connector 相关

// connector 组件
type IConnector interface {
	IComponent                      // 接口继承： 组件接口
	IRecoverIoPanic                 // 接口继承： 设置是否捕获 io 异常
	ISocketOpt                      // 接口继承： socket IO 参数设置/获取
	GetConnectorOpt() *ConnectorOpt // 网络连接配置
	ISessionManage                  // 接口继承： session 管理
}

// 开启 IO 层异常捕获, 在生产版本对外端口应该打开此设置
type IRecoverIoPanic interface {
	SetRecoverIoPanic(v bool) // 设置是否捕获 Io 异常
	GetRecoverIoPanic() bool  // 获取是否捕获 Io 异常
}

// socket IO 参数接口
type ISocketOpt interface {
	GetMaxPacketSize() int                                // 获取 packet 最大字节
	SetSocketReadTimeout(conn net.Conn, callback func())  // 设置 socket 读超时时间
	SetSocketWriteTimeout(conn net.Conn, callback func()) // 设置 socket 写超时时间
}
