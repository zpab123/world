// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- 组件接口

package model

// /////////////////////////////////////////////////////////////////////////////
// acceptor 相关

// connector 组件
type IConnector interface {
	IComponent                      // 接口继承： 组件接口
	GetAddr() *Laddr                // 获取地址信息集合
	GetConnectorOpt() *ConnectorOpt // 网络连接配置
	OnNewSocket(socket ISocket)     // 收到1个新的 socket 连接
	OnSocketClose(socket ISocket)   // 某个 socket  断开
	OnSocketMessage(socket ISocket) // 某个 socket  收到数据
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
	Flush() error
}
