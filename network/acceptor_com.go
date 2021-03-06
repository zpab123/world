// /////////////////////////////////////////////////////////////////////////////
// 网络连接接收器： 组合模式，支持 tcp websocket 连接

package network

import (
	"net"
	"net/http"
	"strings"

	"github.com/pkg/errors"          // 异常库
	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/world/state" // 状态管理
	"github.com/zpab123/world/utils" // 工具库
	"github.com/zpab123/world/wderr" // 异常库
	"github.com/zpab123/zaplog"      // 日志库
	"golang.org/x/net/websocket"     // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// ComAcceptor 对象

// 支持 tcp websocket 连接
type ComAcceptor struct {
	stateMgr     *state.StateManager // 状态管理
	name         string              // 侦听器名字
	tcpListener  net.Listener        // tcp 侦听器
	wsListenAddr string              // 监听成功的 websocket
	certFile     string              // TLS加密文件
	keyFile      string              // TLS解密key
	laddr        *TLaddr             // 监听地址集合
	connMgr      IComConnManager     // 连接管理对象
	httpServer   *http.Server        // http 服务器
	wsListener   net.Listener        // websocket 侦听器
}

// 创建1个 ComAcceptor 对象
func NewComAcceptor(addr *TLaddr, mgr IComConnManager) (IAcceptor, error) {
	var err error

	// 参数效验
	ok := (addr.TcpAddr == "" || addr.WsAddr == "")
	if ok {
		err = errors.New("创建 ComAcceptor 失败。参数 TcpAddr WsAddr 为空")

		return nil, err
	}

	if nil == mgr {
		err = errors.New("创建 ComAcceptor 失败。参数 IComConnManager=nil")

		return nil, err
	}

	// 创建 StateManager
	sm := state.NewStateManager()

	// 创建 ComAcceptor
	comaptor := &ComAcceptor{
		stateMgr: sm,
		name:     C_ACCEPTOR_NAME_COM,
		laddr:    addr,
		connMgr:  mgr,
	}

	// 设置为初始化状态
	comaptor.stateMgr.SetState(state.C_INIT)

	return comaptor, nil
}

// 启动 Acceptor [IAcceptor 接口]
func (this *ComAcceptor) Run() error {
	var err error

	// 改变状态: 正在启动中
	if !this.stateMgr.SwapState(state.C_INIT, state.C_RUNING) {
		err = errors.Errorf("ComAcceptor 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_INIT, this.stateMgr.GetState())

		return err
	}

	// 启动 tcp 侦听
	err = this.runTcpListener()
	if nil != err {
		return err
	}

	// 启动 websocket 侦听
	err = this.runWsListener()
	if nil != err {
		return err
	}

	// 改变状态: 工作中
	if !this.stateMgr.SwapState(state.C_RUNING, state.C_WORKING) {
		zaplog.Errorf("ComAcceptor 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_RUNING, this.stateMgr.GetState())

		return err
	}

	return err
}

// 停止 Acceptor [IAcceptor 接口]
func (this *ComAcceptor) Stop() error {
	// 改变状态: 关闭中
	if !this.stateMgr.SwapState(state.C_WORKING, state.C_STOPING) {
		zaplog.Errorf("ComAcceptor 停止失败，状态错误。正确状态=%d，当前状态=%d", state.C_WORKING, this.stateMgr.GetState())

		return nil
	}

	// 关闭 tcp
	this.tcpListener.Close()

	// 关闭 websocket
	this.httpServer.Close()

	// 改变状态: 关闭完成
	if !this.stateMgr.SwapState(state.C_STOPING, state.C_STOPED) {
		zaplog.Errorf("ComAcceptor 停止失败。状态错误。正确状态=%d，当前状态=%d", state.C_STOPING, this.stateMgr.GetState())

		return nil
	}

	return nil
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
func (this *ComAcceptor) runTcpListener() error {
	var err error
	var ln interface{}

	// 创建侦听器
	f := func(addr *model.TAddress, port int) (interface{}, error) {
		return net.Listen("tcp", addr.HostPortString(port))
	}
	ln, err = utils.DetectPort(this.laddr.TcpAddr, f)

	// 创建失败
	if nil != err {
		return err
	}

	// 创建成功
	var ok bool
	this.tcpListener, ok = ln.(net.Listener)
	if !ok {
		err = errors.New("ComAcceptor 启动失败，创建 net.Listener 失败")

		return err
	}

	zaplog.Infof("ComAcceptor-tcp 启动成功。ip=%s", this.GetTcpAddr())

	// 侦听 tcp 连接
	go this.acceptTcpConn()

	return err
}

// 接收新的 tcp 连接
func (this *ComAcceptor) acceptTcpConn() {
	//  出现错误，关闭监听
	closeF := func() {
		zaplog.Error("ComAcceptor 侦听 tcp 新连接出现错误。关闭 ComAcceptor-tcp")

		this.tcpListener.Close()
	}

	defer closeF()

	// 监听新连接
	for {
		newConn, err := this.tcpListener.Accept()
		if nil != err {
			if wderr.IsTimeoutError(err) {
				continue
			} else {
				break
			}
		}

		// 开启新协程，处理新 tcp 连接
		if nil != this.connMgr {
			go this.connMgr.OnNewTcpConn(newConn)
		}
	}
}

// 启动 websocket 侦听
func (this *ComAcceptor) runWsListener() error {
	// 变量定义
	var (
		addrObj *model.TAddress // 地址变量
		wsPort  int             // 监听成功的 websocket 端口
		err     error           // 错误
		ln      interface{}     // 监听接口
	)

	// 查找1个 可用端口
	f := func(addr *model.TAddress, port int) (interface{}, error) {
		addrObj = addr
		wsPort = port

		return net.Listen("tcp", addr.HostPortString(port))
	}

	ln, err = utils.DetectPort(this.laddr.WsAddr, f)

	// 查找失败
	if nil != err {
		return err
	}

	// 端口查找成功
	this.wsListenAddr = addrObj.String(wsPort)
	this.wsListener = ln.(net.Listener)
	//this.wsListener.Close() // 解除端口占用

	// 侦听新连接
	go this.acceptWsConn()

	return err
}

// 接收新的 websocket 连接
func (this *ComAcceptor) acceptWsConn() {
	// 创建 mux
	mux := http.NewServeMux()
	handler := websocket.Handler(this.connMgr.OnNewWsConn)
	mux.Handle("/ws", handler)

	// 创建 httpServer
	this.httpServer = &http.Server{
		Addr:    this.wsListenAddr,
		Handler: mux,
	}

	// 开启服务器
	zaplog.Infof("ComAcceptor-websocket 启动成功。ip=%s", this.wsListenAddr)

	var err error
	if this.certFile != "" && this.keyFile != "" {
		err = this.httpServer.ServeTLS(this.wsListener, this.certFile, this.keyFile)
	} else {
		err = this.httpServer.Serve(this.wsListener)
	}

	// 错误信息
	if nil != err {
		zaplog.Fatalf("ComAcceptor-websocket 启动失败。ip=%s，err=%s", this.wsListenAddr, err)
	}
}

// 是否需要停止
func (this *ComAcceptor) needStop() bool {
	return this.stateMgr.GetState() == state.C_STOPING
}
