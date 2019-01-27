// /////////////////////////////////////////////////////////////////////////////
// 面向客户端的 session 组件

package session

import (
	"net"

	"github.com/zpab123/syncutil"      // 原子变量
	"github.com/zpab123/world/model"   // 全局模型
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/zplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// /////////////////////////////////////////////////////////////////////////////
// ClientSession 对象

// 面向客户端的 session 对象
type ClientSession struct {
	opts        *model.TSessionOpts      // session 配置参数
	worldConn   *network.WorldConnection // world 引擎连接对象
	sesssionMgr model.ISessionManage     // sessiong 管理对象
	// session_id
	// 用户id
	pktHandler model.ICilentPktHandler // 客户端 packet 消息处理器
	state      syncutil.AtomicUint32   // 状态变量
}

func NewClientSession(socket model.ISocket, opt *model.TSessionOpts) *ClientSession {
	// 创建 WorldConnection
	if nil == opt {
		opt = model.NewTSessionOpts()
	}
	wc := network.NewWorldConnection(socket, opt.WorldConnOpts)

	// 创建对象
	cs := &ClientSession{
		opts:      opt,
		worldConn: wc,
	}

	// 修改为初始化状态
	cs.state.Store(model.C_SES_STATE_INITED)

	return cs
}

// 启动 session
func (this *ClientSession) Run() {
	// 结束线程

	// 开启接收线程
	go this.recvLoop()

	// 开启发送线程
	go this.sendLoop()
}

// 接收线程
func (this *ClientSession) recvLoop() {
	for {
		// 心跳检查
		this.worldConn.CheckClientHeartbeat()

		// 接收消息
		var pkt *packet.Packet
		pkt = this.worldConn.RecvPacket()
		if nil == pkt {
			continue
		}

		// 处理消息
		handlePacket(pkt)
	}
}

// 发送线程
func (this *ClientSession) sendLoop() {
	for {
		// 心跳检查
		this.worldConn.CheckServerHeartbeat()

		// 刷新缓冲区
		this.worldConn.Flush()
	}
}
