// /////////////////////////////////////////////////////////////////////////////
// 游戏服务组件

package component

import (
	"net"

	"github.com/zpab123/syncutil"                // 原子操作工具
	"github.com/zpab123/world/consts"            // 全局常量
	"github.com/zpab123/world/model"             // 全局结构体
	"github.com/zpab123/world/network"           // 网络库
	"github.com/zpab123/world/network/connector" // 网络连接库
	"github.com/zpab123/zplog"                   // log 库
	"golang.org/x/net/websocket"                 // websocket 库
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
	name      string                  // 组件名字
	connNum   syncutil.AtomicUint32   // 当前连接数
	state     syncutil.AtomicInt32    // connector 当前状态
	opt       *connector.ConnectorOpt // 配置参数
	laddr     *model.Laddr            // 监听地址集合
	connector network.IConnector      // 某种类型的 connector 连接器
}

// 新建1个 Connector 对象
func NewConnector(addrs *model.Laddr, param *connector.ConnectorOpt) *Connector {
	// 参数效验
	if nil != parameter.Check() {
		return nil
	}

	// 创建 connector
	cntor := connector.NewConnector(addrs, param)

	// 创建组件
	cpt := &Connector{
		name:  consts.COMPONENT_NAME_CONNECTOR,
		laddr: addrs,
		opt:   param,
	}

	// 保存 Connector
	cpt.connector = cntor

	// 数据初始化
	cpt.init()

	return cpt
}

// 组件名字 [IComponent 实现]
func (this *Connector) Name() string {
	return this.name
}

// 运行 Connector 组件 [IComponent 实现]
func (this *Connector) Run() {
	// 启动 tcp 服务
	if nil != this.tcpServer {
		zplog.Debugf("开启 tcp 服务")
		go this.tcpServer.Run()
	}

	// 启动 websocket 服务
	if nil != this.wsServer {
		go this.wsServer.Run()
	}
}

// 停止运行 [IComponent 实现]
func (this *Connector) Stop() {

}

// 收到1个新的 websocket 连接对象
func (this *Connector) OnWsConn(wsconn *websocket.Conn) {
	// 设置网路数据接收格式
	zplog.Debugf("收到1个新的 websocket 连接，客户端地址为=%s", wsconn.RemoteAddr())
	wsconn.PayloadType = websocket.BinaryFrame //按照二进制方式接收数据

	// 创建统一连接并接收数据
}

// 收到1个新的 Tcp 连接对象
func (this *Connector) OnTcpConn(conn net.Conn) {
	// 打印日志
	zplog.Debugf("收到1个新的 TCP 连接，客户端地址为=%s", conn.RemoteAddr())

	// 将 conn 接口变量 转化为 *net.TCPConn 对象 （conn 需要完全符合 *net.TCPConn）
	tcpConn, ok := conn.(*net.TCPConn)
	if false == ok {
		zplog.Errorf("收到错误的Tcp连接")
		conn.Close()
		return
	}

	// 设置 buffer 参数
	tcpConn.SetWriteBuffer(consts.CONNECTOR_SOCKET_WRITE_BUF_SIZE)
	tcpConn.SetReadBuffer(consts.CONNECTOR_SOCKET_READ_BUF_SIZE)
	tcpConn.SetNoDelay(consts.CONNECTOR_SOCKET_NO_DELAY)

	// 创建统一连接并接收数据
}

// 初始化 Connector 数据
func (this *Connector) init() {
	// 设置配置参数
	this.setConfig()

	// 最大连接数
	if this.config.MaxConn <= 0 {
		this.config.MaxConn = _maxConnNum
	}

	// 创建 tcp 服务器
	if this.laddr.TcpAddr != "" {
		this.tcpServer = network.NewTcpServer(this.laddr.TcpAddr, this)
	}

	// 创建 websocket 服务
	if this.laddr.WsAddr != "" {
		this.wsServer = network.NewWsServer(this.laddr.WsAddr, this)
	}
}

// 设置默认参数
func (this *Connector) setConfig() {
	// 已经设置
	if nil != this.config {
		return
	}

	// 创建 config
	this.config = &model.ConnectorConfig{
		MaxConn: _maxConnNum,
	}
}

// 创建 session 组件
func (this *Connector) createSession(conn net.Conn) {
	// 达到最大连接
	if this.connNum.Load() >= this.config.MaxConn {
		conn.Close()
		zplog.Warnf("Connector 组件达到最大连接数量=%d，断开新连接", this.config.MaxConn)
		return
	}

	// 主动停止接收新连接 -- 后续添加

	// 创建 session

	// session 线程循环
}

// 获取 connector
func (this *Connector) getConnector(typeName string) network.IConnector {

}
