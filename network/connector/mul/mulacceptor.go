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
	connector.AddrManager                // 对象继承： 监听地址管理
	connector             ifs.IConnector // connector 组件
	tcpAcceptor           ifs.IAcceptor  // tcpAcceptor 对象
	wsAcceptor            ifs.IAcceptor  // wsAcceptor 对象
}

// 创建1个 mulAcceptor 对象
func newMulAcceptor(cntor ifs.IConnector) ifs.IAcceptor {
	// 创建 tcpAcceptor
	tcpaptor := connector.NewAcceptor(connector.CONNECTOR_TYPE_TCP, cntor)

	// 创建 wsAcceptor
	wsaptor := connector.NewAcceptor(connector.CONNECTOR_TYPE_WEBSOCKET, cntor)

	// 创建对象
	mulaptor := &mulAcceptor{
		tcpAcceptor: tcpaptor,
		wsAcceptor:  wsaptor,
	}

	return mulaptor
}

// 启动 mulAcceptor [IAcceptor 接口]
func (this *mulAcceptor) Run() {
	// 启动 tcp
	this.tcpAcceptor.Run()

	// 启动 websocket
	this.wsAcceptor.Run()
}

// 停止 mulAcceptor [IAcceptor 接口]
func (this *mulAcceptor) Stop() {

}
