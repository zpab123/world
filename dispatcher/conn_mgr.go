// /////////////////////////////////////////////////////////////////////////////
// dispatcher 客户端连接管理

package dispatcher

import (
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
)

// 分发客户端
type DispatcherConnMgr struct {
	addr    string                // 服务器地址
	option  *TDispatcherClientOpt // 配置参数
	disConn *DispatcherConn       // 连接对象
}

// 新建1个 DispatcherConnMgr
func NewDispatcherConnMgr(addr string, opt *TDispatcherClientOpt) *DispatcherConnMgr {
	dc := &DispatcherConnMgr{
		addr:   addr,
		option: opt,
	}

	return dc
}

// 启动 DispatcherConnMgr
func (this *DispatcherConnMgr) Run() {
	conn, err := net.Dial("tcp", this.addr)

	if nil != err {

	}

	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetReadBuffer(this.option.TcpConnOpt.ReadBufferSize)
	tcpConn.SetWriteBuffer(this.option.TcpConnOpt.WriteBufferSize)

	// 创建代理
	clientProxy = NewDispatcherConn(conn, this.option.WorldConnOpts)
}
