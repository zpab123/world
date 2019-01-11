// /////////////////////////////////////////////////////////////////////////////
// websocket 连接管理

package ws

import (
	"github.com/zpab123/world/ifs"               // 全局接口库
	"github.com/zpab123/world/network/connector" // 连接器
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
	connector.AddrManager              // 对象继承： 监听地址管理
	listener              net.Listener // 侦听器
}

// 创建1个 wsAcceptor 对象
func newWsAcceptor(cntor ifs.IConnector) connector.IAcceptor {
	// 创建对象
	wsaptor := &wsAcceptor{}

	return wsaptor
}

// 启动 wsAcceptor [IAcceptor 接口]
func (this *wsAcceptor) Run() {

}

// 停止 wsAcceptor [IAcceptor 接口]
func (this *wsAcceptor) Stop() {

}
