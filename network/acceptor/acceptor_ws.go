// /////////////////////////////////////////////////////////////////////////////
// websocket 接收器

package acceptor

import (
	"net"
	"net/http"

	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/zplog"       // log 日志库
	"golang.org/x/net/websocket"     // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// wsAcceptor 对象

// websocket 接收器
type WsAcceptor struct {
	state                                // 对象继承：运行状态操作
	name         string                  // 连接器名字
	laddr        model.TLaddr            // 地址集合
	websocketMgr model.IWebsocketManager // websocket 连接管理
	listener     net.Listener            // 侦听器： 用于http服务器
	httpServer   *http.Server            // http 服务器
	wsListenAddr string                  // 侦听成功的 websocket 地址
	certFile     string                  // TLS加密文件
	keyFile      string                  // TLS解密key
}

// 创建1个新的 wsAcceptor 对象
func NewWsAcceptor(addr *model.TLaddr, mgr model.IWebsocketManager) model.IAcceptor {
	// 创建接收器
	aptor := &WsAcceptor{
		name:         model.C_ACCEPTOR_NAME_WEBSOCKET,
		laddr:        addr,
		websocketMgr: mgr,
	}

	return aptor
}

// 启动 wsAcceptor [IAcceptor 接口]
func (this *WsAcceptor) Run() {
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
		zplog.Errorf("WsAcceptor 启动失败。err=%v", err.Error())
		return
	}

	// 端口查找成功
	this.listener = ln.(net.Listener)
	this.listener.Close() // 解除端口占用
	this.wsListenAddr = addrObj.String(wsPort)

	// 侦听新连接
	go this.accept()

}

// 停止 wsAcceptor [IAcceptor 接口]
func (this *WsAcceptor) Stop() {

}

// 侦听连接
func (this *WsAcceptor) accept() {

	// 创建 http 服务器

	// 设置 "/ws" 消息协议处理函数(客户端需要在url后面加上 /ws 路由)
	if nil != this.websocketMgr {
		http.Handle("/ws", websocket.Handler(this.websocketMgr.OnNewWsConn)) // 有新连接的时候，会调用 wsHandler 处理新连接
	}

	var err error // 错误信息
	zplog.Infof("WsAcceptor 启动成功。ip=%s", this.wsListenAddr)

	// 开启服务器
	if this.certFile != "" && this.keyFile != "" {
		zplog.Debugf("WsAcceptor 使用 TLS。 cert=%s, key=%s", this.certFile, this.keyFile)
		err = http.ListenAndServeTLS(this.wsListenAddr, certFile, keyFile, nil)
	} else {
		err = http.ListenAndServe(this.wsListenAddr, nil)
	}

	// 错误信息
	if nil != err {
		zplog.Fatalf("WsAcceptor 启动失败。ip=%s，err=%s", this.wsListenAddr, err)
	}
}
