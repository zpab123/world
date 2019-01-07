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
// tcpConnector 对象

// tcp 接收器
type tcpConnector struct {
	connector.TcpSocketOption              // 对象继承：tcp socket io 参数配置
	connector.State                        // 对象继承：运行状态操作
	connector.BaseInfo                     // 对象继承：基础信息
	listener                  net.Listener // 侦听器
}

// 创建1个新的 tcpConnector 对象
func NewTcpConnector() worldnet.IConnector {
	// 创建对象
	cntor := &tcpConnector{}

	// 配置基础数据
	cntor.TcpSocketOption.Init()

	return cntor
}

// 异步侦听新连接 [worldnet.IConnector 接口]
func (self *tcpConnector) Run() {
	// 阻塞，等到所有线程结束
	self.WaitAllStop()

	// 已经运行
	if self.IsRuning() {
		return
	}

	// 创建侦听器
	ln, err := utils.DetectPort(self.Address(), func(a *utils.Address, port int) (interface{}, error) {
		return net.Listen("tcp", a.HostPortString(port))
	})

	// 创建失败
	if nil != err {
		zplog.Errorf("创建 tcp.tcpConnector 失败，名字=%s；错误=%v", self.Name(), err.Error())
		self.SetRunning(false)
	}

	// 创建成功
	self.listener = ln.(net.Listener)
	//zplog.Infof("创建 tcp.tcpConnector 成功，名字=%s；监听地址=%s", self.Name(), self.ListenAddress())

	// 侦听连接
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

	// 等待线程结束 - 阻塞
	self.WaitAllStop()
}

// 获取类型的名字 [worldnet.IConnector 接口]
func (self *tcpConnector) TypeName() string {
	return consts.CONNECTOR_TYPE_TCP_ACCEPTOR
}
