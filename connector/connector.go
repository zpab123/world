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
	"github.com/zpab123/zaplog"        // 日志库
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
	option     *TConnectorOpt          // 配置参数
	acceptor   network.IAcceptor       // acceptor 连接器
	connNum    syncutil.AtomicUint32   // 当前连接数
	stateMgr   *state.StateManager     // 状态管理
	sessionMgr *session.SessionManager // session 管理对象
}

// 新建1个 Connector 对象
func NewConnector(addr *network.TLaddr, opt *TConnectorOpt) model.IComponent {
	// 参数效验
	if nil == opt {
		opt = NewTConnectorOpt(nil)
	}

	// 创建对象
	sm := state.NewStateManager()
	sesMgr := session.NewSessionManager()

	// 创建 Connector
	cntor := &Connector{
		stateMgr:   sm,
		name:       C_COMPONENT_NAME,
		laddr:      addr,
		option:     opt,
		sessionMgr: sesMgr,
	}

	// 创建 Acceptor
	actor, _ := network.NewAcceptor(opt.AcceptorType, addr, cntor)
	if nil == actor {
		return nil
	} else {
		cntor.acceptor = actor
	}

	// 设置为初始状态
	cntor.stateMgr.SetState(state.C_STATE_INIT)

	return cntor
}

// 获取组件名字 [IComponent 接口]
func (this *Connector) Name() string {
	return this.name
}

// 运行 Connector [IComponent 接口]
func (this *Connector) Run() bool {
	// 改变状态： 启动中
	if !this.stateMgr.SwapState(state.C_STATE_INIT, state.C_STATE_RUNING) {
		zaplog.Errorf("Connector 组件启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INIT, this.stateMgr.GetState())

		return false
	}

	// acceptor 检查
	if nil == this.acceptor {
		zaplog.Error("Connector 组件启动失败。acceptor=nil")

		return false
	}

	// 启动 acceptor
	if err := this.acceptor.Run(); nil == err {
		return false
	}

	// 改变状态： 工作中
	if !this.stateMgr.SwapState(state.C_STATE_RUNING, state.C_STATE_WORKING) {
		zaplog.Errorf("Connector 组件启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_RUNING, this.stateMgr.GetState())

		return false
	}

	zaplog.Debugf("Connector 组件启动成功")

	return true
}

// 停止 Connector [IComponent 接口]
func (this *Connector) Stop() bool {
	// 状态效验
	if !this.stateMgr.SwapState(state.C_STATE_WORKING, state.C_STATE_STOPING) {
		zaplog.Errorf("Connector 组件停止失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_WORKING, this.stateMgr.GetState())

		return false
	}

	// 停止 acceptor
	if err := this.acceptor.Stop(); nil != err {
		return false
	}

	// 关闭所有 session
	this.sessionMgr.CloseAllSession()

	// 改变状态：关闭完成
	if !this.stateMgr.SwapState(state.C_STATE_STOPING, state.C_STATE_STOP) {
		zaplog.Errorf("Connector 组件停止失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_STOPING, this.stateMgr.GetState())

		return false
	}

	zaplog.Infof("Connector 组件停止成功")

	return true
}

// 收到1个新的 Tcp 连接对象
func (this *Connector) OnNewTcpConn(conn net.Conn) {
	// 超过最大连接数
	if this.connNum.Load() >= this.option.MaxConn {
		conn.Close()

		zaplog.Warnf("Connector 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 不符合 tcp 连接对象
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		conn.Close()

		return
	}

	zaplog.Debugf("收到1个新的 tcp 连接。ip=%s", tcpConn.RemoteAddr())

	// 配置 iO 参数
	tcpConn.SetWriteBuffer(this.option.TcpConnOpt.WriteBufferSize)
	tcpConn.SetReadBuffer(this.option.TcpConnOpt.ReadBufferSize)
	tcpConn.SetNoDelay(this.option.TcpConnOpt.NoDelay)

	// 创建 session 对象
	this.createSession(conn, false)
}

// 收到1个新的 websocket 连接对象
func (this *Connector) OnNewWsConn(wsconn *websocket.Conn) {
	// 超过最大连接数
	if this.connNum.Load() >= this.option.MaxConn {
		wsconn.Close()
		zaplog.Debugf("收到1个新的 websocket 连接。ip=%s", wsconn.RemoteAddr())
		zaplog.Debugf("Connector 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 参数设置
	wsconn.PayloadType = websocket.BinaryFrame // 以二进制方式接受数据

	// 创建 session 对象
	this.createSession(wsconn, true)
}

// 创建 session 对象
func (this *Connector) createSession(netconn net.Conn, isWebSocket bool) {
	// 创建 socket
	socket := &network.Socket{
		Conn: netconn,
	}

	// 创建 session
	if this.option.Frontend {
		ses := session.NewFrontendSession(socket, this.sessionMgr, this.option.FrontendSessionOpt)

		ses.Run()
	} else {
		ses := session.NewBackendSession(socket, this.sessionMgr, this.option.BackendSessionOpt)

		ses.Run()
	}
}
