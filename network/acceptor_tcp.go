// /////////////////////////////////////////////////////////////////////////////
// tcp 连接器

package network

import (
	"net"
	"strings"
	"sync"

	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/world/state" // 状态管理
	"github.com/zpab123/world/utils" // 工具库
	"github.com/zpab123/zaplog"      // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// TcpAcceptor 对象

// tcp 接收器
type TcpAcceptor struct {
	name     string              // 连接器名字
	laddr    *TLaddr             // 地址集合
	stateMgr *state.StateManager // 状态管理
	connMgr  ITcpConnManager     // 符合 ITcpAcceptorManager 连接管理接口的对象
	listener net.Listener        // 侦听器
}

// 创建1个新的 TcpAcceptor 对象
func NewTcpAcceptor(addr *TLaddr, mgr ITcpConnManager) IAcceptor {
	// 参数效验
	if addr.TcpAddr == "" {
		return nil
	}

	if nil == mgr {
		zaplog.Warnf("创建 TcpAcceptor。连接管理对象为nil")
	}

	// 对象
	st := state.NewStateManager()

	// 创建 TcpAcceptor
	aptor := &TcpAcceptor{
		name:     C_ACCEPTOR_TYPE_TCP,
		laddr:    addr,
		stateMgr: st,
		connMgr:  mgr,
	}

	aptor.stateMgr.SetState(state.C_STATE_INIT)

	return aptor
}

// 异步侦听新连接 [IAcceptor 接口]
func (this *TcpAcceptor) Run() bool {
	// 状态效验
	s := this.stateMgr.GetState()
	if s != state.C_STATE_INIT && s != state.C_STATE_STOP {
		return false
	}

	// 创建侦听器
	ln, err := utils.DetectPort(this.laddr.TcpAddr, func(a *model.TAddress, port int) (interface{}, error) {
		return net.Listen("tcp", a.HostPortString(port))
	})

	// 创建失败
	if nil != err {
		zaplog.Fatalf("TcpAcceptor 启动失败。ip=%s，err=%v", this.laddr.TcpAddr, err.Error())

		return false
	}

	// 创建成功
	this.stateMgr.SetState(state.C_STATE_WORKING)

	this.listener = ln.(net.Listener)
	zaplog.Infof("TcpAcceptor 启动成功。ip=%s", this.GetListenAddress())

	// 侦听连接
	go this.accept()

	return true
}

// 停止侦听器 [IAcceptor 接口]
func (this *TcpAcceptor) Stop() bool {
	// 关闭侦听器
	this.listener.Close()

	return true
}

// 获取监听成功的端口
func (this *TcpAcceptor) GetPort() int {
	if this.listener == nil {
		return 0
	}

	return this.listener.Addr().(*net.TCPAddr).Port
}

// 获取监听成功的地址
func (this *TcpAcceptor) GetListenAddress() string {
	// 获取 host
	pos := strings.Index(this.laddr.TcpAddr, ":")
	if pos == -1 {
		return this.laddr.TcpAddr
	}
	host := this.laddr.TcpAddr[:pos]

	// 获取 port
	port := this.GetPort()

	return utils.JoinAddress(host, port)
}

// 侦听连接
func (this *TcpAcceptor) accept() {
	// 主循环
	for {
		// 接收新连接
		conn, err := this.listener.Accept()

		// 监听错误
		if nil != err {
			zaplog.Errorf("TcpAcceptor 接收新连接出现错误。err=%v", err.Error())
			break
		}

		// 处理连接进入独立线程, 防止 accept 无法响应
		if nil != this.connMgr {
			go this.connMgr.OnNewTcpConn(conn)
		}
	}

	this.stateMgr.SetState(state.C_STATE_STOP)
}
