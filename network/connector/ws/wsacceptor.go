// /////////////////////////////////////////////////////////////////////////////
// websocket 连接管理

package ws

import (
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"               // websocket 库
	"github.com/zpab123/world/ifs"               // 全局接口库
	"github.com/zpab123/world/model"             // 常用数据类型
	"github.com/zpab123/world/network/connector" // 连接器
	"github.com/zpab123/world/utils"             // 工具库
	"github.com/zpab123/zplog"                   // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 初始化函数
func init() {
	connector.RegisterAcceptor(connector.CONNECTOR_TYPE_WEBSOCKET, newWsAcceptor)
}

// /////////////////////////////////////////////////////////////////////////////
// wsAcceptor 对象

// websocket 连接对象
type wsAcceptor struct {
	connector.AddrManager                    // 对象继承： 监听地址管理
	upgrader              websocket.Upgrader // 暂时不知道 upgrader
	listener              net.Listener       // 侦听器： 用于可用端口查找
	httpServer            *http.Server       // http 服务器
	certfile              string             // 加密文件
	keyfile               string             // key
	connector             ifs.IConnector     // connector 对象
}

// 创建1个 wsAcceptor 对象
func newWsAcceptor(cntor ifs.IConnector) ifs.IAcceptor {
	// 检查函数
	origin := func(r *http.Request) bool {
		return true
	}

	// 创建 Upgrader
	uper := websocket.Upgrader{
		CheckOrigin: origin,
	}

	// 创建对象
	wsaptor := &wsAcceptor{
		upgrader:  uper,
		connector: cntor,
	}

	return wsaptor
}

// 启动 wsAcceptor [IAcceptor 接口]
func (this *wsAcceptor) Run() {
	// 变量定义
	var (
		addrObj *model.Address // 地址变量
		err     error          // 错误
	)

	// 查找1个 可用端口
	f := func(addr *model.Address, port int) (interface{}, error) {
		addrObj = addr

		return net.Listen("tcp", addr.HostPortString(port))
	}

	ln, err := utils.DetectPort(this.GetAddr().WsAddr, f)

	// 查找失败
	if nil != err {
		zplog.Errorf("启动 wsAcceptor 失败。错误=%v", err.Error())
		return
	}

	// 端口查找成功
	this.listener = ln.(net.Listener)

	// 创建多路转接器
	mux := http.NewServeMux()

	// http handler 函数
	handler := func(w http.ResponseWriter, r *http.Request) {
		// 新的连接
		conn, err := this.upgrader.Upgrade(w, r, nil)
		if nil != err {
			zplog.Errorf("wsAcceptor 接收新连接出现错误。错误=%v", err.Error())
			return
		}

		// 创建 socket
		newWsSocket(conn, this.connector)
	}

	// HandleFunc 注册一个处理器函数 handler 和对应的模式 pattern。
	if "" == addrObj.Path {
		addrObj.Path = "/"
	}
	mux.HandleFunc(addrObj.Path, handler)

	// 创建 http 服务器
	this.httpServer = &http.Server{
		Addr:    addrObj.HostPortString(this.GetPort()),
		Handler: mux,
	}

	// 启动 http 服务器
	go func() {
		zplog.Infof("启动 wsAcceptor 成功。监听地址=%s", addrObj.String(this.GetPort()))

		if this.certfile != "" && this.keyfile != "" {
			err = this.httpServer.ServeTLS(this.listener, this.certfile, this.keyfile)
		} else {
			err = this.httpServer.Serve(this.listener)
		}

		if nil != err {
			zplog.Errorf("启动 wsAcceptor 失败。错误=%v", err.Error())
		}
	}()
}

// 停止 wsAcceptor [IAcceptor 接口]
func (this *wsAcceptor) Stop() {

}

// 获取监听成功的端口
func (this *wsAcceptor) GetPort() int {
	if this.listener == nil {
		return 0
	}

	return this.listener.Addr().(*net.TCPAddr).Port
}
