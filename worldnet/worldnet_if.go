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
