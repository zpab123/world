// /////////////////////////////////////////////////////////////////////////////
// 面向客户端的 session 组件

package session

import (
	"net"
	"sync"

	"github.com/zpab123/syncutil"      // 原子变量
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/state"   // 状态管理
	"github.com/zpab123/zplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// ClientSession 对象

// 面向客户端的 session 对象
type ClientSession struct {
	*state.StateManager                          // 对象继承： 状态管理
	opts                *model.TSessionOpts      // session 配置参数
	worldConn           *network.WorldConnection // world 引擎连接对象
	sesssionMgr         model.ISessionManage     // sessiong 管理对象
	stopGroup           sync.WaitGroup           // 结束线程组
	state               syncutil.AtomicUint32    // 状态变量
	sessionId           syncutil.AtomicInt64     // session ID
	// 用户id
}

// 创建1个新的 ClientSession 对象
func NewClientSession(socket model.ISocket, mgr model.ISessionManage, opt *model.TSessionOpts) *ClientSession {
	// 创建 StateManager
	st := state.NewStateManager()

	// 创建 WorldConnection
	if nil == opt {
		opt = model.NewTSessionOpts()
	}
	wc := network.NewWorldConnection(socket, opt.WorldConnOpts)

	// 创建对象
	cs := &ClientSession{
		StateManager: st,
		opts:         opt,
		worldConn:    wc,
		sesssionMgr:  mgr,
	}

	// 修改为初始化状态
	cs.SetState(model.C_STATE_INIT)

	return cs
}

// 启动 session
func (this *ClientSession) Run() {
	// 非 INIT 状态
	if this.state.Load() != model.C_STATE_INIT {
		zplog.Errorf("ClientSession 启动失败。状态不在初始化状态")
		return
	}

	// 变量重置？ 状态? 发送队列？

	// 需要接收和发送线程都启动/关闭才算完成
	this.AddRunGo(2)
	this.AddStopGo(2)

	// 将 session 添加到管理器, 在线程处理前添加到管理器(分配id), 避免ID还未分配,就开始使用id的竞态问题
	this.sesssionMgr.OnNewSession(this)

	// 开启接收线程
	go this.recvLoop()

	// 开启发送线程
	go this.sendLoop()

	// 阻塞
	this.RunWait()

	// 启动完成
	this.SetState(model.C_STATE_WORKING)
}

// 关闭 session [ISession 接口]
func (this *ClientSession) Close() {
	// 非运行状态
	if this.state.Load() != model.C_STATE_WORKING {
		return
	}

	// 状态改变为关闭中
	this.state.Store(model.C_SES_STATE_CLOSING)

	// 关闭连接
	this.worldConn.Close()

	// 关闭阻塞
	this.StopWait()

	// 关闭完成
	this.SetState(model.C_STATE_CLOSED)

	// 通知 session 管理
	this.sesssionMgr.OnSessionClose(this)
}

// 获取 session ID [ISession 接口]
func (this *ClientSession) GetId() int64 {
	return this.sessionId.Load()
}

// 设置 session ID [ISession 接口]
func (this *ClientSession) SetId(v int64) int64 {
	return this.sessionId.Store(v)
}

// 接收线程
func (this *ClientSession) recvLoop() {
	// 启动线程完成1个
	this.RunDone()

	for {
		// 退出检查
		if exit := this.isCloseIng(); exit {
			break
		}

		// 心跳检查
		this.worldConn.CheckClientHeartbeat()

		// 接收消息
		var pkt *packet.Packet
		pkt = this.worldConn.RecvPacket()
		if nil == pkt {
			continue
		}

		// 处理消息
		//handlePacket(pkt)
	}

	// 接收线程结束
	this.stopGroup.Done()
}

// 发送线程
func (this *ClientSession) sendLoop() {
	// 启动线程完成1个
	this.RunDone()

	for {
		// 退出检查
		if exit := this.isCloseIng(); exit {
			break
		}

		// 心跳检查
		this.worldConn.CheckServerHeartbeat()

		// 刷新缓冲区
		this.worldConn.Flush()
	}

	// 发送线程结束
	this.stopGroup.Done()
}

// 是否正在结束
func (this *ClientSession) isCloseIng() bool {
	return this.state.Load() == model.C_SES_STATE_CLOSING
}
