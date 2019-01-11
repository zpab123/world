// /////////////////////////////////////////////////////////////////////////////
// websocket 客户端 <-> 服务器 通信

package ws

import (
	"github.com/gorilla/websocket" // websocket 库
	"github.com/zpab123/world/ifs" // 全局接口库
)

// /////////////////////////////////////////////////////////////////////////////
// wsSocket 对象

// websocket 管理
type wsSocket struct {
	wsConn    *websocket.Conn // websocket 原始连接
	connector ifs.IConnector  // connector 组件
}

// 创建1个新的 wsSocket 对象
func newWsSocket(conn *websocket.Conn, cntor ifs.IConnector) ifs.ISocket {
	// 创建 socket
	socket := &wsSocket{
		wsConn:    conn,
		connector: cntor,
	}

	return socket
}

// 刷新缓冲区
func (this *wsSocket) Flush() error {
	return nil
}
