// /////////////////////////////////////////////////////////////////////////////
// 同时支持 tcp websocket

package network

import (
	"sync"

	"github.com/zpab123/world/state" // 状态管理
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// mulAcceptor 对象

// 同时支持 tcp websocket 的对象
type MulAcceptor struct {
	*state.StateManager                // 对象继承： 状态管理
	name                string         // 连接器名字
	laddr               *TLaddr        // 地址集合
	tcpAcceptor         IAcceptor      // tcpAcceptor 对象
	wsAcceptor          IAcceptor      // wsAcceptor 对象
	stopGroup           sync.WaitGroup // 协程停止组
}

// 创建1个 mulAcceptor 对象
func NewMulAcceptor(addr *TLaddr, mgr IMulConnManager) (IAcceptor, error) {
	var err error

	// 参数效验
	ok := (addr.TcpAddr == "" || addr.WsAddr == "")
	if ok {
		err = errors.New("创建 MulAcceptor 失败。参数 TcpAddr WsAddr 为空")

		return nil, err
	}

	if nil == mgr {
		err = errors.New("创建 MulAcceptor 失败。参数 IMulConnManager=nil")

		return nil, err
	}

	// 创建对象
	sm := state.NewStateManager()
	tcpaptor := NewTcpAcceptor(addr, mgr)
	wsaptor := NewWsAcceptor(addr, mgr)

	// 创建 MulAcceptor
	mulaptor := &MulAcceptor{
		StateManager: sm,
		name:         C_ACCEPTOR_NAME_MUL,
		laddr:        addr,
		tcpAcceptor:  tcpaptor,
		wsAcceptor:   wsaptor,
	}

	// 设置为初始化状态
	mulaptor.SetState(state.C_STATE_INIT)

	return mulaptor
}

// 启动 mulAcceptor [IAcceptor 接口]
func (this *MulAcceptor) Run() error {
	var err error
	// 状态效验
	if this.GetState() != state.C_STATE_INIT {
		return err
	}

	// 改变状态: 正在启动中
	this.SetState(state.C_STATE_RUNING)

	// 添加启动线程数量
	this.stopGroup.Add(2)

	// 启动 tcp
	this.tcpAcceptor.Run()

	// 启动 websocket
	this.wsAcceptor.Run()

	// 阻塞
	//this.RunWait()

	// 启动完成

	return err
}

// 停止 mulAcceptor [IAcceptor 接口]
func (this *MulAcceptor) Stop() error {
	// 停止 tcp
	this.tcpAcceptor.Stop()

	// 停止 websocket
	this.wsAcceptor.Stop()

	return nil
}
