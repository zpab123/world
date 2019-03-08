// /////////////////////////////////////////////////////////////////////////////
// 面向服务器连接的 session 组件

package session

import (
	"github.com/pkg/errors"            // 异常库
	"github.com/zpab123/syncutil"      // 原子变量
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/state"   // 状态管理
	"github.com/zpab123/world/wderr"   // 异常库
	"github.com/zpab123/zaplog"        // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// ServerSession 对象

// 面向服务器连接的 session 对象
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
	if !this.stateMgr.SwapState(state.C_INIT, state.C_RUNING) {
		if !this.stateMgr.SwapState(state.C_STOPED, state.C_RUNING) {
			err = errors.Errorf("ServerSession 启动失败，状态错误。当前状态=%d，正确状态=%d或%d", this.stateMgr.GetState(), state.C_INIT, state.C_STOPED)

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

	// 主循环

	// 改变状态： 工作中
	if !this.stateMgr.SwapState(state.C_RUNING, state.C_WORKING) {
		err = errors.Errorf("ServerSession 启动失败，状态错误。当前状态=%d，正确状态=%d", this.stateMgr.GetState(), state.C_RUNING)

		return
	}

	return
}

// 关闭 session [ISession 接口]
func (this *ServerSession) Stop() (err error) {
	// 状态改变为关闭中
	if !this.stateMgr.SwapState(state.C_WORKING, state.C_CLOSEING) {
		err = errors.Errorf("ServerSession %s 关闭失败，状态错误。当前状态=%d, 正确状态=%d", this, this.stateMgr.GetState(), state.C_WORKING)

		return
	}

	// 关闭连接
	err = this.worldConn.Close()
	if nil != err {
		err = errors.Errorf("ServerSession %s 关闭失败。错误=%s", this, err)

		return
	}

	// 状态改变为关闭完成
	if !this.stateMgr.SwapState(state.C_CLOSEING, state.C_CLOSED) {
		err = errors.Errorf("ServerSession %s 关闭失败，状态错误。当前状态=%d, 正确状态=%d", this, this.stateMgr.GetState(), state.C_CLOSEING)

		return
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

// 打印信息
func (this *ServerSession) String() string {
	return this.worldConn.String()
}

// 发送心跳消息
func (this *ServerSession) SendHeartbeat() {
	this.worldConn.SendHeartbeat()
}

// 发送通用消息
func (this *ServerSession) SendData(data []byte) {
	this.worldConn.SendData(data)
}

// 接收线程
func (this *ServerSession) recvLoop() {
	defer func() {
		this.Stop()

		if err := recover(); nil != err && !wderr.IsConnectionError(err.(error)) {
			zaplog.TraceError("ServerSession %s 接收数据出现错误：%s", this, err.(error))
		} else {
			zaplog.Debugf("ServerSession %s 断开连接", this)
		}
	}()

	for this.stateMgr.GetState() == state.C_WORKING {
		// 接收消息
		pkt, err := this.worldConn.RecvPacket()

		// 错误处理
		if nil != err && !wderr.IsTimeoutError(err) {
			if wderr.IsConnectionError(err) {
				break
			} else {
				panic(err)
			}
		}

		// 消息处理
		if nil == pkt {
			continue
		}

		if this.msgHandler != nil {
			this.msgHandler.OnServerMessage(this, pkt) // 这里还需要增加异常处理
		}

	}
}

// 发送线程
func (this *ServerSession) sendLoop() {
	var err error

	for this.stateMgr.GetState() == state.C_WORKING {
		err = this.worldConn.Flush() // 刷新缓冲区

		if nil != err {
			break
		}
	}
}

// 主循环
func (this *ServerSession) mainLoop() {

}
