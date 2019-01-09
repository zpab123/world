// /////////////////////////////////////////////////////////////////////////////
// tcp 连接器

package tcp

import (
	"net"

	"github.com/zpab123/world/consts"             // 全局常量
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
// tcpConnector 对象

// tcp 接收器
type tcpConnector struct {
	connector.ISessionManager              // 接口继承：符合 SessionManager 接口对象的
	connector.TcpSocketOption              // 对象继承：tcp socket io 参数配置
	connector.State                        // 对象继承：运行状态操作
	connector.BaseInfo                     // 对象继承：基础信息
	connector.RecoverIoPanic               // 对象继承: io 异常捕获
	connector.PacketManager                // 对象继承：packet 消息管理
	listener                  net.Listener // 侦听器
}

// 创建1个新的 tcpConnector 对象
func newTcpConnector() worldnet.IConnector {
	// 创建对象
	cntor := &tcpConnector{
		ISessionManager: new(connector.SessionManager),
	}

	// 配置基础数据
	cntor.TcpSocketOption.Init()

	return cntor
}

// 异步侦听新连接 [worldnet.IConnector 接口]
func (self *tcpConnector) Run() {
	// 阻塞，等到所有线程结束
	self.WaitAllStop()

	// 正在运行
	if self.IsRuning() {
		return
	}

	// 创建侦听器
	ln, err := utils.DetectPort(self.GetAddress(), func(a *utils.Address, port int) (interface{}, error) {
		return net.Listen("tcp", a.HostPortString(port))
	})

	// 创建失败
	if nil != err {
		zplog.Errorf("tcp.tcpConnector 创建失败：名字=%s；错误=%v", self.GetName(), err.Error())
		self.SetRunning(false)

		return
	}

	// 创建成功
	self.listener = ln.(net.Listener)
	zplog.Infof("tcp.tcpConnector 创建成功：名字=%s；监听地址=%s", self.GetName(), self.GetListenAddress())

	// 侦听连接
	go self.accept()
}

// 停止侦听器 [worldnet.IConnector 接口]
func (self *tcpConnector) Stop() {
	// 非运行状态
	if self.IsRuning() {
		return
	}

	// 正在停止
	if self.IsStopping() {
		return
	}

	// 开始停止
	self.StartStop()

	// 关闭侦听器
	self.listener.Close()

	// 断开所有 Session
	self.CloseAllSession()

	// 等待线程结束 - 阻塞
	self.WaitAllStop()
}

// 获取 Connector 的类型 [worldnet.IConnector 接口]
func (self *tcpConnector) GetType() string {
	return connector.TYPE_TCP_CONNECTOR
}

// 获取监听成功的端口
func (self *tcpConnector) GetPort() int {
	if self.listener == nil {
		return 0
	}

	return self.listener.Addr().(*net.TCPAddr).Port
}

// 获取监听成功的地址
func (self *tcpConnector) GetListenAddress() string {
	// 获取 host
	pos := strings.Index(self.GetAddress(), ":")
	if pos == -1 {
		return self.GetAddress()
	}
	host := self.GetAddress()[:pos]

	// 获取 port
	port := self.GetPort()

	return utils.JoinAddress(host, port)
}

// 侦听连接
func (self *tcpConnector) accept() {
	// 设置为运行状态
	self.SetRunning(true)

	// 主循环
	for {
		// 接收新连接
		conn, err := self.listener.Accept()

		// 正在停止
		if self.IsStopping() {
			break
		}

		// 监听错误
		if nil != err {
			zplog.Errorf("tcp.tcpConnector 接收新连接出现错误：名字=%s 错误=%v", self.GetName(), err.Error())
			break
		}

		// 处理连接进入独立线程, 防止 accept 无法响应
		go self.createSession(conn)
	}

	// 监听异常
	self.SetRunning(false) // 设置为非运行状态
	self.EndStop()         // 结束停止
}

// 创建新的 tcpSession
func (self *tcpConnector) createSession(conn net.Conn) {
	// 设置 conn io 参数
	self.SetSocketOption(conn)

	// 创建 session
	ses := newSession(conn, self, nil)

	// 启动 session
	ses.Run()
}
