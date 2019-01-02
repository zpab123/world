// /////////////////////////////////////////////////////////////////////////////
// 游戏服务组件

package component

import (
	"net"

	"github.com/zpab123/syncutil"      // 原子操作工具
	"github.com/zpab123/world/consts"  // 全局常量
	"github.com/zpab123/world/model"   // 全局结构体
	"github.com/zpab123/world/network" // 网络库
	"golang.org/x/net/websocket"       // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 常量
const (
	_maxConnNum uint32 = 100000 // 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// connector 对象

// 网络连接对象，支持 websocket tcp
type Connector struct {
	name      string                 // 组件名字
	maxConn   uint32                 // 最大连接数量，超过此数值后，不在接收新连接
	connNum   syncutil.AtomicUint32  // 当前连接数
	state     syncutil.AtomicInt32   // connector 当前状态
	config    *model.ConnectorConfig // 配置参数
	tcpServer *network.TcpServer     // tcp 服务器
	wsServer  *network.WsServer      // websocket 服务器
}

// 新建1个 Connector 对象
func NewConnector(parameter *model.ConnectorConfig) *Connector {
	// 参数效验
	if nil != parameter.Check() {
		return nil
	}

	// 创建对象
	server := &Connector{
		name:    consts.COMPONENT_NAME_CONNECTOR,
		maxConn: _maxConnNum,
		config:  parameter,
	}

	// 数据初始化
	server.init()

	return server
}

// 组件名字 [IComponent 实现]
func (this *Connector) Name() string {
	return this.name
}

// 运行 Connector 组件 [IComponent 实现]
func (this *Connector) Run() {
	// 启动 tcp 服务
	if nil != this.tcpServer {
		this.tcpServer.Run()
	}

	// 启动 websocket 服务
	if nil != this.wsServer {
		this.wsServer.Run()
	}
}

// 停止运行 [IComponent 实现]
func (this *Connector) Stop() {

}

// 设置最大连数
func (this *Connector) SetMaxConn(num uint32) {
	this.maxConn = num
}

// 收到1个新的 websocket 连接对象
func (this *Connector) OnWsConn(wsconn *websocket.Conn) {

}

// 收到1个新的 Tcp 连接对象
func (this *Connector) OnTcpConn(conn net.Conn) {

}

// 初始化 Connector 数据
func (this *Connector) init() {
	// 最大连接数
	if this.config.MaxConn > 0 {
		this.maxConn = this.config.MaxConn
	}

	// 创建 tcp 服务器
	if this.config.TcpAddr != "" {
		this.tcpServer = network.NewTcpServer(this.config.TcpAddr, this)
	}

	// 创建 websocket 服务
	if this.config.WsAddr != "" {
		this.wsServer = network.NewWsServer(this.config.WsAddr, this)
	}
}
