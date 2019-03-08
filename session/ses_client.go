// /////////////////////////////////////////////////////////////////////////////
// 面向客户端的 session 组件

package session

import (
	"github.com/pkg/errors"            // 异常库
	"github.com/zpab123/syncutil"      // 原子变量
	"github.com/zpab123/world/config"  // 配置文件
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/state"   // 状态管理
	"github.com/zpab123/zaplog"        // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// ClientSession 对象

// 面向客户端的 session 对象
type ClientSession struct {
	option      *TClientSessionOpt       // 配置参数
	stateMgr    *state.StateManager      // 状态管理
	worldConn   *network.WorldConnection // server 端连接对象
	sesssionMgr ISessionManage           // sessiong 管理对象
	sessionId   syncutil.AtomicInt64     // session ID
	msgHandler  IClientMsgHandler        // 消息处理器
}

// 创建1个新的 ClientSession 对象
func NewClientSession(socket network.ISocket, mgr ISessionManage, opt *TClientSessionOpt) ISession {
	// 创建 StateManager
	st := state.NewStateManager()

	// 创建 WorldConnection
	if nil == opt {
		opt = NewTClientSessionOpt(nil)
		opt.WorldConnOpt.ShakeKey = config.GetWorldConfig().ShakeKey
	}
	sc := network.NewWorldConnection(socket, opt.WorldConnOpt)

	// 创建对象
	cs := &ClientSession{
		option:      opt,
		stateMgr:    st,
		worldConn:   sc,
		sesssionMgr: mgr,
		msgHandler:  opt.MsgHandler,
	}

	// 修改为初始化状态
	cs.stateMgr.SetState(C_SES_STATE_INIT)

	return cs
}

// 启动 session [ISession 接口]
func (this *ClientSession) Run() (err error) {
	// 改变状态： 启动中
	if !this.stateMgr.SwapState(C_SES_STATE_INIT, C_SES_STATE_RUNING) {
		if !this.stateMgr.SwapState(C_SES_STATE_STOPED, C_SES_STATE_RUNING) {
			err = errors.Errorf("ClientSession 启动失败，状态错误。当前状态=%d，正确状态=%d或%d", this.stateMgr.GetState(), C_SES_STATE_INIT, C_SES_STATE_STOPED)

			return
		}
	}

	// 变量重置？ 状态? 发送队列？

	// 将 session 添加到管理器, 在线程处理前添加到管理器(分配id), 避免ID还未分配,就开始使用id的竞态问题
	this.sesssionMgr.OnNewSession(this)

	// 开启接收线程
	go this.recvLoop()

	// 开启发送线程
	go this.sendLoop()

	// 改变状态： 工作中
	if !this.stateMgr.SwapState(state.C_RUNING, state.C_WORKING) {
		zaplog.Errorf("ClientSession 启动失败，状态错误。正确状态=%d，当前状态=%d", state.C_RUNING, this.stateMgr.GetState())
	}

	return
}

// 关闭 session [ISession 接口]
func (this *ClientSession) Stop() (err error) {
	// 状态改变为关闭中
	if !this.stateMgr.SwapState(state.C_WORKING, state.C_CLOSEING) {
		err = errors.Errorf("ClientSession 关闭失败，状态错误。当前状态=%d，正确状态=%d", this.stateMgr.GetState(), C_SES_STATE_WORKING)

		return
	}

	// 关闭连接
	this.worldConn.Close()

	// 状态改变为关闭完成
	if !this.stateMgr.SwapState(state.C_CLOSEING, state.C_CLOSED) {
		zaplog.Errorf("ClientSession 关闭失败，状态错误。正确状态=%d，当前状态=%d", state.C_CLOSEING, this.stateMgr.GetState())
	}

	// 通知 session 管理
	this.sesssionMgr.OnSessionClose(this)

	return
}

// 获取 session ID [ISession 接口]
func (this *ClientSession) GetId() int64 {
	return this.sessionId.Load()
}

// 设置 session ID [ISession 接口]
func (this *ClientSession) SetId(v int64) {
	this.sessionId.Store(v)
}

// 接收线程
func (this *ClientSession) recvLoop() {
	for {
		// 接收消息
		var pkt *network.Packet
		pkt, _ = this.worldConn.RecvPacket()
		if nil == pkt {
			continue
		}

		// 消息处理
		if this.msgHandler != nil {
			this.msgHandler.OnClientMessage(this, pkt)
		}
	}
}

// 发送线程
func (this *ClientSession) sendLoop() {
	var err error

	for {
		// 刷新缓冲区
		err = this.worldConn.Flush()
		if nil != err {
			break
		}
	}
}
