// /////////////////////////////////////////////////////////////////////////////
// Session 会话对象

package tcp

import (
	"net"
	"sync"
	"time"

	"github.com/zpab123/syncutil"                 // 原子变量
	"github.com/zpab123/world/utils"              // 工具库
	"github.com/zpab123/world/worldnet"           // 网络库
	"github.com/zpab123/world/worldnet/connector" // 连接器
	"github.com/zpab123/zplog"                    // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// tcpSession 对象

// Socket 会话对象
type tcpSession struct {
	conn                    net.Conn             // Socket 原始连接
	connector               worldnet.IConnector  // 符合 IConnector 接口的对象
	closeNotify             func()               // tcpSession 关闭成功后的回调函数
	sendQueue               *worldnet.Pipe       // 发送消息队列
	closing                 syncutil.AtomicInt64 // 是否正在关闭中 0=否；1=是
	closeWaitGroup          sync.WaitGroup       // 关闭同步器
	connector.PacketManager                      // 对象继承： packet 消息管理对象
	connector.IdMananger                         // 对象继承：ID管理
}

// 创建1个新的 tcpSession 对象
//
// endNotify=tcpSession 关闭成功后的回调函数
func newSession(conn net.Conn, cntor worldnet.IConnector, endNotify func()) *tcpSession {
	// 创建发送队列
	que := worldnet.NewPipe()

	// 消息处理器
	dm := cntor.(interface {
		GetDataMananger() *connector.DataManager
	}).GetDataMananger()
	//dm := cntor.GetDataMananger()

	// 创建 session
	ses := &tcpSession{
		conn:        conn,
		connector:   cntor,
		closeNotify: endNotify,
		sendQueue:   que,
		DataManager: dm,
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

// 关闭前，调用过 Session.Close
func (self *tcpSession) IsManualClosed() bool {
	return self.closing.Load() != 0
}

// 接收循环
func (self *tcpSession) recvLoop() {
	// 是否进行 io 异常捕获
	var capturePanic bool
	if i, ok := self.GetConnector().(worldnet.IRecoverIoPanic); ok {
		capturePanic = i.GetRecoverIoPanic()
	}

	// 接收数据
	for nil != self.conn {
		// 接收数据
		var pkt interface{}
		var err error
		if capturePanic {
			pkt, err = self.safeReadPacket()
		} else {
			pkt, err = self.ReadPacket(self)
		}

		// 接收错误
		if nil != err {
			// EOFO错误
			if !utils.IsEOFOrNetReadError(err) {
				zplog.Errorf("tcpSession 关闭， id=%d, err=%s", self.GetId(), err)
			}

			// 关闭发送线程
			self.sendQueue.Add(nil)

			// 发送关闭事件 - 标记为手动关闭原因
			closeMsg := &worldnet.SessionClosed{} // 消息
			if self.IsManualClosed() {
				closeMsg.Reason = worldnet.CloseReason_Manual
			}
			closeEvt := &worldnet.PacketEvent{ // 事件
				Session: self,
				Packet:  closeMsg,
			}
			self.SendEvent(closeEvt) // 发送事件

			break
		}

		// 派发数据 -- 这里是否用对象池？
		evt := &worldnet.PacketEvent{
			Session: self,
			Packet:  pkt,
		}
		self.SendEvent(evt)
	}

	// 关闭线程组数量 - 1
	self.closeWaitGroup.Done()
}

// 发送循环
func (self *tcpSession) sendLoop() {
	// 复制出发送队列数据
	var writeList []interface{}
	for {
		writeList = writeList[0:0]
		exit := self.sendQueue.Pick(&writeList) // 如果队列中存在 nil 数据 则 exit = true

		// 遍历要发送的数据
		for _, pkt := range writeList {
			// 创建事件 -- ？ 对象池？
			evt := &worldnet.PacketEvent{
				Session: self,
				Packet:  pkt,
			}

			// 发送事件
			self.SendPacket(evt)
		}

		// 需要结束线程
		if exit {
			break
		}
	}

	// 关闭 sokcet
	self.conn.Close()

	// 关闭线程组数量 - 1
	self.closeWaitGroup.Done()
}

// 安全地读取1个 packet 数据
//
// 该方法，会异常捕获，不会引起线程崩溃
func (self *tcpSession) safeReadPacket() (pkt interface{}, err error) {
	// 异常捕获
	defer func() {
		if err := recover(); err != nil {
			zplog.Warnf("io panic 恐慌=%s", err)
			self.conn.Close()
		}
	}()

	// 读取数据
	pkt, err = self.ReadPacket(self)

	return
}
