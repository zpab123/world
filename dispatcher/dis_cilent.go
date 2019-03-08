// /////////////////////////////////////////////////////////////////////////////
// 分发客户端

package dispatcher

import (
	"fmt"

	"github.com/pkg/errors"            // 异常库
	"github.com/zpab123/world/config"  // 配置文件
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// DispatcherClient

// 分发客户端
type DispatcherClient struct {
	name          string                // 组件名字
	option        *TDispatcherClientOpt // 配置参数
	disServerInfo []*config.TServerInfo // dispatcher 服务器配置信息
	connectors    []*Connector          // 连接对象数组
}

// 新建1个 DispatcherClient
func NewDispatcherClient(opt *TDispatcherClientOpt) (dc model.IComponent, err error) {
	// 参数效验
	if nil == opt {
		opt = NewTDispatcherClientOpt(nil)
	}

	// 创建连接管理
	servers := config.GetServerMap()
	disServers := servers[model.C_SERVER_TYPE_DIS]

	disNum := len(disServers)
	cts := make([]*Connector, disNum)

	if 0 == disNum {
		err = errors.New("创建 DispatcherClient 出现异常。 dispatcher 服务器数量为0")

		return
	} else if disNum > 0 {
		for key, dis := range disServers {
			host := dis.Host
			port := dis.Port
			ipAddr := fmt.Sprintf("%s:%d", host, port)

			addr := &network.TLaddr{
				TcpAddr: ipAddr,
				WsAddr:  ipAddr,
				UdpAddr: ipAddr,
				KcpAddr: ipAddr,
			}

			cts[key] = NewConnector(addr, opt)
		}
	}

	// 创建 DispatcherClient
	dc = &DispatcherClient{
		name:          C_COMPONENT_NAME_CLIENT,
		option:        opt,
		disServerInfo: disServers,
		connectors:    cts,
	}

	return
}

// 获取组件名字
func (this *DispatcherClient) Name() string {
	return this.name
}

// 启动
func (this *DispatcherClient) Run() error {
	// 连接所有 DispatcherServer 服务器
	if nil != this.connectors && len(this.connectors) > 0 {
		for _, ct := range this.connectors {
			ct.Connect()
		}
	}

	// 主循环

	return nil
}

// 关闭
func (this *DispatcherClient) Stop() error {
	// 关闭所有 DispatcherServer 连接

	return nil
}
