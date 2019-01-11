// /////////////////////////////////////////////////////////////////////////////
// 同时支持 tcp websocket

package mul

import (
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"               // websocket 库
	"github.com/zpab123/world/ifs"               // 全局接口库
	"github.com/zpab123/world/network/connector" // 连接器
	"github.com/zpab123/world/utils"             // 工具库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 初始化函数
func init() {
	connector.RegisterAcceptor(connector.CONNECTOR_TYPE_MUL, newMulAcceptor)
}

// /////////////////////////////////////////////////////////////////////////////
// mulAcceptor 对象

// 同时支持 tcp websocket 的对象
type mulAcceptor struct {
	connector.AddrManager                        // 对象继承： 监听地址管理
	connector.TCPSocketOption                    // 对象继承： socket 基础参数管理
	upgrader                  websocket.Upgrader // 暂时不知道 upgrader
	tcpListener               net.Listener       // tcp 侦听器
	wsListener                net.Listener       // websocket 侦听器
	certfile                  string             // 加密文件
	keyfile                   string             // key
	connector                 ifs.IConnector     // connector 组件
	httpServer                *http.Server       // http 服务器
}

// 创建1个 mulAcceptor 对象
func newMulAcceptor(cntor ifs.IConnector) ifs.IAcceptor {
	// 创建 tcpAcceptor
	//tcpaptor := connector.NewAcceptor()

	// 创建对象
	mulaptor := &mulAcceptor{
		connector: cntor,
	}

	//  设置地址
	//mulaptor.SetAddr()

	return mulaptor
}

// 启动 mulAcceptor [IAcceptor 接口]
func (this *mulAcceptor) Run() {
	// 启动 tcp
	this.runTcp()

	// 启动 websocket
	this.runWebsocket()
}

// 停止 mulAcceptor [IAcceptor 接口]
func (this *mulAcceptor) Stop() {

}

// 获取监听成功的 tcp 端口
func (this *mulAcceptor) GetTcpPort() int {
	if this.tcpListener == nil {
		return 0
	}

	return this.tcpListener.Addr().(*net.TCPAddr).Port
}

// 获取监听成功的 websocket 端口
func (this *mulAcceptor) GetWsPort() int {
	if this.wsListener == nil {
		return 0
	}

	return this.wsListener.Addr().(*net.TCPAddr).Port
}

// 获取监听成功的 tcp 地址
func (this *mulAcceptor) GetTcpAddr() string {
	// 获取 host
	pos := strings.Index(this.GetAddr().TcpAddr, ":")
	if pos == -1 {
		return this.GetAddr().TcpAddr
	}
	host := this.GetAddr().TcpAddr[:pos]

	// 获取 port
	port := this.GetTcpPort()

	return utils.JoinAddress(host, port)
}

// 获取监听成功的 websocket 地址
func (this *mulAcceptor) GetWsAddr() string {
	// 获取 host
	pos := strings.Index(this.GetAddr().WsAddr, ":")
	if pos == -1 {
		return this.GetAddr().WsAddr
	}
	host := this.GetAddr().WsAddr[:pos]

	// 获取 port
	port := this.GetWsPort()

	return utils.JoinAddress(host, port)
}

// 启动 tcp 侦听器
func (this *mulAcceptor) runTcp() {

}

// 启动 websocket 侦听器
func (this *mulAcceptor) runWebsocket() {

}
