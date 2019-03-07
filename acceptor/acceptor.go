// /////////////////////////////////////////////////////////////////////////////
// 支持格式配置的网络接收器

package acceptor

import (
	"net"

	"github.com/pkg/errors"            // 异常库
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
// Acceptor 对象

// 网络连接对象，支持 websocket tcp
type Acceptor struct {
	name       string                  // 组件名字
	laddr      *network.TLaddr         // 监听地址集合
	option     *TAcceptorOpt           // 配置参数
	acceptor   network.IAcceptor       // acceptor 连接器
	connNum    syncutil.AtomicUint32   // 当前连接数
	stateMgr   *state.StateManager     // 状态管理
	sessionMgr *session.SessionManager // session 管理对象
}

// 新建1个 Acceptor 对象
func NewAcceptor(addr *network.TLaddr, opt *TAcceptorOpt) (model.IComponent, error) {
	var err error
	var a network.IAcceptor

	// 参数效验
	if nil == opt {
		opt = NewTAcceptorOpt(nil)
	}

	// 创建对象
	sm := state.NewStateManager()
	sesMgr := session.NewSessionManager()

	// 创建 Acceptor
	actor := &Acceptor{
		stateMgr:   sm,
		name:       C_COMPONENT_NAME,
		laddr:      addr,
		option:     opt,
		sessionMgr: sesMgr,
	}

	// 创建 Acceptor
	a, err = network.NewAcceptor(opt.AcceptorName, addr, actor)
	if nil != err {
		return nil, err
	} else {
		actor.acceptor = a
	}

	// 设置为初始状态
	actor.stateMgr.SetState(state.C_INIT)

	return actor, nil
}

// 获取组件名字 [IComponent 接口]
func (this *Acceptor) Name() string {
	return this.name
}

// 运行 Acceptor [IComponent 接口]
func (this *Acceptor) Run() (err error) {
	// 改变状态： 启动中
	if !this.stateMgr.SwapState(state.C_INIT, state.C_RUNING) {
		if !this.stateMgr.SwapState(state.C_STOPED, state.C_RUNING) {
			err = errors.Errorf("Acceptor 组件启动失败，状态错误。当前状态=%d，正确状态=%d或=%d", this.stateMgr.GetState(), state.C_INIT, state.C_STOPED)
		}

		return
	}

	// acceptor 检查
	if nil == this.acceptor {
		err = errors.New("Acceptor 组件启动失败。acceptor=nil")

		return
	}

	// 启动 acceptor
	if err = this.acceptor.Run(); nil == err {
		return
	}

	// 改变状态： 工作中
	if !this.stateMgr.SwapState(state.C_RUNING, state.C_WORKING) {
		err = errors.Errorf("Acceptor 组件启动失败，状态错误。当前状态=%d，正确状态=%d", this.stateMgr.GetState(), state.C_RUNING)

		return
	}

	zaplog.Debugf("Acceptor 组件启动成功")

	return
}

// 停止 Acceptor [IComponent 接口]
func (this *Acceptor) Stop() (err error) {
	// 状态效验
	if !this.stateMgr.SwapState(state.C_WORKING, state.C_STOPING) {
		err = errors.Errorf("Acceptor 组件停止失败，状态错误。当前状态=%d，正确状态=%d", this.stateMgr.GetState(), state.C_WORKING)

		return
	}

	// 停止 acceptor
	if err = this.acceptor.Stop(); nil != err {
		return
	}

	// 关闭所有 session
	this.sessionMgr.CloseAllSession()

	// 改变状态：关闭完成
	if !this.stateMgr.SwapState(state.C_STOPING, state.C_STOPED) {
		err = errors.Errorf("Acceptor 组件停止失败，状态错误。当前状态=%d，正确状态=%d", this.stateMgr.GetState(), state.C_STOPING)

		return
	}

	zaplog.Infof("Acceptor 组件停止成功")

	return
}

// 收到1个新的 Tcp 连接对象
func (this *Acceptor) OnNewTcpConn(conn net.Conn) {
	zaplog.Debugf("收到1个新的 tcp 连接。ip=%s", conn.RemoteAddr())

	// 超过最大连接数
	if this.connNum.Load() >= this.option.MaxConn {
		conn.Close()

		zaplog.Warnf("Acceptor 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 不符合 tcp 连接对象
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		conn.Close()

		return
	}

	// 配置 iO 参数
	tcpConn.SetWriteBuffer(this.option.TcpConnOpt.WriteBufferSize)
	tcpConn.SetReadBuffer(this.option.TcpConnOpt.ReadBufferSize)
	tcpConn.SetNoDelay(this.option.TcpConnOpt.NoDelay)

	// 创建 session 对象
	this.createSession(conn, false)
}

// 收到1个新的 websocket 连接对象
func (this *Acceptor) OnNewWsConn(wsconn *websocket.Conn) {
	zaplog.Debugf("收到1个新的 websocket 连接。ip=%s", wsconn.RemoteAddr())

	// 超过最大连接数
	if this.connNum.Load() >= this.option.MaxConn {
		wsconn.Close()

		zaplog.Debugf("Acceptor 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 参数设置
	wsconn.PayloadType = websocket.BinaryFrame // 以二进制方式接受数据

	// 创建 session 对象
	this.createSession(wsconn, true)
}

// 创建 session 对象
func (this *Acceptor) createSession(netconn net.Conn, isWebSocket bool) {
	// 创建 socket
	socket := &network.Socket{
		Conn: netconn,
	}

	// 创建 session
	if this.option.ForClient {
		cses := session.NewClientSession(socket, this.sessionMgr, this.option.ClientSesOpt)

		cses.Run()
	} else {
		sses := session.NewServerSession(socket, this.sessionMgr, this.option.ServerSesOpt)

		sses.Run()
	}

	this.connNum.Add(1)
}
