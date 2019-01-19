// /////////////////////////////////////////////////////////////////////////////
// tcp 连接器

package network

import (
	"net"

	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/world/utils" // 工具库
	"github.com/zpab123/zplog"       //日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// TcpAcceptor 对象

// tcp 接收器
type TcpAcceptor struct {
	state                             // 对象继承：运行状态操作
	name      string                  // 连接器名字
	laddr     model.TLaddr            // 地址集合
	socketMgr model.ITcpSocketManager // 符合 tcpsocket 连接管理接口的对象
	listener  net.Listener            // 侦听器
}

// 创建1个新的 TcpAcceptor 对象
func NewTcpAcceptor(addr *model.TLaddr, mgr model.ITcpSocketManager) model.IAcceptor {
	// 创建对象
	aptor := &TcpAcceptor{
		name:      model.C_ACCEPTOR_NAME_TCP,
		laddr:     addr,
		socketMgr: mgr,
	}

	return aptor
}

// 异步侦听新连接 [IAcceptor 接口]
func (this *TcpAcceptor) Run() {
	// 阻塞，等到所有线程结束
	this.WaitAllStop()

	// 正在运行
	if this.IsRuning() {
		return
	}

	// 创建侦听器
	ln, err := utils.DetectPort(this.laddr.TcpAddr, func(a *model.TAddress, port int) (interface{}, error) {
		return net.Listen("tcp", a.HostPortString(port))
	})

	// 创建失败
	if nil != err {
		zplog.Fatalf("TcpAcceptor 启动失败。ip=%s，err=%v", this.laddr.TcpAddr, err.Error())
		this.SetRunning(false)

		return
	}

	// 创建成功
	this.listener = ln.(net.Listener)
	zplog.Infof("TcpAcceptor 启动成功。ip=%s", this.GetListenAddress())

	// 侦听连接
	go this.accept()
}

// 停止侦听器 [IAcceptor 接口]
func (this *TcpAcceptor) Stop() {
	// 非运行状态
	if !this.IsRuning() {
		return
	}

	// 正在停止
	if this.IsStopping() {
		return
	}

	// 开始停止
	this.StartStop()

	// 关闭侦听器
	this.listener.Close()

	// 断开所有连接
	this.socketMgr.CloseAllTcpConn()

	// 等待线程结束 - 阻塞
	this.WaitAllStop()
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
	// 设置为运行状态
	this.SetRunning(true)

	// 主循环
	for {
		// 接收新连接
		conn, err := this.listener.Accept()

		// 正在停止
		if this.IsStopping() {
			break
		}

		// 监听错误
		if nil != err {
			zplog.Errorf("TcpAcceptor 接收新连接出现错误。err=%v", err.Error())
			break
		}

		// 处理连接进入独立线程, 防止 accept 无法响应
		go this.socketMgr.OnNewTcpConn(conn)
	}

	// 侦听线程退出
	this.SetRunning(false) // 设置为非运行状态
	this.EndStop()         // 结束停止
}
