// /////////////////////////////////////////////////////////////////////////////
// 能够读写 packet 数据的 socket

package socket

import (
	"github.com/zpab123/world/model"   // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/queue"   // 消息队列
	"github.com/zpab123/world/utils"   // 工具库
	"github.com/zpab123/zplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// PacketSocket 对象

// PacketSocket
type PacketSocket struct {
	socket                model.ISocket    // 接口继承： 符合 ISocket 的对象
	connector             model.IConnector // connector 组件
	network.PacketManager                  // 对象继承： packet 消息管理对象
}

// 创建1个新的 PacketSocket 对象
func NewPacketSocket(st model.ISocket, cntor model.IConnector) *PacketSocket {

	// 设置 packet 处理器

	// 创建对象
	pktSocket := &PacketSocket{
		socket:    st,
		connector: cntor,
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
	go this.recvLoop()

	// 启动并发发送 goroutine
	go this.sendLoop()
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
	var capturePanic bool = this.connector.GetRecoverIoPanic()

	// 数据处理
	for nil != this.socket {
		var pkt interface{}
		var err error

		// 接收数据
		if capturePanic {
			pkt, err = this.protectedRecv()
		} else {
			pkt, err = this.RecvPacket()
		}

		// 接收出错
		if nil != err {
			// EOFO 错误
			if !utils.IsEOFOrNetReadError(err) {
				zplog.Errorf("PacketSocket 关闭。 socketId=, err=%s", err)
			}

			// 准备关闭发送线程
			this.sendQueue.Add(nil)

			// 通知 connector 组件

			break
		}

		// 自身数据处理 -- 握手、心跳等

		// 通知 connector 的数据
	}

	// 结束发送线程
}

// 发送循环
func (this *PacketSocket) sendLoop() {
	// 发送切片
	var writeList []interface{}

	// 复制数据
	for {
		writeList = writeList[0:0]
		exit := this.sendQueue.Pick(&writeList)

		// 遍历要发送的数据
		for _, pkt := range writeList {
			this.SendPacket(pkt)
		}

		if exit {
			break
		}
	}

	// 完整关闭
	this.socket.Close()

	// 通知其他线程

}

// 保护性读取 packet 数据 （带异常捕获）
func (this *PacketSocket) protectedRecv() (pkt interface{}, err error) {
	// 异常捕获
	defer func() {
		if err := recover(); err != nil {
			zplog.Panicf("PacketSocket 接收 packet 异常。错误=%s", err)
			this.socket.Close()
		}
	}()

	pkt, err = this.RecvPacket()

	return
}
