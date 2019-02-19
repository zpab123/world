// /////////////////////////////////////////////////////////////////////////////
// World 框架内部通信使用的客户端连接对象

package network

import (
	"net"

	"github.com/zpab123/world/state" // 状态管理
	"github.com/zpab123/zaplog"      // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// WorldConnClient

// World 框架内部通信使用的客户端连接对象
type WorldConnClient struct {
	addr     string               // 连接地址
	option   *TWorldConnClientOpt // 配置参数
	stateMgr *state.StateManager  // 状态管理
	tcpConn  *net.TCPConn         // tcp 连接对象
}

// 新建1个 WorldConnClient
func NewWorldConnClient(addr string, opt *TWorldConnClientOpt) *WorldConnClient {
	// 创建对象
	if nil == opt {
		opt = NewTWorldConnClientOpt()
	}

	st := state.NewStateManager()

	// 创建 WorldConnClient
	wc := &WorldConnClient{
		addr:     addr,
		option:   opt,
		stateMgr: st,
	}

	wc.stateMgr.SetState(state.C_STATE_INIT)

	return wc
}

// 连接服务器
func (this *WorldConnClient) Connect() {
	this.connectByTcp()
}

// 发送握手

// 发送心跳

// 其他

// 通过 tcp 连接
func (this *WorldConnClient) connectByTcp() error {
	// 创建 tcp 连接
	conn, err := net.Dial("tcp", this.addr)
	if nil != err {
		zaplog.Errorf("WorldConnClient 连接服务器失败。ip=%s", this.addr)

		return err
	}

	// io 参数
	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetReadBuffer(this.option.TcpConnOpt.ReadBufferSize)
	tcpConn.SetWriteBuffer(this.option.TcpConnOpt.WriteBufferSize)
	tcpConn.SetNoDelay(this.option.TcpConnOpt.NoDelay)

	this.tcpConn = tcpConn

	zaplog.Debugf("WorldConnClient 连接服务器成功。ip=%s", this.addr)

	return nil
}
