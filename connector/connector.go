// /////////////////////////////////////////////////////////////////////////////
// 支持格式配置的连接器

package connector

import (
	"net"

	"github.com/zpab123/syncutil"      // 原子操作工具
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
	"github.com/zpab123/world/state"   // 状态管理
	"github.com/zpab123/zplog"         // 日志库
	"golang.org/x/net/websocket"       // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 常量
const (
	_maxConnNum uint32 = 100000 // 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// /////////////////////////////////////////////////////////////////////////////
// connector 对象

// 网络连接对象，支持 websocket tcp
type Connector struct {
	*state.StateManager                           // 对象继承： 状态管理
	*session.SessionManager                       // 对象继承： session 管理对象
	name                    string                // 组件名字
	laddr                   *model.TLaddr         // 监听地址集合
	opts                    *model.TConnectorOpt  // 配置参数
	acceptor                model.IAcceptor       // 某种类型的 acceptor 连接器
	connNum                 syncutil.AtomicUint32 // 当前连接数
}

// 新建1个 Connector 对象
func NewConnector(addr *model.TLaddr, opt *model.TConnectorOpt) model.IConnector {
	// 地址检查？

	// 参数效验
	if nil == opt {
		opt := model.NewTConnectorOpt()
	}

	if nil != opt.Check() {
		return nil
	}

	// 创建 SessionManager
	sesMgr := session.NewSessionManager()

	// 创建 Acceptor
	aptor, _ := newAcceptor(opt.AcceptorName, addr, cntor)

	// 创建组件
	cntor := &Connector{
		name:           model.C_CPT_NAME_CONNECTOR,
		laddr:          addr,
		opts:           opt,
		SessionManager: sesMgr,
	}

	// 创建 Acceptor
	aptor, _ := newAcceptor(opt.AcceptorName, addr, cntor)
	cntor.acceptor = aptor

	// 设置为初始状态
	cntor.state.Store(model.C_STATE_INIT)

	return cntor
}

// 运行 Connector [IComponent 接口]
func (this *Connector) Name() string {
	return this.name
}

// 运行 Connector [IComponent 接口]
func (this *Connector) Run() bool {
	// 状态效验
	if !this.CompareAndSwap(model.C_STATE_INIT, model.C_STATE_RUNING) {
		zplog.Errorf("Connector 组件启动失败，状态错误。正确状态=%d，当前状态=%d", model.C_STATE_INIT, this.state.Load())

		return false
	}

	// 添加启动线程数量
	this.AddRunGo(1)

	// 停止线程+1
	this.AddStopGo(1)

	// 启动 acceptor
	this.acceptor.Run()

	// 阻塞
	this.RunWait()

	// 改变状态： 工作中
	if !this.CompareAndSwap(model.C_STATE_RUNING, model.C_STATE_WORKING) {
		zplog.Errorf("Connector 组件启动失败，状态错误。正确状态=%d，当前状态=%d", model.C_STATE_RUNING, this.state.Load())

		return false
	}

	zplog.Infof("Connector 组件启动成功")

	return true
}

// 停止 Connector [IComponent 接口]
func (this *Connector) Stop() bool {
	// 状态效验
	if !this.CompareAndSwap(model.C_STATE_WORKING, model.C_STATE_STOPING) {
		zplog.Errorf("Connector 组件停止失败，状态错误。正确状态=%d，当前状态=%d", model.C_STATE_WORKING, this.state.Load())

		return false
	}

	// 停止 acceptor
	this.acceptor.Stop()

	// 阻塞
	this.StopWait()

	// 关闭所有 session
	this.CloseAllSession()

	// 改变状态：关闭完成
	if !this.CompareAndSwap(model.C_STATE_STOPING, model.C_STATE_STOP) {
		zplog.Errorf("Connector 组件停止失败，状态错误。正确状态=%d，当前状态=%d", model.C_STATE_STOPING, this.state.Load())

		return false
	}

	zplog.Infof("Connector 组件停止成功")

	return true
}

// 收到1个新的 Tcp 连接对象
func (this *Connector) OnNewTcpConn(conn net.Conn) {
	// 超过最大连接数
	if this.connNum.Load() >= this.opts.MaxConn {
		conn.Close()
		zplog.Debugf("收到1个新的 tcp 连接。ip=%s", tcpConn.RemoteAddr())
		zplog.Debugf("Connector 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 不符合 tcp 连接对象
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return
	}

	zplog.Debugf("收到1个新的 tcp 连接。ip=%s", tcpConn.RemoteAddr())

	// 配置 iO 参数
	tcpConn.SetWriteBuffer(this.opts.TcpConnOpts.WriteBufferSize)
	tcpConn.SetReadBuffer(this.opts.TcpConnOpts.ReadBufferSize)
	tcpConn.SetNoDelay(this.opts.TcpConnOpts.NoDelay)

	// 创建 session 对象
}

// 收到1个新的 websocket 连接对象
func (this *Connector) OnNewWsConn(wsconn *websocket.Conn) {
	// 超过最大连接数
	if this.connNum.Load() >= this.opts.MaxConn {
		wsconn.Close()
		zplog.Debugf("收到1个新的 websocket 连接。ip=%s", wsconn.RemoteAddr())
		zplog.Debugf("Connector 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 参数设置
	wsconn.PayloadType = websocket.BinaryFrame // 以二进制方式接受数据

	// 创建 session 对象
}

// 某个 Acceptor 启动完成 [IAcceptorState 接口]
func (this *Connector) OnAcceptorWorkIng() {
	// 启动线程完成1个
	this.RunDone()
}

// 某个 Acceptor 关闭完成 [IAcceptorState 接口]
func (this *Connector) OnAcceptorClosed() {
	// 结束线程完成1个
	this.StopDone()
}

// 创建 session 对象
func (this *Connector) createSession(netconn net.Conn, isWebSocket bool) {
	// 创建 socket
	socket := &network.Socket{
		conn: netconn,
	}

	// 创建 session
}
