// /////////////////////////////////////////////////////////////////////////////
// go 空文件模板

package dispatcher

// 分发客户端
type ClientProxy struct {
	addr string                // 地址
	opt  *TDispatcherClientOpt // 配置参数
}

// 新建1个 ClientProxy
func NewClientProxy(opt *TDispatcherClientOpt) *ClientProxy {

}

// 连接服务器
func (this *ClientProxy) connectServer() {
	var addr string
	conn, err := net.Dial("tcp", addr)

	if nil != err {

	}

	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetReadBuffer(this.opt.TcpConnOpt.ReadBufferSize)
	tcpConn.SetWriteBuffer(this.opt.TcpConnOpt.WriteBufferSize)
}
