// /////////////////////////////////////////////////////////////////////////////
// 对 PacketSocket 的封装，定义一些 world 内部常用的消息

package proto

import (
	"encoding/binary"
	"net"
	"sync"
	"time"

	"github.com/zpab123/world/model"          // 全局模型
	"github.com/zpab123/world/network"        // 网络库
	"github.com/zpab123/world/network/packet" // 消息包
	"github.com/zpab123/world/queue"          // 消息队列
	"github.com/zpab123/world/utils"          // 工具库
	"github.com/zpab123/zplog"                // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// WorldConnection 对象

// world 框架内部需要用到的一些常用网络消息
type WorldConnection struct {
	opts         *model.TWorldSocketOpts // 配置参数
	packetSocket *network.PacketSocket   // 接口继承： 符合 IPacketSocket 的对象
}

// 新建1个 WorldConnection 对象
func NewWorldConnection(socket model.ISocket, opt *model.TWorldSocketOpts) *WorldConnection {
	// 创建 packetSocket
	bufSocket := network.NewBufferSocket(socket)
	pktSocket := network.NewPacketSocket(bufSocket)

	// 创建对象
	wc := &WorldConnection{
		packetSocket: pktSocket,
	}

	return wc
}

// 接收1个 msg 消息
func (this *WorldConnection) RecvMsg() {

}

// 发送1个 msg 消息
func (this *WorldConnection) SendMsg() {

}

// 回应握手消息
func (this *WorldConnection) HandshakeResponse() {

}

// 心跳消息

// 游戏常用内部消息
