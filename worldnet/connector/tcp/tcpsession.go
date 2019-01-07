// /////////////////////////////////////////////////////////////////////////////
// Session 会话对象

package tcp

import (
	"net"

	"github.com/zpab123/syncutil"       // 原子变量
	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// tcpSession 对象

// Socket 会话对象
type tcpSession struct {
	conn      net.Conn            // Socket 原始连接
	connector worldnet.IConnector // 符合 IConnector 接口的对象
	callBack  func()              // tcpSession run 成功后的回调函数
	sendQueue *worldnet.Pipe      // 发送消息队列
}

// 创建1个新的 tcpSession 对象
func newSession(conn net.Conn, cntor worldnet.IConnector, callBack func()) *tcpSession {
	// 创建发送队列
	que := worldnet.NewPipe()

	// 消息处理器

	// 创建 session
	ses := &tcpSession{
		conn:      conn,
		connector: cntor,
		callBack:  callBack,
		sendQueue: que,
	}

	return ses
}

// 启动 session
func (self *tcpSession) Run() {

	// 启动并发接收 goroutine
	go self.recvLoop()

	// 启动并发发送 goroutine
	go self.sendLoop()
}

// 接收循环
func (self *tcpSession) recvLoop() {

}

// 发送循环
func (self *tcpSession) sendLoop() {

}
