// /////////////////////////////////////////////////////////////////////////////
// dispatcher 客户端连接管理

package dispatcher

import (
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
)

// 分发客户端
type DispatcherConnMgr struct {
	addr       string                  // 服务器地址
	sessionMgr *session.SessionManager // session 管理对象
	option     *TDispatcherClientOpt   // 配置参数
	disConn    *DispatcherConn         // 连接对象

}

// 新建1个 DispatcherConnMgr
func NewDispatcherConnMgr(addr string, mgr ISessionManage, opt *TDispatcherClientOpt) *DispatcherConnMgr {
	dc := &DispatcherConnMgr{
		addr:       addr,
		sessionMgr: mgr,
		option:     opt,
	}

	return dc
}

// 启动 DispatcherConnMgr
func (this *DispatcherConnMgr) Run() {
	// 连接服务器
	this.connectServer()

	// 主循环
}

//

// 连接服务器
func (this *DispatcherConnMgr) connectServer() (*session.BackendSession, error) {
	conn, err := net.Dial("tcp", this.addr)
	if nil != err {
		return nil, err
	}

	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetReadBuffer(this.option.TcpConnOpt.ReadBufferSize)
	tcpConn.SetWriteBuffer(this.option.TcpConnOpt.WriteBufferSize)

	// 创建代理
	//clientProxy = NewDispatcherConn(conn, this.option.WorldConnOpts)

	// 创建 BackendSession
	socket := &network.Socket{
		Conn: netconn,
	}
	ss := session.NewBackendSession(sock, this.sessionMgr, this.option.SessionOpt)

	ss.Run()

	return ss, nil
}
