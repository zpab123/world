// /////////////////////////////////////////////////////////////////////////////
// 同时支持 tcp websocket

package mul

import (
	"github.com/zpab123/world/ifs"               // 全局接口库
	"github.com/zpab123/world/network/connector" // 连接器
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 初始化函数
func init() {
	connector.RegisterAcceptor(connector.CONNECTOR_TYPE_MUL, newMulAcceptor)
}

// /////////////////////////////////////////////////////////////////////////////
// mulAcceptor 对象

// 同时支持 tcp websocket 的对象
type mulAcceptor struct {
	connector.AddrManager              // 对象继承： 监听地址管理
	listener              net.Listener // 侦听器
}

// 创建1个 mulAcceptor 对象
func newMulAcceptor(cntor ifs.IConnector) connector.IAcceptor {
	// 创建对象
	wsaptor := &mulAcceptor{}

	return wsaptor
}

// 启动 mulAcceptor [IAcceptor 接口]
func (this *mulAcceptor) Run() {

}

// 停止 mulAcceptor [IAcceptor 接口]
func (this *mulAcceptor) Stop() {

}
