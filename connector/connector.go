// /////////////////////////////////////////////////////////////////////////////
// 支持格式配置的连接器

package connector

import (
	"net"

	"github.com/zpab123/syncutil"      // 原子操作工具
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络模型
	"github.com/zpab123/world/session" // session 库
	"github.com/zpab123/world/state"   // 状态管理
	"github.com/zpab123/zplog"         // 日志库
	"golang.org/x/net/websocket"       // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// public api

// /////////////////////////////////////////////////////////////////////////////
// connector 对象

// 网络连接对象，支持 websocket tcp
type Connector struct {
	name       string                  // 组件名字
	laddr      *network.TLaddr         // 监听地址集合
	opts       *TConnectorOpt          // 配置参数
	acceptor   network.IAcceptor       // 某种类型的 acceptor 连接器
	connNum    syncutil.AtomicUint32   // 当前连接数
	stateMgr   *state.StateManager     // 状态管理
	sessionMgr *session.SessionManager // session 管理对象
}

// 新建1个 Connector 对象
func NewConnector(addr *network.TLaddr, opt *TConnectorOpt) model.IComponent {
	// 地址检查？

	// 参数效验
	if nil == opt {
		opt = NewTConnectorOpt()
	}

	if nil != opt.Check() {
		return nil
	}

	// 创建 StateManager
	sm := state.NewStateManager()

	// 创建 SessionManager
	sesMgr := session.NewSessionManager()

	// 创建组件
	cntor := &Connector{
		stateMgr:   sm,
		name:       COMPONENT_NAME,
		laddr:      addr,
		opts:       opt,
		sessionMgr: sesMgr,
	}

	// 创建 Acceptor
	aptor, _ := newAcceptor(opt.AcceptorName, addr, cntor)
	cntor.acceptor = aptor

	// 设置为初始状态
	cntor.stateMgr.SetState(state.C_STATE_INIT)

	return cntor
}

// 运行 Connector [IComponent 接口]
func (this *Connector) Name() string {
	return this.name
}

// 运行 Connector [IComponent 接口]
func (this *Connector) Run() bool {
	// 改变状态： 启动中
	if !this.stateMgr.SwapState(state.C_STATE_INIT, state.C_STATE_RUNING) {
		zplog.Errorf("Connector 组件启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INIT, this.stateMgr.GetState())

		return false
	}

	// acceptor 检查
	if nil == this.acceptor {
		zplog.Error("Connector 组件启动失败。acceptor=nil")

		return false
	}

	// 启动 acceptor
	if !this.acceptor.Run() {
		return false
	}

	// 改变状态： 工作中
	if !this.stateMgr.SwapState(state.C_STATE_RUNING, state.C_STATE_WORKING) {
		zplog.Errorf("Connector 组件启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_RUNING, this.stateMgr.GetState())

		return false
	}

	zplog.Infof("Connector 组件启动成功")

	return true
}

// 停止 Connector [IComponent 接口]
func (this *Connector) Stop() bool {
	// 状态效验
	if !this.stateMgr.SwapState(state.C_STATE_WORKING, state.C_STATE_STOPING) {
		zplog.Errorf("Connector 组件停止失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_WORKING, this.stateMgr.GetState())

		return false
	}

	// 停止 acceptor
	if !this.acceptor.Stop() {
		return false
	}

	// 关闭所有 session
	this.sessionMgr.CloseAllSession()

	// 改变状态：关闭完成
	if !this.stateMgr.SwapState(state.C_STATE_STOPING, state.C_STATE_STOP) {
		zplog.Errorf("Connector 组件停止失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_STOPING, this.stateMgr.GetState())

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

		zplog.Warnf("Connector 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 不符合 tcp 连接对象
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		conn.Close()

		return
	}

	zplog.Debugf("收到1个新的 tcp 连接。ip=%s", tcpConn.RemoteAddr())

	// 配置 iO 参数
	tcpConn.SetWriteBuffer(this.opts.TcpConnOpts.WriteBufferSize)
	tcpConn.SetReadBuffer(this.opts.TcpConnOpts.ReadBufferSize)
	tcpConn.SetNoDelay(this.opts.TcpConnOpts.NoDelay)

	// 创建 session 对象
	this.createSession(conn, false)
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

// 创建 session 对象
func (this *Connector) createSession(netconn net.Conn, isWebSocket bool) {
	// 创建 socket
	socket := &network.Socket{
		Conn: netconn,
	}

	// 创建 session
	opt := session.NewTSessionOpts()
	ses := session.NewFrontendSession(socket, this.sessionMgr, opt)

	// 启动 session
	ses.Run()
}
