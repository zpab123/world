// /////////////////////////////////////////////////////////////////////////////
// 分发客户端

package dispatcher

import (
	"net"
)

// /////////////////////////////////////////////////////////////////////////////
// go 空文件模板

// 分发客户端
type DispatcherClient struct {
	name    string                // 组件名字
	opt     *TDispatcherClientOpt // 配置参数
	clients []*ClientProxy        // 连接切片
}

// 新建1个 DispatcherClient
func NewDispatcherClient(opt *TDispatcherClientOpt) *DispatcherClient {
	dc := &DispatcherClient{
		name: C_COMPONENT_NAME_CLIENT,
	}

	return dc
}

// 启动
func (this *DispatcherClient) Run() {
	// 连接所有 DispatcherServer 服务器

	// 主循环
}

// 关闭
func (this *DispatcherClient) Stop() {
	// 关闭所有 DispatcherServer 连接
}
