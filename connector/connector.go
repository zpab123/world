// /////////////////////////////////////////////////////////////////////////////
// 支持格式配置的连接器

package connector

import (
	"net"

	"github.com/zpab123/syncutil"      // 原子操作工具
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
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
	*session.SessionManager                       // 对象继承： session 管理对象
	name                    string                // 组件名字
	laddr                   *model.TLaddr         // 监听地址集合
	opts                    *model.TConnectorOpt  // 配置参数
	acceptor                model.IAcceptor       // 某种类型的 acceptor 连接器
	connNum                 syncutil.AtomicUint32 // 当前连接数
	state                   syncutil.AtomicUint32 // connector 当前状态
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
func (this *Connector) Run() {
	// 状态效验
	if this.state.Load() != model.C_STATE_INIT {
		return
	}

	// 改变状态： 正在启动中
	this.state.Store(model.C_STATE_RUNING)

	// 启动 acceptor
	this.acceptor.Run()
}

// 停止 Connector [IComponent 接口]
func (this *Connector) Stop() {
	// 停止 acceptor
	this.acceptor.Stop()
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

// 关闭所有连接
func (this *Connector) CloseAllConn() {

}

// 创建 session 对象
func (this *Connector) createSession(netconn net.Conn, isWebSocket bool) {
	// 创建 socket
	socket := &network.Socket{
		conn: netconn,
	}

	// 创建 session
}
