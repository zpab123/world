// /////////////////////////////////////////////////////////////////////////////
// 顶级接口 -- 组件接口

package ifs

// /////////////////////////////////////////////////////////////////////////////
// component 相关

// component 组件
type IComponent interface {
	Name() string // 获取组件名字
	Run()         // 组件开始运行
	Stop()        // 组件停止运行
}

// /////////////////////////////////////////////////////////////////////////////
// connector 相关

// socket 组件
type ISocket interface {
	Flush() error
}

// connector 组件
type IConnector interface {
	OnNewSocket(socket ISocket)     // 收到1个新的 socket 连接
	OnSocketClose(socket ISocket)   // 某个 socket  断开
	OnSocketMessage(socket ISocket) // 某个 socket  收到数据
}

// session 组件
type ISession interface {
	Send(data interface{})
}

// message 组件
type IMessage interface {
}
