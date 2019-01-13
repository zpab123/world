// /////////////////////////////////////////////////////////////////////////////
// websocket 客户端 <-> 服务器 通信

package mul

import (
	"net"

	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
)

// /////////////////////////////////////////////////////////////////////////////
// mulSocket 对象

// websocket 管理
type mulSocket struct {
	tcpConn   net.Conn         // Socket原始连接
	connector model.IConnector // connector 组件
}

// 创建1个新的 mulSocket 对象
func newMulSocket(conn net.Conn, cntor model.IConnector) interface{} {
	// 创建 socket
	socket := &mulSocket{
		tcpConn:   conn,
		connector: cntor,
	}

	return socket
}

// 刷新缓冲区
func (this *mulSocket) Flush() error {
	return nil
}
