// /////////////////////////////////////////////////////////////////////////////
// websocket 接收器

package network

import (
	"net"
	"net/http"

	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/world/utils" // 工具库
	"github.com/zpab123/zaplog"       // log 日志库
	"golang.org/x/net/websocket"     // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// wsAcceptor 对象

// websocket 接收器
type WsAcceptor struct {
	name         string         // 连接器名字
	laddr        *TLaddr        // 地址集合
	connMgr      IWsConnManager // websocket 连接管理
	listener     net.Listener   // 侦听器： 用于http服务器
	httpServer   *http.Server   // http 服务器
	wsListenAddr string         // 侦听成功的 websocket 地址
	certFile     string         // TLS加密文件
	keyFile      string         // TLS解密key
}

// 创建1个新的 wsAcceptor 对象
func NewWsAcceptor(addr *TLaddr, mgr IWsConnManager) IAcceptor {
	// 创建接收器
	aptor := &WsAcceptor{
		name:    C_ACCEPTOR_NAME_WEBSOCKET,
		laddr:   addr,
		connMgr: mgr,
	}

	return aptor
}

// 启动 wsAcceptor [IAcceptor 接口]
func (this *WsAcceptor) Run() bool {
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
		zaplog.Errorf("WsAcceptor 启动失败。err=%v", err.Error())
		return false
	}

	// 端口查找成功
	this.listener = ln.(net.Listener)
	//this.listener.Close() // 解除端口占用
	this.wsListenAddr = addrObj.String(wsPort)

	// 侦听新连接
	go this.accept()

	return true
}

// 停止 wsAcceptor [IAcceptor 接口]
func (this *WsAcceptor) Stop() bool {
	return true
}

// 侦听连接
func (this *WsAcceptor) liten() {

	// 创建 http 服务器

	// 设置 "/ws" 消息协议处理函数(客户端需要在url后面加上 /ws 路由)
	if nil != this.connMgr {
		http.Handle("/ws", websocket.Handler(this.connMgr.OnNewWsConn)) // 有新连接的时候，会调用 wsHandler 处理新连接
	}

	var err error // 错误信息
	zaplog.Infof("WsAcceptor 启动成功。ip=%s", this.wsListenAddr)

	// 开启服务器
	if this.certFile != "" && this.keyFile != "" {
		zaplog.Debugf("WsAcceptor 使用 TLS。 cert=%s, key=%s", this.certFile, this.keyFile)
		err = http.ListenAndServeTLS(this.wsListenAddr, this.certFile, this.keyFile, nil)
	} else {
		err = http.ListenAndServe(this.wsListenAddr, nil)
	}

	// 错误信息
	if nil != err {
		zaplog.Fatalf("WsAcceptor 启动失败。ip=%s，err=%s", this.wsListenAddr, err)
	}
}

// 侦听连接
func (this *WsAcceptor) accept() {
	// 创建 mux
	mux := http.NewServeMux()
	// 路由函数
	handler := websocket.Handler(this.connMgr.OnNewWsConn)
	//mux.HandleFunc("/ws", handler)
	mux.Handle("/ws", handler)

	// 创建 httpServer
	this.httpServer = &http.Server{
		Addr:    this.wsListenAddr,
		Handler: mux,
	}

	// 开启服务器
	var err error
	zaplog.Infof("WsAcceptor 启动成功。ip=%s", this.wsListenAddr)
	if this.certFile != "" && this.keyFile != "" {
		err = this.httpServer.ServeTLS(this.listener, this.certFile, this.keyFile)
	} else {
		err = this.httpServer.Serve(this.listener)
	}

	// 错误信息
	if nil != err {
		zaplog.Fatalf("WsAcceptor 启动失败。ip=%s，err=%s", this.wsListenAddr, err)
	}
}
