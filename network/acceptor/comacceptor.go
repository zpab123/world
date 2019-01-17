// /////////////////////////////////////////////////////////////////////////////
// 网络连接接收器： 组合模式，支持 tcp websocket 连接

package acceptor

import (
	"net"
	"net/http"
	"strings"

	"github.com/zpab123/world/model"             // 全局模型
	"github.com/zpab123/world/network"           // 网络库
	"github.com/zpab123/world/network/connector" // 连接器
	"github.com/zpab123/world/network/socket"    // socket 库
	"github.com/zpab123/world/session"           // socket会话
	"github.com/zpab123/world/utils"             // 工具库
	"github.com/zpab123/zplog"                   // 日志库
	"golang.org/x/net/websocket"                 // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// comAcceptor 对象

// 支持 tcp websocket 连接
type comAcceptor struct {
	network.TCPSocketOption                      // 对象继承： socket 基础参数管理
	tcpListener             net.Listener         // tcp 侦听器
	wsListenAddr            string               // 监听成功的 websocket
	certFile                string               // TLS加密文件
	keyFile                 string               // TLS解密key
	laddr                   *model.TLaddr        // 监听地址集合
	sessionMgr              model.ISessionManage // session 管理
}

// 创建1个 comAcceptor 对象
func NewComAcceptor(addr model.TLaddr, mgr model.ISessionManage) model.IAcceptor {
	// 创建 comAcceptor
	comaptor := &comAcceptor{
		sessionMgr: mgr,
		laddr:      addr,
	}

	return comaptor
}

// 启动 Acceptor [IAcceptor 接口]
func (this *comAcceptor) Run() {
	// 启动 tcp 侦听
	this.runTcpListener()

	// 启动 websocket 侦听
	this.runWsListener()
}

// 停止 Acceptor [IAcceptor 接口]
func (this *comAcceptor) Stop() {

}

// 获取 侦听成功的 tcp 端口
func (this *comAcceptor) GetTcpPort() int {
	if nil == this.tcpListener {
		return 0
	}

	port := this.tcpListener.Addr().(*net.TCPAddr).Port

	return port
}

// 获取 侦听成功的 tcpAddr 端口
func (this *comAcceptor) GetTcpAddr() string {
	// 获取 host
	pos := strings.Index(this.laddr.TcpAddr, ":")
	if pos == -1 {
		return this.laddr.TcpAddr
	}
	host := this.laddr.TcpAddr[:pos]

	// 获取 port
	port := this.GetTcpPort()

	return utils.JoinAddress(host, port)
}

// 启动 tcp 侦听
func (this *comAcceptor) runTcpListener() {
	// 创建侦听器
	f := func(addr *model.TAddress, port int) (interface{}, error) {
		return net.Listen("tcp", addr.HostPortString(port))
	}
	ln, err := utils.DetectPort(this.laddr.TcpAddr, f)

	// 创建失败
	if nil != err {
		zplog.Errorf("tcpAcceptor-tcp 启动失败。错误=%v", err.Error())

		return
	}

	// 创建成功
	this.tcpListener = ln.(net.Listener)
	zplog.Infof("comAcceptor-tcp 启动成功。ip=%s", this.GetTcpAddr())

	// 侦听 tcp 连接
	go this.acceptTcpConn()
}

// 接收新的 tcp 连接
func (this *comAcceptor) acceptTcpConn() {
	//  出现错误，关闭监听
	closeF := func() {
		zplog.Error("comAcceptor 侦听 tcp 新连接出现错误。关闭 comAcceptor-tcp")
		this.tcpListener.Close()
	}
	defer closeF()

	// 监听新连接
	for {
		newConn, err := this.tcpListener.Accept()
		if nil != err {
			if utils.IsTimeoutError(err) {
				continue
			} else {
				break
			}
		}

		// 开启新线程 处理新 tcp 连接
		go this.onNewTcpConn(newConn)
	}
}

// 收到1个新的 tcp 连接
func (this *comAcceptor) onNewTcpConn(conn net.Conn) {
	// 日志
	zplog.Debugf("收到1个新的 tcp 连接。客户端ip=%s", conn.RemoteAddr())

	// 配置 io 参数
	this.ApplySocketOption(conn)

	// 创建 session
	this.createSession(conn)
}

// 启动 websocket 侦听
func (this *comAcceptor) runWsListener() {
	// 变量定义
	var (
		addrObj *model.TAddress // 地址变量
		wsPort  int             // 监听成功的 websocket 端口
	)

	// 查找1个 可用端口
	f := func(addr *model.TAddress, port int) (interface{}, error) {
		addrObj = addr
		wsPort = port
		return net.Listen("tcp", addr.HostPortString(port))
	}

	ln, err := utils.DetectPort(this.laddr.WsAddr, f)

	// 查找失败
	if nil != err {
		zplog.Errorf("comAcceptor-websocket 启动失败。错误=%v", err.Error())
		return
	}

	// 端口查找成功
	this.wsListenAddr = addrObj.String(wsPort)
	ln.(net.Listener).Close() // 解除端口占用

	// 设置 "/ws" 消息协议处理函数(客户端需要在url后面加上 /ws 路由)
	http.Handle("/ws", websocket.Handler(this.onNewWsConn)) // 有新连接的时候，会调用 onNewWsConn 处理新连接

	// 侦听新连接
	go func() {
		var err error
		zplog.Infof("comAcceptor-websocket 启动成功。ip=%s", this.wsListenAddr)
		if "" == this.keyFile && "" == this.certFile {
			err = http.ListenAndServe(this.wsListenAddr, nil)
		} else {
			zplog.Infof("comAcceptor-websocket 使用 TLS。 key=%s, cert=%s", this.keyFile, this.certFile)
			err = http.ListenAndServeTLS(this.wsListenAddr, this.certFile, this.keyFile, nil)
		}

		// 错误信息
		if nil != err {
			zplog.Fatalf("comAcceptor-websocket 启动失败。ip=%s，错误=%s", this.wsListenAddr, err)
		}
	}()
}

// 收到1个新的 websocket 连接
func (this *comAcceptor) onNewWsConn(wsConn *websocket.Conn) {
	// 日志
	zplog.Debugf("收到1个新的 websocket 连接。客户端ip=%s", wsConn.RemoteAddr())

	// 设置为接收二进制数据
	wsConn.PayloadType = websocket.BinaryFrame

	// 创建 session
	this.createSession(wsConn)
}

// 创建 session
func (this *comAcceptor) createSession(conn net.Conn) {
	// 创建 session
	bufferSocket := network.NewBufferSocket(conn)
	ses := session.NewClientSession(bufferSocket, this.sessionMgr, nil)

	// 通知管理器
	if nil != this.sessionMgr {
		this.sessionMgr.OnNewSession(ses)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// private api
