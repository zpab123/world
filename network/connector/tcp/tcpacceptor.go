// /////////////////////////////////////////////////////////////////////////////
// tcp 连接器

package tcp

import (
	"net"
	"strings"

	"github.com/zpab123/world/model"             // 常用数据类型
	"github.com/zpab123/world/network/connector" // 连接器
	"github.com/zpab123/world/utils"             // 工具库
	"github.com/zpab123/zplog"                   // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

func init() {
	// 注册创建函数
	connector.RegisterAcceptor(connector.CONNECTOR_TYPE_TCP, newTcpAcceptor)
}

// /////////////////////////////////////////////////////////////////////////////
// tcpAcceptor 对象

// tcp 接收器
type tcpAcceptor struct {
	connector.AddrManager                      // 对象继承： 监听地址管理
	connector.TCPSocketOption                  // 对象继承： socket 基础参数管理
	listener                  net.Listener     // 侦听器
	connector                 model.IConnector // connector 对象
}

// 创建1个新的 tcpAcceptor 对象
func newTcpAcceptor(cntor model.IConnector) model.IAcceptor {
	// 创建对象
	aptor := &tcpAcceptor{
		connector: cntor,
	}

	// 参数初始化
	aptor.Init()

	return aptor
}

// 启动 tcpAcceptor [IAcceptor 接口]
func (this *tcpAcceptor) Run() {
	// 创建侦听器
	f := func(addr *model.Address, port int) (interface{}, error) {
		return net.Listen("tcp", addr.HostPortString(port))
	}
	ln, err := utils.DetectPort(this.GetAddr().TcpAddr, f)

	// 创建失败
	if nil != err {
		zplog.Errorf("启动 tcpAcceptor 失败。错误=%v", err.Error())

		return
	}

	// 创建成功
	this.listener = ln.(net.Listener)
	zplog.Infof("启动 tcpAcceptor 成功。监听地址=%s", this.GetListenAddress())

	// 侦听连接
	go this.accept()
}

// 停止 tcpAcceptor [IAcceptor 接口]
func (this *tcpAcceptor) Stop() {
	// 关闭侦听器
	this.listener.Close()
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
	pos := strings.Index(this.GetAddr().TcpAddr, ":")
	if pos == -1 {
		return this.GetAddr().TcpAddr
	}
	host := this.GetAddr().TcpAddr[:pos]

	// 获取 port
	port := this.GetPort()

	return utils.JoinAddress(host, port)
}

// 侦听连接
func (this *tcpAcceptor) accept() {
	// 主循环
	for {
		// 接收新连接
		conn, err := this.listener.Accept()

		// 监听错误
		if nil != err {
			zplog.Errorf("tcpAcceptor 接收新连接出现错误，停止工作。错误=%v", err.Error())

			break
		}

		// 处理连接进入独立线程, 防止 accept 无法响应
		go this.onNewConn(conn)
	}
}

// 收到1个新的 socket 连接
func (this *tcpAcceptor) onNewConn(conn net.Conn) {
	// 配置 io 参数
	this.ApplySocketOption(conn)

	// 创建 socket 对象
	socket := newtcpSocket(conn, this.connector)

	// 通知 Connector 组件
	this.connector.OnNewSocket(socket)
}
