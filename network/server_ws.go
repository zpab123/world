// /////////////////////////////////////////////////////////////////////////////
// WebSocket 服务器

package network

import (
	"net/http"

	"github.com/zpab123/zplog"   // log 日志库
	"golang.org/x/net/websocket" // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// wsServer 对象

// websocket server 对象
type WsServer struct {
	listenAddr string     // 监听地址(格式 -> 127.0.0.1:6532)
	service    IWsService // 符合 IWsService 接口的对象
	certFile   string     // TLS加密文件
	keyFile    string     // TLS解密key
}

// 创建1个新的 WsServer 对象
func NewWsServer(laddr string, svc IWsService) *WsServer {

	// 创建对象
	ws := &WsServer{
		listenAddr: laddr,
		service:    svc,
	}

	// 数据初始化
	// ws.init()

	return ws
}

// 运行游戏服务器
func (ws *WsServer) Run() error {
	// 开启 websocket 服务器
	err := runWsServer(ws.listenAddr, ws.service.OnWsConn, ws.certFile, ws.keyFile)

	return err
}

// 停止服务器
func (ws *WsServer) Stop() error {
	return nil
}

// 设置 TLS 参数
func (ws *WsServer) SetTls(certFile string, keyFile string) {
	ws.certFile = certFile
	ws.keyFile = keyFile
}

// /////////////////////////////////////////////////////////////////////////////
// 私有api

// 开启 websocket 服务器
//
// listenAddr=监听地址(格式 -> 127.0.0.1:6532);wsHandler=websocket连接对象处理函数；certFile=加密文件；keyFile=解密key
func runWsServer(listenAddr string, wsHandler func(ws *websocket.Conn), certFile string, keyFile string) error {
	// log
	zplog.Infof("websocket 服务开启成功，ip=%s", listenAddr)
	if "" != keyFile || "" != certFile {
		zplog.Infof("websocket 使用 TLS, key=%s, cert=%s", keyFile, certFile)
	}

	// 设置 "/ws" 消息协议处理函数(客户端需要在url后面加上 /ws 路由)
	if wsHandler != nil {
		http.Handle("/ws", websocket.Handler(wsHandler)) // 有新连接的时候，会调用 wsHandler 处理新连接
	}

	// 侦听新连接
	var err error
	if "" == keyFile && "" == certFile {
		err = http.ListenAndServe(listenAddr, nil)
	} else {
		err = http.ListenAndServeTLS(listenAddr, certFile, keyFile, nil)
	}

	// 错误信息
	if nil != err {
		zplog.Fatalf("websocket 开启失败，ip=%s，错误=%s", listenAddr, err)
	}

	return err
}
