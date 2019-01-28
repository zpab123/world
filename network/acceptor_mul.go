// /////////////////////////////////////////////////////////////////////////////
// 同时支持 tcp websocket

package network

import (
	"github.com/zpab123/world/model" // 全局模型
	"github.com/zpab123/world/state" // 状态管理
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// mulAcceptor 对象

// 同时支持 tcp websocket 的对象
type MulAcceptor struct {
	*state.StateManager                 // 对象继承： 状态管理
	name                string          // 连接器名字
	laddr               *model.TLaddr   // 地址集合
	tcpAcceptor         model.IAcceptor // tcpAcceptor 对象
	wsAcceptor          model.IAcceptor // wsAcceptor 对象
}

// 创建1个 mulAcceptor 对象
func NewMulAcceptor(addr *model.TLaddr, mgr model.IMulConnManager) model.IAcceptor {
	// 创建 StateManager
	sm := state.NewStateManager()

	// 创建 tcpAcceptor
	tcpaptor := NewTcpAcceptor(addr, mgr)

	// 创建 wsAcceptor
	wsaptor := NewWsAcceptor(addr, mgr)

	// 创建对象
	mulaptor := &MulAcceptor{
		StateManager: sm,
		name:         model.C_ACCEPTOR_NAME_MUL,
		laddr:        addr,
		tcpAcceptor:  tcpaptor,
		wsAcceptor:   wsaptor,
	}

	// 设置为初始化状态
	mulaptor.SetState(model.C_STATE_INIT)

	return mulaptor
}

// 启动 mulAcceptor [IAcceptor 接口]
func (this *MulAcceptor) Run() bool {
	// 状态效验
	if this.GetState() != model.C_STATE_INIT {
		return false
	}

	// 改变状态: 正在启动中
	this.SetState(model.C_STATE_RUNING)

	// 添加启动线程数量
	this.AddRunGo(2)

	// 启动 tcp
	this.tcpAcceptor.Run()

	// 启动 websocket
	this.wsAcceptor.Run()

	// 阻塞
	this.RunWait()

	// 启动完成

	return true
}

// 停止 mulAcceptor [IAcceptor 接口]
func (this *MulAcceptor) Stop() bool {
	// 停止 tcp
	this.tcpAcceptor.Stop()

	// 停止 websocket
	this.wsAcceptor.Stop()

	return true
}
