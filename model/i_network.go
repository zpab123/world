// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- network 网络消息包

package model

import (
	"net"

	"golang.org/x/net/websocket" // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// acceptor 相关

// tcpSocket 连接管理
type ITcpSocketManager interface {
	OnNewTcpConn(conn net.Conn) // 收到1个新的 Tcp 连接对象
	CloseAllConn()              // 关闭所有连接
}

// websocket 连接管理
type IWebsocketManager interface {
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
	CloseAllConn()                      // 关闭所有连接
}

// MulAcceptor 连接管理
type IMulConnManager interface {
	OnNewTcpConn(conn net.Conn)         // 收到1个新的 Tcp 连接对象
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
	CloseAllConn()                      // 关闭所有连接
}

// ComAcceptor 连接管理
type IComConnManager interface {
	OnNewTcpConn(conn net.Conn)         // 收到1个新的 Tcp 连接对象
	OnNewWsConn(wsconn *websocket.Conn) // 收到1个新的 websocket 连接对象
	CloseAllConn()                      // 关闭所有连接
}

// acceptor 接口
type IAcceptor interface {
	Run()  // 组件开始运行
	Stop() // 组件停止运行
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
