// /////////////////////////////////////////////////////////////////////////////
// 面向前端的 session 组件

package session

import (
	"github.com/zpab123/syncutil"      // 原子变量
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/state"   // 状态管理
	"github.com/zpab123/zaplog"        // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// ServerSession 对象

// 面向后端的 session 对象
type ServerSession struct {
	option      *TServerSessionOpt       // 配置参数
	stateMgr    *state.StateManager      // 状态管理
	worldConn   *network.WorldConnection // world 引擎连接对象
	sesssionMgr ISessionManage           // sessiong 管理对象
	sessionId   syncutil.AtomicInt64     // session ID
	msgHandler  IServerMsgHandler        // 消息处理器
}

// 创建1个新的 ServerSession 对象
func NewServerSession(socket network.ISocket, mgr ISessionManage, opt *TServerSessionOpt) ISession {
	// 创建 StateManager
	st := state.NewStateManager()

	// 创建 WorldConnection
	if nil == opt {
		opt = NewTServerSessionOpt(nil)
	}
	wc := network.NewWorldConnection(socket, opt.WorldConnOpt)

	// 创建对象
	ss := &ServerSession{
		option:      opt,
		stateMgr:    st,
		worldConn:   wc,
		sesssionMgr: mgr,
		msgHandler:  opt.ServerMsgHandler,
	}

	// 修改为初始化状态
	ss.stateMgr.SetState(state.C_INIT)

	return ss
}

// 启动 session [ISession 接口]
func (this *ServerSession) Run() (err error) {
	// 状态效验
	s := this.stateMgr.GetState()
	if s != state.C_INIT && s != state.C_CLOSED {
		zaplog.Errorf("ServerSession 启动失败，状态错误。正确状态=%d或%d，当前状态=%d", state.C_INIT, state.C_CLOSED, s)

		return
	}

	// 改变状态： 启动中
	this.stateMgr.SetState(state.C_RUNING)

	// 变量重置？ 状态? 发送队列？

	// 将 session 添加到管理器, 在线程处理前添加到管理器(分配id), 避免ID还未分配,就开始使用id的竞态问题
	this.sesssionMgr.OnNewSession(this)

	// 开启接收线程
	go this.recvLoop()

	// 开启发送线程
	go this.sendLoop()

	// 改变状态： 工作中
	if !this.stateMgr.SwapState(state.C_RUNING, state.C_WORKING) {
		zaplog.Errorf("ServerSession 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_RUNING, this.stateMgr.GetState())
	}

	return
}

// 关闭 session [ISession 接口]
func (this *ServerSession) Stop() (err error) {
	// 状态改变为关闭中
	if !this.stateMgr.SwapState(state.C_WORKING, state.C_CLOSEING) {
		zaplog.Errorf("ServerSession 关闭失败，状态错误。正确状态=%d，当前状态=%d", state.C_WORKING, this.stateMgr.GetState())

		return
	}

	// 关闭连接
	this.worldConn.Close()

	// 状态改变为关闭完成
	if !this.stateMgr.SwapState(state.C_CLOSEING, state.C_CLOSED) {
		zaplog.Errorf("ServerSession 关闭失败，状态错误。正确状态=%d，当前状态=%d", state.C_CLOSEING, this.stateMgr.GetState())
	}

	// 通知 session 管理
	this.sesssionMgr.OnSessionClose(this)

	return
}

// 获取 session ID [ISession 接口]
func (this *ServerSession) GetId() int64 {
	return this.sessionId.Load()
}

// 设置 session ID [ISession 接口]
func (this *ServerSession) SetId(v int64) {
	this.sessionId.Store(v)
}

// 接收线程
func (this *ServerSession) recvLoop() {
	for {
		// 接收消息
		pkt, _ := this.worldConn.RecvPacket()
		if nil == pkt {
			continue
		}

		// 消息处理
		if this.msgHandler != nil {
			this.msgHandler.OnServerMessage(this, pkt)
		}

	}
}

// 发送线程
func (this *ServerSession) sendLoop() {
	var err error

	for {
		// 刷新缓冲区
		err = this.worldConn.Flush()
		if nil != err {
			break
		}
	}
}
