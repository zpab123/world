// /////////////////////////////////////////////////////////////////////////////
// 支持格式配置的连接器

package connector

import (
	"fmt"
	"net"

	"github.com/zpab123/syncutil"      // 原子操作工具
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
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
	name     string                // 组件名字
	laddr    *model.TLaddr         // 监听地址集合
	opts     *model.TConnectorOpt  // 配置参数
	acceptor model.IAcceptor       // 某种类型的 acceptor 连接器
	connNum  syncutil.AtomicUint32 // 当前连接数
	state    syncutil.AtomicInt32  // connector 当前状态
}

// 新建1个 Connector 对象
func NewConnector(addrs *model.TLaddr, opt *model.TConnectorOpt) model.IConnector {
	// 参数效验
	if nil != opt.Check() {
		return nil
	}

	// 地址检查？

	// 创建组件
	cntor := &Connector{
		name:  model.C_CPT_NAME_CONNECTOR,
		laddr: addrs,
		opts:  opt,
	}

	// 创建 Acceptor
	aptor := newAcceptor(opt.AcceptorName, cntor)
	cntor.acceptor = aptor

	return cntor
}

// 运行 Connector [IComponent 接口]
func (this *Connector) Name() string {
	return this.name
}

// 运行 Connector [IComponent 接口]
func (this *Connector) Run() {
	// 启动 acceptor
	this.acceptor.Run()
}

// 停止 Connector [IComponent 接口]
func (this *Connector) Run() {
	// 停止 acceptor
	this.acceptor.Stop()
}

// 收到1个新的 Tcp 连接对象
func (this *Connector) OnNewTcpConn(conn net.Conn) {
	// 不符合 tcp 连接对象
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return
	}

	// 配置 iO 参数
	tcpConn.SetWriteBuffer(this.opts.TcpConnOpts.WriteBufferSize)
	tcpConn.SetReadBuffer(this.opts.TcpConnOpts.ReadBufferSize)
	tcpConn.SetNoDelay(this.opts.TcpConnOpts.NoDelay)
}

// 收到1个新的 websocket 连接对象
func (this *Connector) OnNewWsConn(wsconn *websocket.Conn) {

}

// 关闭所有连接
func (this *Connector) CloseAllConn() {

}

// 创建 session 对象
func (this *Connector) createSession(netconn net.Conn, isWebSocket bool) {

}
