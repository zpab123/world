// /////////////////////////////////////////////////////////////////////////////
// 同时支持 tcp websocket

package acceptor

import (
	"github.com/zpab123/world/model" // 全局模型
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// mulAcceptor 对象

// 同时支持 tcp websocket 的对象
type MulAcceptor struct {
	State                    // 对象继承：运行状态操作
	name        string       // 连接器名字
	laddr       model.TLaddr // 地址集合
	tcpAcceptor *TcpAcceptor // tcpAcceptor 对象
	wsAcceptor  *WsAcceptor  // wsAcceptor 对象
}

// 创建1个 mulAcceptor 对象
func newMulAcceptor(addr *model.TLaddr, mgr model.IMulSocketManager) model.IAcceptor {
	// 创建 tcpAcceptor
	tcpaptor := NewTcpAcceptor(addr, mgr)

	// 创建 wsAcceptor
	wsaptor := NewWsAcceptor(addr, mgr)

	// 创建对象
	mulaptor := &mulAcceptor{
		name:        model.C_ACCEPTOR_NAME_MUL,
		laddr:       addr,
		tcpAcceptor: tcpaptor,
		wsAcceptor:  wsaptor,
	}

	return mulaptor
}

// 启动 mulAcceptor [IAcceptor 接口]
func (this *MulAcceptor) Run() {
	// 启动 tcp
	this.tcpAcceptor.Run()

	// 启动 websocket
	this.wsAcceptor.Run()
}

// 停止 mulAcceptor [IAcceptor 接口]
func (this *MulAcceptor) Stop() {
	// 停止 tcp
	this.tcpAcceptor.Stop()

	// 停止 websocket
	this.wsAcceptor.Stop()
}
