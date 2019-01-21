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
	opts         *model.TSessionOpts  // session 配置参数
	packetSocket *socket.PacketSocket // 对象继承： 继承至 PacketSocket 对象
	sesssionMgr  model.ISessionManage // sessiong 管理对象
	// session_id
	// 用户id
	pktHandler model.ICilentPktHandler // 客户端 packet 消息处理器
	state      syncutil.AtomicUint32   // 状态变量
}

func NewClientSession(opt *model.TSessionOpts) *ClientSession {
	// 创建 pktSocket
	socket := &network.Socket{
		conn: netconn,
	}
	bufSocket := network.NewBufferSocket(socket)
	pktSocket := network.NewPacketSocket(bufSocket)

	// 创建对象
	cs := &ClientSession{
		opts:         opt,
		packetSocket: pktSocket,
		sesssionMgr:  mgr,
		pktHandler:   handler,
	}

	// 修改为初始化状态
	cs.state.Store(model.C_SES_STATE_INITED)

	return cs
}

// run
func (this *ClientSession) Run() {
	// 结束线程

	// 开启接收线程
	go this.recvLoop()

	// 开启发送线程
	go this.sendLoop()
}

// 获取 session 状态 [IState 接口]
func (this *ClientSession) GetState() uint32 {
	return this.state.Load()
}

// 设置 session 状态 [IState 接口]
func (this *ClientSession) SetState(v uint32) {
	this.state.Store(v)
}

// 获取配置参数
func (this *ClientSession) GetOpts() *model.TSessionOpts {
	return this.opts
}

// 接收线程
func (this *ClientSession) recvLoop() {
	for {
		// 接收消息
		var pkt *packet.Packet
		pkt = this.packetSocket.RecvPacket()
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
		this.packetSocket.Flush()
	}
}
