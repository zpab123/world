// /////////////////////////////////////////////////////////////////////////////
// 能够读写 packet 数据的 socket

package socket

import (
	"encoding/binary"
	"net"
	"sync"

	"github.com/zpab123/world/model"          // 全局模型
	"github.com/zpab123/world/network"        // 网络库
	"github.com/zpab123/world/network/packet" // 消息包
	"github.com/zpab123/world/queue"          // 消息队列
	"github.com/zpab123/world/utils"          // 工具库
	"github.com/zpab123/zplog"                // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 初始化

// 常量
const (
	_HEAD_LEN        = model.C_PACKET_HEAD_LEN                          // 消息头长度
	_MAX_BODY_LENGTH = model.C_PACKET_MAX_LEN - model.C_PACKET_HEAD_LEN // body 数据最大长度 （ pcket总长度 - 消息头）
)

// 变量
var (
	NETWORK_ENDIAN = binary.LittleEndian // 小端读取对象
	errRecvAgain   = _ErrRecvAgain{}     // 重新接收错误
)

// /////////////////////////////////////////////////////////////////////////////
// PacketSocket 对象

// PacketSocket
type PacketSocket struct {
	socket    model.ISocket    // 接口继承： 符合 ISocket 的对象
	connector model.IConnector // connector 组件
	goMutex   sync.Mutex       // 线程互斥锁
	sendQueue []*packet.Packet // 发送队列
	recvedLen uint32           // 从 socket 的 readbuffer 中已经读取的数据大小：字节（用于消息读取记录）
	headBuff  [_HEAD_LEN]byte  // 存放消息头二进制数据
	pType     uint16           // 本次 packet 类型
	bodylen   uint32           // 本次 pcket body 总大小
	newPacket *packet.Packet   // 用于存储 本次即将接收的 Packet 对象
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

// 接收下1个 packet 数据
//
// 返回, rerutn=nil=没收到完整的 packet 数据; rerutn=packet=完整的 packet 数据包
func (this *PacketSocket) RecvPacket() (*packet.Packet, error) {
	// 还未收到消息头
	if this.recvedLen < _HEAD_LEN {
		n, err := this.socket.Read(this.headBuff[this.recvedLen:]) // 读取数据
		this.recvedLen += uint32(n)

		// 还是没收到完整消息头
		if this.recvedLen < _HEAD_LEN {
			if nil == err {
				err = errRecvAgain
			}
		}

		// 收到消息头: 保存消息类型
		this.pType = NETWORK_ENDIAN.Uint16(this.headBuff[0:2])

		// 收到消息头: 保存本次 packet 消息 body 总大小
		this.bodylen = NETWORK_ENDIAN.Uint16(this.headBuff[2:])

		// 解密

		// 长度效验
		if this.bodylen > _MAX_BODY_LENGTH {
			err := errors.Errorf("packet 消息包数据 body 长度大于最大长度。长度=：%v", pc.bodylen)
			this.resetRecvStates()
			this.Close()

			return nil, err
		}

		// 创建新的 packet 对象
		this.recvedLen = 0 // 重置，准备记录 body
		this.newPacket = packet.NewPacket()
		this.newPacket.AllocBuffer(this.bodylen)
	}

	// 接收 pcket 数据的 body 部分
	n, err := this.socket.Read(this.newPacket.GetBytes()[_HEAD_LEN+this.recvedLen : _HEAD_LEN+this.bodylen])
	this.recvedLen += uint32(n)

	// 接收完成， packet 数据包完整
	if this.recvedLen == this.bodylen {
		// 解密

		// 准备接收下1个
		packet := this.newPacket
		this.resetRecvStates()

		return packet, nil
	}

	// body 收完
	if nil == err {
		err = errRecvAgain
	}

	return nil, err
}

// 发送1个 *Packe 数据
func (this *PacketSocket) SendPacket(pkt *packet.Packet) error {
	// 添加到消息队列
	this.goMutex.Lock()
	this.sendQueue = append(this.sendQueue, pkt)
	this.goMutex.Unlock()

	return nil
}

// 将消息队列中的数据写入 writebuff
func (this *PacketSocket) Flush(reason *string) (err error) {
	// 复制数据
	this.goMutex.Lock()
	if len(this.sendQueue) == 0 {
		this.goMutex.Unlock()

		return
	}

	// 交换数据, 并把原来的数据置空
	packets := make([]*packet.Packet, 0, len(this.sendQueue))
	packets, this.sendQueue = this.sendQueue, packets // 交换值
	this.goMutex.Unlock()

	// 刷新数据
	if 1 == len(packets) {
		pkt := packets[0]

		// 将 data 写入 conn
		err = utils.WriteAll(this.socket, pkt.Data())
		pkt.Release()

		if nil == err {
			err = this.socket.Flush()
		}

		return
	}

	for _, pkt := range packets {
		err = utils.WriteAll(this.socket, pkt.Data())
		pkt.Release()
	}

	if nil == err {
		err = this.socket.Flush()
	}

	return
}

// 关闭 socket
func (this *PacketSocket) Close() error {
	return this.socket.Close()
}

// 获取客户端 ip 地址
func (this *PacketSocket) RemoteAddr() net.Addr {
	return this.socket.RemoteAddr()
}

// 获取本地 ip 地址
func (this *PacketSocket) LocalAddr() net.Addr {
	return this.socket.LocalAddr()
}

// fmt 字符串输出接口
func (pc *PacketSocket) String() string {
	return fmt.Sprintf("[%s >>> %s]", this.LocalAddr(), this.RemoteAddr())
}

// 重置数据接收状态
func (this *PacketSocket) resetRecvStates() {
	this.recvedLen = 0
	this.bodylen = 0
	this.pType = model.C_MSG_TYPE_INVALID
	this.newPacket = nil
}

// /////////////////////////////////////////////////////////////////////////////
// _ErrRecvAgain 对象

type _ErrRecvAgain struct{}

func (err _ErrRecvAgain) Error() string {
	return "继续接收 packet"
}

func (err _ErrRecvAgain) Temporary() bool {
	return true
}

func (err _ErrRecvAgain) Timeout() bool {
	return true
}
