// /////////////////////////////////////////////////////////////////////////////
// tcp 连接器

package tcp

import (
	"net"

	"github.com/zpab123/world/consts"            // 全局常量
	"github.com/zpab123/world/network"           // 网络库
	"github.com/zpab123/world/network/connector" // 连接器
	"github.com/zpab123/world/utils"             // 工具库
	"github.com/zpab123/zplog"                   // 日志库
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
	connector.Address              // 对象继承： 监听地址信息
	listener          net.Listener // 侦听器
}

// 创建1个新的 tcpConnector 对象
func newTcpConnector() {
	// 创建对象
	cntor := &tcpConnector{}

	return cntor
}

// 启动 connector [IConnector 接口]
func (this *tcpConnector) Run() {
	// 创建侦听器
	f := func(addr *utils.Address, port int) (interface{}, error) {
		return net.Listen("tcp", addr.HostPortString(port))
	}
	ln, err := utils.DetectPort(f)

	// 创建失败
	if nil != err {
		zplog.Errorf("启动 tcpConnector 失败。错误=%v", err.Error())

		return
	}

	// 创建成功
	this.listener = ln.(net.Listener)
	zplog.Infof("启动 tcpConnector 成功。监听地址=%s", this.GetListenAddress())

	// 侦听连接
	go this.accept()
}

// 停止 connector [IConnector 接口]
func (this *tcpConnector) Stop() {
	// 关闭侦听器
	this.listener.Close()
}

// 获取 connector 类型，例如 tcp.Connector/udp.Acceptor [IConnector 接口]
func (this *tcpConnector) GetType() {
	return consts.NETWORK_CONNECTOR_TYPE_TCP
}

// 获取监听成功的端口
func (this *tcpConnector) GetPort() int {
	if this.listener == nil {
		return 0
	}

	return this.listener.Addr().(*net.TCPAddr).Port
}

// 获取监听成功的地址
func (this *tcpConnector) GetListenAddress() string {
	// 获取 host
	pos := strings.Index(this.GetAddress(), ":")
	if pos == -1 {
		return this.GetAddress()
	}
	host := this.GetAddress()[:pos]

	// 获取 port
	port := this.GetPort()

	return utils.JoinAddress(host, port)
}

// 侦听连接
func (this *tcpConnector) accept() {
	// 主循环
	for {
		// 接收新连接
		conn, err := this.listener.Accept()

		// 监听错误
		if nil != err {
			zplog.Errorf("tcpConnector 接收新连接出现错误，停止工作。错误=%v", err.Error())

			break
		}

		// 处理连接进入独立线程, 防止 accept 无法响应
		//go self.createSession(conn)
	}
}
