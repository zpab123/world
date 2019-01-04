// /////////////////////////////////////////////////////////////////////////////
// websocket 接收器

package ws

import (
	"net"
	"net/http"

	"github.com/gorilla/websocket" // websocket 库
)

// /////////////////////////////////////////////////////////////////////////////
// wsAcceptor 对象

// websocket 接收器
type wsAcceptor struct {
	listener net.Listener       // 侦听器
	sv       *http.Server       // http 服务器
	upgrader websocket.Upgrader // ws 协议处理
}

// 创建1个新的 tcpAcceptor 对象
func newWsAcceptor() *wsAcceptor {
	// 创建接收器
	acceptor := &wsAcceptor{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	return acceptor
}

// 启动 wsAcceptor
func (self *wsAcceptor) Run() {

}
