// /////////////////////////////////////////////////////////////////////////////
// worldnet 接口汇总

package worldnet

// /////////////////////////////////////////////////////////////////////////////
// Session 相关

// /////////////////////////////////////////////////////////////////////////////
// connector 相关

// connector 接口
type IConnector interface {
	Run()             // 启动 connector
	Stop()            // 停止 connector
	TypeName() string // connector 的类型(protocol.type)，例如tcp.Connector/udp.Acceptor
}

// Session 存取器接口
type ISessionAccessor interface {
	GetSession(int64) ISession        // 从 session 存取器中获取一个连接
	VisitSession(func(ISession) bool) // 遍历连接
	SessionCount() int                // 活跃的会话数量
	CloseAllSession()                 // 关闭所有连接
}

// 开启 IO 层异常捕获,在生产版本对外端口应该打开此设置
type IRecoverIoPanic interface {
	SetRecoverIoPanic(v bool) // 设置 IO 层是否异常捕获
	GetRecoverIoPanic() bool  // 获取 IO 层是否异常捕获
}

// /////////////////////////////////////////////////////////////////////////////
// Session 相关

// 长连接
type ISession interface {
	GetSocket() interface{}   // 获得原始的 Socket 连接
	GetConnector() IConnector // 获得 Session 归属的Peer
	Send(msg interface{})     // 发送消息，消息需要以指针格式传入
	Close()                   // 断开
	ID() int64                // 标示 ID
}

// /////////////////////////////////////////////////////////////////////////////
// packet 相关

// Packet 消息收发器
type IPacketManager interface {
	RecvPacket(ses ISession) (pkt interface{}, err error) // 接收 Packet 消息
	SendPacket(ses ISession, pkt interface{}) error       // 发送 Packet 消息
}

// 事件
type IEvent interface {
	GetSession() ISession   // 获取事件的 session 对象
	GetPacket() interface{} // 获取事件携带的 Packet 消息
}

// 用户端处理
type EventCallback func(ev IEvent)
