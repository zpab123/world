// /////////////////////////////////////////////////////////////////////////////
// tcp 连接器

package tcp

import (
	"net"

	"github.com/zpab123/world/consts"            // 全局常量
	"github.com/zpab123/world/network"           // 网络库
	"github.com/zpab123/world/network/connector" // 连接器
	"github.com/zpab123/world/utils"             // 工具库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

func init() {
	// 注册创建函数
	connector.RegisterCreator(newTcpConnector)
}

// /////////////////////////////////////////////////////////////////////////////
// tcpConnector 对象

// tcp 接收器
type tcpConnector struct {
	connector.BaseInfo // 对象继承： 基础信息
}

// 创建1个新的 tcpConnector 对象
func newTcpConnector() {
	// 创建对象
	cntor := &tcpConnector{}

	return cntor
}

// 启动 connector [IConnector 接口]
func (this *tcpConnector) Run() {

}

// 停止 connector [IConnector 接口]
func (this *tcpConnector) Stop() {

}

// 获取 connector 类型，例如 tcp.Connector/udp.Acceptor [IConnector 接口]
func (this *tcpConnector) GetType() {
	return consts.NETWORK_CONNECTOR_TYPE_TCP
}
