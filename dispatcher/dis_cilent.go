// /////////////////////////////////////////////////////////////////////////////
// 分发客户端

package dispatcher

import (
	"fmt"

	"github.com/zpab123/world/config"  // 配置文件
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/session" // session 库
	"github.com/zpab123/zaplog"        // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// DispatcherClient

// 分发客户端
type DispatcherClient struct {
	name          string                  // 组件名字
	option        *TDispatcherClientOpt   // 配置参数
	disServerInfo []*config.TServerInfo   // dispatcher 服务器配置信息
	connMgr       []*DispatcherConnMgr    // 连接管理
	sessionMgr    *session.SessionManager // session 管理对象
}

// 新建1个 DispatcherClient
func NewDispatcherClient(opt *TDispatcherClientOpt) model.IComponent {
	// 参数效验
	if nil == opt {
		opt = NewTDispatcherClientOpt(nil)
	}

	// 创建组件
	sesMgr := session.NewSessionManager()

	// 创建连接管理
	servers := config.GetServerMap()
	disServers := servers[model.C_SERVER_TYPE_DIS]

	disNum := len(disServers)
	connMgrs := make([]*DispatcherConnMgr, disNum)

	if 0 == disNum {
		zaplog.Fatal("创建 DispatcherClient 出现异常。 dispatcher 服务器数量为0")
	} else if disNum > 0 {
		for key, dis := range disServers {
			host := dis.Host
			port := dis.Port
			addr := fmt.Sprintf("%s:%d", host, port)

			connMgrs[key] = NewDispatcherConnMgr(addr, opt)
		}
	}

	// 创建 DispatcherClient
	dc := &DispatcherClient{
		name:          C_COMPONENT_NAME_CLIENT,
		option:        opt,
		disServerInfo: disServers,
		connMgr:       connMgrs,
		sessionMgr:    sesMgr,
	}

	return dc
}

// 获取组件名字
func (this *DispatcherClient) Name() string {
	return this.name
}

// 启动
func (this *DispatcherClient) Run() bool {
	// 连接所有 DispatcherServer 服务器
	if nil != this.connMgr && len(this.connMgr) > 0 {
		for _, mgr := range this.connMgr {
			mgr.Connect()
		}
	}

	// 主循环

	return true
}

// 关闭
func (this *DispatcherClient) Stop() bool {
	// 关闭所有 DispatcherServer 连接

	return true
}
