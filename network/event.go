// /////////////////////////////////////////////////////////////////////////////
// 网络事件 [代码完整]

package worldnet

// /////////////////////////////////////////////////////////////////////////////
// PacketEvent

// 接收到消息事件
type PacketEvent struct {
	Session ISession    // 符合 ISession 的对象
	Packet  interface{} // 消息数据
}

// 获取 session [IEvent 接口]
func (self *PacketEvent) GetSession() ISession {
	return self.Ses
}

// 获取消息 [IEvent 接口]
func (self *PacketEvent) GetMessage() interface{} {
	return self.Pkt
}
