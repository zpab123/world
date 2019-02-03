// /////////////////////////////////////////////////////////////////////////////
// 消息转发服务

package dispatcher

import (
	"net"

	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
	"github.com/zpab123/world/state"   // 状态管理
)

// 分发服务
type DispatcherServer struct {
	maxConn    uint32                  // 最大连接数量，超过此数值后，不再接收新连接
	option     *TDispatcherServerOpts  // 配置参数
	connNum    syncutil.AtomicUint32   // 当前连接数
	acceptor   *network.TcpAcceptor    // tcp 连接器
	stateMgr   *state.StateManager     // 状态管理
	sessionMgr *session.SessionManager // session 管理对象
}

// 新建1个分发服务
func NewDispatcherServer(addr *network.TLaddr, opts *TDispatcherServerOpts) *DispatcherServer {
	// 参数效验
	if nil == opts {
		opts = NewTDispatcherServerOpts()
	}

	// 创建组件
	sm := state.NewStateManager()
	sesMgr := session.NewSessionManager()

	// 创建 DispatcherServer
	ds := &DispatcherServer{
		maxConn:    opts.MaxConn,
		option:     opts,
		stateMgr:   sm,
		sessionMgr: sesMgr,
	}

	// 创建 acceptor
	acpor := network.NewTcpAcceptor(addr, ds)
	ds.acceptor = acpor

	// 设置为初始状态
	ds.stateMgr.SetState(state.C_STATE_INIT)

	return ds
}

// 启动 DispatcherServer
func (this *DispatcherServer) Run() {
	this.acceptor.Run()
}

// 收到1个新的 Tcp 连接对象
func (this *DispatcherServer) OnNewTcpConn(conn net.Conn) {
	// 超过最大连接数
	if this.connNum.Load() >= this.maxConn {
		conn.Close()

		zaplog.Warnf("DispatcherServer 达到最大连接数，关闭新连接。当前连接数=%d", this.connNum.Load())
	}

	// 不符合 tcp 连接对象
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		conn.Close()

		return
	}

	zaplog.Debugf("DispatcherServer 收到1个新的 tcp 连接。ip=%s", tcpConn.RemoteAddr())

	// 配置 iO 参数
	tcpConn.SetReadBuffer(this.option.TcpConnOpts.ReadBufferSize)
	tcpConn.SetWriteBuffer(this.option.TcpConnOpts.WriteBufferSize)
	tcpConn.SetNoDelay(this.option.TcpConnOpts.NoDelay)

	// 创建服务器 session
	this.createSession(conn)
}

// 创建 session 对象
func (this *DispatcherServer) createSession(netconn net.Conn) {

}
