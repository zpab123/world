// /////////////////////////////////////////////////////////////////////////////
// tcp 连接器

package acceptor

import (
	"net"

	"github.com/zpab123/world/consts"             // 全局常量
	"github.com/zpab123/world/model"              // 全局模型
	"github.com/zpab123/world/utils"              // 工具库
	"github.com/zpab123/world/worldnet"           // 网络库
	"github.com/zpab123/world/worldnet/connector" // 连接器
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

func init() {
	// 注册创建函数
	connector.RegisterCreator(newTcpConnector)
}

// /////////////////////////////////////////////////////////////////////////////
// tcpAcceptor 对象

// tcp 接收器
type tcpAcceptor struct {
	name                      string                  // 连接器名字
	laddr                     model.TLaddr            // 地址集合
	opts                      *model.TTcpSocketOpts   // 配置参数
	State                                             // 对象继承：运行状态操作
	socketMgr                 model.ITcpSocketManager // 符合 tcpsocket 连接管理接口的对象
	connector.ISessionManager                         // 接口继承：符合 SessionManager 接口对象的
	connector.TcpSocketOption                         // 对象继承：tcp socket io 参数配置
	connector.RecoverIoPanic                          // 对象继承: io 异常捕获
	connector.PacketManager                           // 对象继承：packet 消息管理
	listener                  net.Listener            // 侦听器
}

// 创建1个新的 tcpAcceptor 对象
func NewTcpAcceptor(addr *model.TLaddr, opt *model.TTcpSocketOpts, mgr model.ITcpSocketManager) worldnet.IConnector {
	// 创建对象
	aptor := &tcpAcceptor{
		name:      model.C_ACCEPTOR_NAME_TCP,
		laddr:     addr,
		opts:      opt,
		socketMgr: mgr,
	}

	// 配置基础数据
	//cntor.TcpSocketOption.Init()

	return aptor
}

// 异步侦听新连接 [worldnet.IConnector 接口]
func (this *tcpAcceptor) Run() {
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
		zplog.Errorf("tcp.tcpAcceptor 创建失败。错误=%v", err.Error())
		this.SetRunning(false)

		return
	}

	// 创建成功
	this.listener = ln.(net.Listener)
	zplog.Infof("tcp.tcpAcceptor 创建成功。监听地址=%s", this.GetListenAddress())

	// 侦听连接
	go this.accept()
}

// 停止侦听器 [worldnet.IConnector 接口]
func (this *tcpAcceptor) Stop() {
	// 非运行状态
	if this.IsRuning() {
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

	// 断开所有 Session
	this.CloseAllSession()

	// 等待线程结束 - 阻塞
	this.WaitAllStop()
}

// 获取 Connector 的类型 [worldnet.IConnector 接口]
func (this *tcpAcceptor) GetType() string {
	return connector.TYPE_TCP_CONNECTOR
}

// 获取监听成功的端口
func (this *tcpAcceptor) GetPort() int {
	if this.listener == nil {
		return 0
	}

	return this.listener.Addr().(*net.TCPAddr).Port
}

// 获取监听成功的地址
func (this *tcpAcceptor) GetListenAddress() string {
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
func (this *tcpAcceptor) accept() {
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
			zplog.Errorf("tcp.tcpAcceptor 接收新连接出现错误。错误=%v", err.Error())
			break
		}

		// 处理连接进入独立线程, 防止 accept 无法响应
		go this.socketMgr.OnNewTcpConn(conn)
	}

	// 监听异常
	this.SetRunning(false) // 设置为非运行状态
	this.EndStop()         // 结束停止
}
