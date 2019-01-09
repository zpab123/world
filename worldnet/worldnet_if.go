// /////////////////////////////////////////////////////////////////////////////
// worldnet 接口汇总

package worldnet

import (
	"github.com/zpab123/world/worldnet/connector" // 连接器
)

// /////////////////////////////////////////////////////////////////////////////
// connector 接口

// Session 存取器接口
type ISessionAccessor interface {
	GetSession(int64) ISession        // 从 session 存取器中获取一个连接
	VisitSession(func(ISession) bool) // 遍历连接
	SessionCount() int                // 活跃的会话数量
	CloseAllSession()                 // 关闭所有连接
}

// Session 接口
type ISession interface {
	Raw() interface{}     // 获得原始的 Socket 连接
	Connector()           // 获得 Session 归属的 Connector
	Send(pkt interface{}) // 发送 pkt 消息，消息需要以指针格式传入
	Close()               // 断开连接
	GetId() int64         // 获取 session Id
	SetId(v int64)        // 设置 session Id
}

// connector 接口
type IConnector interface {
	//connector.IDataMananger // 继承 IDataMananger 接口
	Run() IConnector // 启动 connector
	Stop()           // 停止通讯端
	GetType() string // 获取 connector 类型，例如 tcp.Connector/udp.Acceptor
}

// Connector 基础属性
type IConnectorProperty interface {
	Name() string          // 获取名字
	Address() string       // 获取地址
	Queue() EventQueue     // 事件队列
	SetName(v string)      // 设置名称（可选）
	SetAddress(v string)   // 设置 Connector 地址
	SetQueue(v EventQueue) // 设置 Connector 挂接队列（可选）
}

// 基本的通用Peer
type IBaseConnector interface {
	IConnector         // 接口继承： connector 接口
	IConnectorProperty // 接口继承： Connector 基础属性
}

// 事件队列 接口
type IEventQueue interface {
	StartLoop() IEventQueue    // 事件队列开始工作
	StopLoop() IEventQueue     // 停止事件队列
	Wait()                     // 等待退出消息 -- 阻塞线程
	Post(callback func())      // 投递事件, 通过队列到达消费者端
	EnableCapturePanic(v bool) // 是否捕获异常
}

// io 异常捕获接口
//
// 开启IO层异常捕获,在生产版本对外端口应该打开此设置
type ICaptureIOPanic interface {
	EnableCaptureIOPanic(v bool) // 是否开启IO层异常捕获
	CaptureIOPanic() bool        // 获取当前异常捕获值
}

// 设置和获取自定义属性
type IContextSet interface {
	// 为对象设置一个自定义属性
	SetContext(key interface{}, v interface{})
	// 从对象上根据key获取一个自定义属性
	GetContext(key interface{}) (interface{}, bool)
	// 给定一个值指针, 自动根据值的类型GetContext后设置到值
	FetchContext(key, valuePtr interface{}) bool
}

// /////////////////////////////////////////////////////////////////////////////
// packet 数据处理相关

// packet 消息收发接口
type IPacketManager interface {
	// 接收 1个 packet
	RecvPacket(ses ISession) (msg interface{}, err error)
	// 发送 1个 packet
	SendPacket(ses ISession, msg interface{}) error
}

// 处理钩子(参数输入, 返回输出, 不给MessageProccessor处理时，可以将Event设置为nil)
type IEventHooker interface {
	OnInboundEvent(input Event) (output Event)  // 入站(接收)的事件处理
	OnOutboundEvent(input Event) (output Event) // 出站(发送)的事件处理
}

// 消息事件
type IEvent interface {
	Session() Session     // 事件对应的会话
	Message() interface{} // 事件携带的消息
}

// 消息回调函数
type EventCallback func(ev IEvent)

// /////////////////////////////////////////////////////////////////////////////
// Codec 数据编码/解码处理相关

// 编码/解码 接口
type ICodec interface {
	// 将数据转换为字节数组
	Encode(msgObj interface{}, ctx IContextSet) (data interface{}, err error)
	// 将字节数组转换为数据
	Decode(data interface{}, msgObj interface{}) error
	// 编码器的名字
	Name() string
	// 兼容http类型
	MimeType() string
}
