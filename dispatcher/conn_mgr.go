// /////////////////////////////////////////////////////////////////////////////
// dispatcher 客户端连接管理

package dispatcher

import (
	"github.com/zpab123/world/network" // 网络库
)

// 分发客户端
type DispatcherConnMgr struct {
	addr            string                   // 服务器地址
	option          *TDispatcherClientOpt    // 配置参数
	worldConnClient *network.WorldConnClient // 连接对象
}

// 新建1个 DispatcherConnMgr
func NewDispatcherConnMgr(addr string, opt *TDispatcherClientOpt) *DispatcherConnMgr {
	// 创建组件
	wc := network.NewWorldConnClient(addr, opt.WorldConnClientOpt)

	// 创建对象
	dc := &DispatcherConnMgr{
		addr:            addr,
		option:          opt,
		worldConnClient: wc,
	}

	return dc
}

// 启动 DispatcherConnMgr
func (this *DispatcherConnMgr) Connect() {
	// 连接服务器
	this.worldConnClient.Connect()

	// 主循环
	for {

	}
}
