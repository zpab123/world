// /////////////////////////////////////////////////////////////////////////////
// Session 会话对象

package tcp

import (
	"net"
	"sync"
	"time"

	"github.com/zpab123/syncutil"                 // 原子变量
	"github.com/zpab123/world/worldnet"           // 网络库
	"github.com/zpab123/world/worldnet/connector" // 连接器
)

// /////////////////////////////////////////////////////////////////////////////
// tcpSession 对象

// Socket 会话对象
type tcpSession struct {
	conn           net.Conn             // Socket 原始连接
	connector      worldnet.IConnector  // 符合 IConnector 接口的对象
	closeNotify    func()               // tcpSession 关闭成功后的回调函数
	sendQueue      *worldnet.Pipe       // 发送消息队列
	closing        syncutil.AtomicInt64 // 是否正在关闭中 0=否；1=是
	closeWaitGroup sync.WaitGroup       // 关闭同步器
}

// 创建1个新的 tcpSession 对象
//
// endNotify=tcpSession 关闭成功后的回调函数
func newSession(conn net.Conn, cntor worldnet.IConnector, endNotify func()) *tcpSession {
	// 创建发送队列
	que := worldnet.NewPipe()

	// 消息处理器

	// 创建 session
	ses := &tcpSession{
		conn:        conn,
		connector:   cntor,
		closeNotify: endNotify,
		sendQueue:   que,
	}

	return ses
}

// 启动 session
func (self *tcpSession) Run() {
	// 改变为非关闭状态
	self.closing.Store(0)

	// 清除发送队列
	self.sendQueue.Reset()

	// 需要接收和发送线程都退出，才算真正的退出
	self.closeWaitGroup.Add(2)

	// 将 session 添加到管理器, 在线程处理前添加到管理器(分配id), 避免ID还未分配,就开始使用id的竞态问题
	self.GetConnector().(connector.ISessionManager).Add(self)

	// 关闭监听线程
	go func() {
		// 等待2个线程结束
		self.closeWaitGroup.Wait()

		// 将 session 从管理器移除
		self.GetConnector().(connector.ISessionManager).Remove(self)

		// 关闭通知
		if nil != self.closeNotify {
			self.closeNotify()
		}
	}()

	// 启动并发接收 goroutine
	go self.recvLoop()

	// 启动并发发送 goroutine
	go self.sendLoop()
}

// 关闭 session
func (self *tcpSession) Close() {
	// 状态交换
	closing := self.closing.Swap(1)
	if self.closing.Load() != 0 {
		return
	}

	// 关闭连接
	if nil != self.conn {
		// 关闭读
		tcpcon := self.conn.(*net.TCPConn)
		tcpcon.CloseRead()

		// 手动读超时
		tcpcon.SetReadDeadline(time.Now())
	}
}

// 获取 session 对象的
func (self *tcpSession) GetConnector() worldnet.IConnector {
	return self.connector
}

// 接收循环
func (self *tcpSession) recvLoop() {

}

// 发送循环
func (self *tcpSession) sendLoop() {

}
