// /////////////////////////////////////////////////////////////////////////////
// dispatcher 连接对象

package dispatcher

import (
	"net"

	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // session 库
	"github.com/zpab123/world/state"   // 状态管理
)

// dispatcher 连接对象
type DispatcherConn struct {
	stateMgr  *state.StateManager      // 状态管理
	worldConn *network.WorldConnection // world 引擎连接对象
}

// 新建1个 DispatcherConn
func NewDispatcherConn(netconn net.Conn, opt *network.TWorldConnOpts) *DispatcherConn {
	// 创建组件
	st := state.NewStateManager()

	socket := &network.Socket{
		Conn: netconn,
	}
	wc := network.NewWorldConnection(socket, opt)

	// 创建 DispatcherConn
	dc := &DispatcherConn{
		stateMgr:  st,
		worldConn: ws,
	}

	// 修改为初始化状态
	dc.stateMgr.SetState(state.C_STATE_INIT)

	return dc
}

// 启动 DispatcherConn
func (this *DispatcherConn) Run() {
	// 接收线程

	// 发送线程
}
