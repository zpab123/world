// /////////////////////////////////////////////////////////////////////////////
// 面向前端的 session 组件

package session

import (
	"github.com/zpab123/syncutil"      // 原子变量
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/state"   // 状态管理
	"github.com/zpab123/zplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// FrontendSession 对象

// 面向前端的 session 对象
type FrontendSession struct {
	stateMgr    *state.StateManager      // 对象继承： 状态管理
	worldConn   *network.WorldConnection // world 引擎连接对象
	sesssionMgr ISessionManage           // sessiong 管理对象
	sessionId   syncutil.AtomicInt64     // session ID
	msgHandler  IMsgHandler              // 消息处理器
}

// 创建1个新的 FrontendSession 对象
func NewFrontendSession(socket network.ISocket, mgr ISessionManage, opt *TSessionOpts) ISession {
	// 创建 StateManager
	st := state.NewStateManager()

	// 创建 WorldConnection
	if nil == opt {
		opt = NewTSessionOpts(nil)
	}
	wc := network.NewWorldConnection(socket, opt.WorldConnOpts)

	// 创建对象
	cs := &FrontendSession{
		stateMgr:    st,
		worldConn:   wc,
		sesssionMgr: mgr,
		msgHandler:  opt.MsgHandler,
	}

	// 修改为初始化状态
	cs.stateMgr.SetState(state.C_STATE_INIT)

	return cs
}

// 启动 session [ISession 接口]
func (this *FrontendSession) Run() {
	// 改变状态： 启动中
	if !this.stateMgr.SwapState(state.C_STATE_INIT, state.C_STATE_RUNING) {
		zplog.Errorf("FrontendSession 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_INIT, this.stateMgr.GetState())

		return
	}

	// 变量重置？ 状态? 发送队列？

	// 将 session 添加到管理器, 在线程处理前添加到管理器(分配id), 避免ID还未分配,就开始使用id的竞态问题
	this.sesssionMgr.OnNewSession(this)

	// 开启接收线程
	go this.recvLoop()

	// 开启发送线程
	go this.sendLoop()

	// 改变状态： 工作中
	if !this.stateMgr.SwapState(state.C_STATE_RUNING, state.C_STATE_WORKING) {
		zplog.Errorf("FrontendSession 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_RUNING, this.stateMgr.GetState())
	}
}

// 关闭 session [ISession 接口]
func (this *FrontendSession) Close() {
	// 状态改变为关闭中
	if !this.stateMgr.SwapState(state.C_STATE_WORKING, state.C_STATE_CLOSEING) {
		zplog.Errorf("FrontendSession 关闭失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_WORKING, this.stateMgr.GetState())

		return
	}

	// 关闭连接
	this.worldConn.Close()

	// 状态改变为关闭完成
	if !this.stateMgr.SwapState(state.C_STATE_CLOSEING, state.C_STATE_CLOSED) {
		zplog.Errorf("FrontendSession 关闭失败，状态错误。正确状态=%d，当前状态=%d", state.C_STATE_CLOSEING, this.stateMgr.GetState())
	}

	// 通知 session 管理
	this.sesssionMgr.OnSessionClose(this)
}

// 获取 session ID [ISession 接口]
func (this *FrontendSession) GetId() int64 {
	return this.sessionId.Load()
}

// 设置 session ID [ISession 接口]
func (this *FrontendSession) SetId(v int64) {
	this.sessionId.Store(v)
}

// 接收线程
func (this *FrontendSession) recvLoop() {
	for {
		// 心跳检查
		this.worldConn.CheckClientHeartbeat()

		// 接收消息
		var pkt *network.Packet
		pkt, _ = this.worldConn.RecvPacket()
		if nil == pkt {
			continue
		}

		// 创建消息：后续使用对象池？
		msg := &Message{
			session: this,
			packet:  pkt,
		}

		// 消息处理
		if this.msgHandler != nil {
			this.msgHandler.OnNewMessage(msg)
		}
	}
}

// 发送线程
func (this *FrontendSession) sendLoop() {
	for {
		// 心跳检查
		this.worldConn.CheckServerHeartbeat()

		// 刷新缓冲区
		this.worldConn.Flush()
	}
}