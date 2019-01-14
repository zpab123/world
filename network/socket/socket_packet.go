// /////////////////////////////////////////////////////////////////////////////
// 能够读写 packet 数据的 socket

package socket

import (
	"github.com/zpab123/world/model"   // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// PacketSocket 对象

// PacketSocket
type PacketSocket struct {
	socket                model.ISocket    // 接口继承： 符合 ISocket 的对象
	sendQueue             *network.Pipe    // 发送队列
	connector             model.IConnector // connector 组件
	network.PacketManager                  // 对象继承： packet 消息管理对象
}

// 创建1个新的 PacketSocket 对象
func NewPacketSocket(st model.ISocket, cntor model.IConnector) *PacketSocket {
	// 创建发送队列
	que := network.NewPipe()

	// 设置 packet 处理器

	// 创建对象
	pktSocket := &PacketSocket{
		socket:    st,
		connector: cntor,
		sendQueue: que,
	}

	return pktSocket
}

// 启动 PacketSocket
func (this *PacketSocket) Run() {
	// 状态处理

	// 清空发送队列
	this.sendQueue.Reset()

	// 设置线程结束数量

	// 将会话添加到管理器
	//this.connector.

	// 结束监听 goroutine

	// 启动并发接收 goroutine

	// 启动并发发送 goroutine
}

// 关闭 PacketSocket
func (this *PacketSocket) Close() {
	// 状态交换

	// 关闭连接

	// 超时续时
}

// 接收循环
func (this *PacketSocket) recvLoop() {
	// 是否进行 io 异常捕获
	var capturePanic bool
	//if i, ok := this.connector.(network.)

	for nil != this.socket {
		// 接收数据
		pkt := this.ReadPacket()

		// 自身数据处理 -- 握手、心跳等

		// 通知 connector 的数据
	}
}

// 发送循环
func (this *PacketSocket) sendLoop() {

}

// 接收下1个 packet 数据
//
// 返回, nil=没收到完整的 packet 数据; packet=完整的 packet 数据包
