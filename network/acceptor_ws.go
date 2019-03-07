// /////////////////////////////////////////////////////////////////////////////
// websocket 接收器

package network

import (
	"net"
	"net/http"

	"github.com/pkg/errors"          // 异常库
	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/world/utils" // 工具库
	"github.com/zpab123/zaplog"      // log 日志库
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
func NewWsAcceptor(addr *TLaddr, mgr IWsConnManager) (IAcceptor, error) {
	var err error

	// 参数效验
	if addr.WsAddr == "" {
		err = errors.New("创建 WsAcceptor 失败。参数 WsAddr 为空")

		return nil, err
	}

	if nil == mgr {
		err = errors.New("创建 WsAcceptor 失败。参数 IWsConnManager=nil")

		return nil, err
	}

	// 创建接收器
	aptor := &WsAcceptor{
		name:    C_ACCEPTOR_NAME_WS,
		laddr:   addr,
		connMgr: mgr,
	}

	return aptor, nil
}

// 启动 wsAcceptor [IAcceptor 接口]
func (this *WsAcceptor) Run() error {
	// 变量定义
	var (
		addrObj *model.TAddress // 地址变量
		wsPort  int             // 监听成功的 websocket 端口
		ln      interface{}     // Listener
		err     error           // 错误
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
	var ok bool
	this.listener, ok = ln.(net.Listener)
	if !ok {
		err = errors.New("WsAcceptor 启动失败，创建 net.Listener 失败")

		return err
	}

	//this.listener.Close() // 解除端口占用
	this.wsListenAddr = addrObj.String(wsPort)

	// 侦听新连接
	go this.accept()

	return nil
}

// 停止 wsAcceptor [IAcceptor 接口]
func (this *WsAcceptor) Stop() error {

	return this.httpServer.Close()
}

// 侦听连接
func (this *WsAcceptor) accept() {
	// 创建 mux
	mux := http.NewServeMux()
	handler := websocket.Handler(this.connMgr.OnNewWsConn) // 路由函数
	mux.Handle("/ws", handler)                             // 客户端需要在url后面加上 /ws 路由

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
		zaplog.Fatalf("WsAcceptor 停止服务。ip=%s，err=%s", this.wsListenAddr, err)
	}
}
