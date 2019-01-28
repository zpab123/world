// /////////////////////////////////////////////////////////////////////////////
// 网络连接接收器： 组合模式，支持 tcp websocket 连接

package network

import (
	"net"
	"net/http"
	"strings"

	"github.com/zpab123/syncutil"    // 原子变量
	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/world/state" // 状态管理
	"github.com/zpab123/world/utils" // 工具库
	"github.com/zpab123/zplog"       // 日志库
	"golang.org/x/net/websocket"     // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// ComAcceptor 对象

// 支持 tcp websocket 连接
type ComAcceptor struct {
	*state.StateManager                           // 对象继承： 状态管理
	name                string                    // 侦听器名字
	tcpListener         net.Listener              // tcp 侦听器
	wsListenAddr        string                    // 监听成功的 websocket
	certFile            string                    // TLS加密文件
	keyFile             string                    // TLS解密key
	laddr               *model.TLaddr             // 监听地址集合
	acceptorMgr         model.IComAcceptorManager // 接收器管理对象
	httpServer          *http.Server              // http 服务器
	wsListener          net.Listener              // websocket 侦听器
}

// 创建1个 ComAcceptor 对象
func NewComAcceptor(addr *model.TLaddr, mgr model.IComAcceptorManager) model.IAcceptor {
	// 创建 StateManager
	sm := state.NewStateManager()

	// 创建 ComAcceptor
	comaptor := &ComAcceptor{
		StateManager: sm,
		name:         model.C_ACCEPTOR_NAME_COM,
		laddr:        addr,
		acceptorMgr:  mgr,
	}

	// 设置为初始化状态
	mulaptor.SetState(model.C_STATE_INIT)

	return comaptor
}

// 启动 Acceptor [IAcceptor 接口]
func (this *ComAcceptor) Run() bool {
	// 改变状态: 正在启动中
	if !this.SwapState(model.C_STATE_INIT, model.C_STATE_RUNING) {
		zplog.Errorf("ComAcceptor 启动失败，状态错误。正确状态=%d，当前状态=%d", model.C_STATE_INIT, this.GetState())

		return false
	}

	// 启动 tcp 侦听
	if !this.runTcpListener() {
		return false
	}

	// 启动 websocket 侦听
	if !this.runWsListener() {
		return false
	}

	// 改变状态: 工作中
	if !this.SwapState(model.C_STATE_RUNING, model.C_STATE_WORKING) {
		zplog.Errorf("ComAcceptor 启动失败，状态错误。正确状态=%d，当前状态=%d", model.C_STATE_RUNING, this.GetState())

		return false
	}

	return true
}

// 停止 Acceptor [IAcceptor 接口]
func (this *ComAcceptor) Stop() bool {
	// 改变状态: 关闭中
	if !this.SwapState(model.C_STATE_WORKING, model.C_STATE_STOPING) {
		zplog.Errorf("ComAcceptor 停止失败，状态错误。正确状态=%d，当前状态=%d", model.C_STATE_WORKING, this.GetState())

		return false
	}

	// 关闭 tcp
	this.tcpListener.Close()

	// 关闭 websocket
	this.httpServer.Close()

	// 改变状态: 关闭完成
	if !this.SwapState(model.C_STATE_STOPING, model.C_APP_STATE_STOP) {
		zplog.Errorf("ComAcceptor 停止失败。状态错误。正确状态=%d，当前状态=%d", model.C_STATE_STOPING, this.GetState())

		return false
	}

	return true
}

// 获取 侦听成功的 tcp 端口
func (this *ComAcceptor) GetTcpPort() int {
	if nil == this.tcpListener {
		return 0
	}

	port := this.tcpListener.Addr().(*net.TCPAddr).Port

	return port
}

// 获取 侦听成功的 tcpAddr 端口
func (this *ComAcceptor) GetTcpAddr() string {
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
func (this *ComAcceptor) runTcpListener() bool {
	// 创建侦听器
	f := func(addr *model.TAddress, port int) (interface{}, error) {
		return net.Listen("tcp", addr.HostPortString(port))
	}
	ln, err := utils.DetectPort(this.laddr.TcpAddr, f)

	// 创建失败
	if nil != err {
		zplog.Errorf("ComAcceptor-tcp 启动失败。err=%v", err.Error())

		return false
	}

	// 创建成功
	this.tcpListener = ln.(net.Listener)
	zplog.Infof("ComAcceptor-tcp 启动成功。ip=%s", this.GetTcpAddr())

	// 侦听 tcp 连接
	go this.acceptTcpConn()

	return true
}

// 接收新的 tcp 连接
func (this *ComAcceptor) acceptTcpConn() {
	//  出现错误，关闭监听
	closeF := func() {
		zplog.Error("ComAcceptor 侦听 tcp 新连接出现错误。关闭 ComAcceptor-tcp")
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

		// 开启新线程，处理新 tcp 连接
		go this.acceptorMgr.OnNewTcpConn(newConn)
	}
}

// 启动 websocket 侦听
func (this *ComAcceptor) runWsListener() bool {
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
		zplog.Errorf("ComAcceptor-websocket 启动失败。err=%v", err.Error())
		return false
	}

	// 端口查找成功
	this.wsListenAddr = addrObj.String(wsPort)
	this.wsListener = ln.(net.Listener)
	//this.wsListener.Close() // 解除端口占用

	// 侦听新连接
	go this.acceptWsConn()
	//go this.acceptWebsocket()

	return true
}

// 接收新的 websocket 连接
func (this *ComAcceptor) acceptWsConn() {
	// 创建 mux
	mux := http.NewServeMux()
	handler := websocket.Handler(this.acceptorMgr.OnNewWsConn)
	mux.Handle("/ws", handler)

	// 创建 httpServer
	this.httpServer = &http.Server{
		Addr:    this.wsListenAddr,
		Handler: mux,
	}

	// 开启服务器
	var err error
	zplog.Infof("ComAcceptor-websocket 启动成功。ip=%s", this.wsListenAddr)
	if this.certFile != "" && this.keyFile != "" {
		err = this.httpServer.ServeTLS(this.wsListener, this.certFile, this.keyFile)
	} else {
		err = this.httpServer.Serve(this.wsListener)
	}

	// 错误信息
	if nil != err {
		zplog.Fatalf("ComAcceptor-websocket 启动失败。ip=%s，err=%s", this.wsListenAddr, err)
	}
}

// 接收新的 websocket 连接
func (this *ComAcceptor) acceptWebsocket() {
	// 设置 "/ws" 消息协议处理函数(客户端需要在url后面加上 /ws 路由)
	http.Handle("/ws", websocket.Handler(this.acceptorMgr.OnNewWsConn)) // 有新连接的时候，会调用 OnNewWsConn 处理新连接

	// 侦听新连接
	var err error
	zplog.Infof("ComAcceptor-websocket 启动成功。ip=%s", this.wsListenAddr)
	if "" != this.keyFile && "" != this.certFile {
		zplog.Debugf("ComAcceptor-websocket 使用 TLS。 key=%s, cert=%s", this.keyFile, this.certFile)
		err = http.ListenAndServeTLS(this.wsListenAddr, this.certFile, this.keyFile, nil)
	} else {
		err = http.ListenAndServe(this.wsListenAddr, nil)
	}

	// 错误信息
	if nil != err {
		zplog.Fatalf("ComAcceptor-websocket 启动失败。ip=%s，错误=%s", this.wsListenAddr, err)
	}
}

// 是否需要停止
func (this *ComAcceptor) needStop() bool {
	return this.GetState() == model.C_STATE_STOPING
}
