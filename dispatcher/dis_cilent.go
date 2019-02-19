// /////////////////////////////////////////////////////////////////////////////
// 分发客户端

package dispatcher

import (
	"fmt"
	"net"

	"github.com/zpab123/world/config"  // 配置文件
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/session" // session 库
	"github.com/zpab123/zaplog"        // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// DispatcherClient

// 分发客户端
type DispatcherClient struct {
	name        string                  // 组件名字
	option      *TDispatcherClientOpt   // 配置参数
	gateServers []*TServerInfo          // gate 服务器配置信息
	sessionMgr  *session.SessionManager // session 管理对象
	connMgrs    []*DispatcherConnMgr    // 连接管理
}

// 新建1个 DispatcherClient
func NewDispatcherClient(opt *TDispatcherClientOpt) *DispatcherClient {
	// 参数效验
	if nil == opt {
		opt = NewTDispatcherClientOpt(nil)
	}

	// 创建组件
	sesMgr := session.NewSessionManager()

	servers := config.GetServerMap()
	gates := servers[model.C_SERVER_TYPE_GATE]

	gateNum := len(gates)
	if 0 == gateNum {
		zaplog.Fatal("创建 DispatcherClient 出现异常。gate 服务器参数获取失败")
	}
	mgrs := make([]*DispatcherConnMgr, gateNum)
	if gateNum > 0 {
		for index, gate := range gates {
			host := gate.Host
			port := gate.Port
			addr := fmt.Sprintf("%s:%d", host, port)

			mgrs[index] = NewDispatcherConnMgr(addr, sesMgr, opt)
		}
	}

	// 创建 DispatcherClient
	dc := &DispatcherClient{
		name:        C_COMPONENT_NAME_CLIENT,
		option:      opt,
		gateServers: gates,
		sessionMgr:  sesMgr,
		connMgrs:    mgrs,
	}

	return dc
}

// 启动
func (this *DispatcherClient) Run() {
	// 连接所有 DispatcherServer 服务器
	for _, mgr := range this.connMgrs {
		go mgr.Run()
	}

	// 主循环
}

// 关闭
func (this *DispatcherClient) Stop() {
	// 关闭所有 DispatcherServer 连接
}
