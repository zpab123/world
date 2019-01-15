// /////////////////////////////////////////////////////////////////////////////
// tcp 消息收发器

package tcp

import (
	"encoding/binary"
	"io"
	"net"

	"github.com/zpab123/world/model"          // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/world/network/packet" // packet 管理
	"github.com/zpab123/world/queue"          // 消息队列
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
// TcpDataManager 对象

// tcp 消息收发管理
type TcpDataManager struct {
	handler   model.IPacketHandler          // 数据处理对象
	recvedLen uint32                        // 从 socket 的 readbuffer 中已经读取的数据大小：字节（用于消息读取记录）
	headBuff  [consts.PACKET_HEAD_SIZE]byte // 存放消息头二进制数据
	pType     uint16                        // 本次 packet 类型
	bodylen   uint32                        // 本次 pcket body 总大小
	newPacket *packet.Packet                // 用于存储 本次即将接收的 Packet 对象
	sendQueue *queue.Pipe                   // 发送队列
}

// 创建1个 TcpDataManager 对象
func NewTcpDataManager(handler model.IPacketHandler) *TcpDataManager {
	// 创建发送队列
	que := queue.NewPipe()

	// 创建 dm
	dm := &TcpDataManager{
		handler:   handler,
		sendQueue: que,
	}

	return dm
}

// 接收1个 packet [IDataManager 接口]
func (this *TcpDataManager) RecvPacket(socket model.IPacketSocket) (interface{}, error) {
	// 接收消息头
	if this.recvedLen < _HEAD_LEN {
		// 读取数据
		n, err := socket.GetSocket().Read(this.headBuff[this.recvedLen:])
		this.recvedLen += uint32(n)

		// 还是没收到完整消息头
		if this.recvedLen < _HEAD_LEN {
			if nil == err {
				err = errRecvAgain
			}

			return nil, err
		}

		// 保存消息类型
		this.pType = NETWORK_ENDIAN.Uint16(this.headBuff[0:1])

		// 保存本次 packet 消息 body 总大小
		this.bodylen = NETWORK_ENDIAN.Uint16(this.headBuff[2:])

		// 解密

		// 长度效验
		if this.bodylen > _MAX_BODY_LENGTH {
			err := errors.Errorf("packet 消息包数据 body 长度错误。 长度=：%v", this.bodylen)
			this.resetRecvStates() // 重置接收状态
			socket.Close()         // 关闭连接

			return nil, err
		}

		// 创建新的 packet 对象
		this.recvedLen = 0 // 重置
		this.newPacket = packet.NewPacket()
		this.newPacket.AllocBuffer(this.bodylen)
	}

	// 接收 pcket 数据的 body 部分
	bytes := this.newPacket.GetBytes()
	n, err := socket.GetSocket().Read(bytes[_HEAD_LEN+this.recvedLen : _HEAD_LEN+this.bodylen])
	this.recvedLen += uint32(n)

	// 接收完成， packet 数据包完整
	if this.recvedLen == this.bodylen {
		// 解密
		packet := this.newPacket
		this.resetRecvStates()

		return packet, nil
	}

	if nil == err {
		err == errRecvAgain
	}

	return nil, err
}

// 发送1个 packet [IDataManager 接口]
func (this *TcpDataManager) SendPacket(socket model.IPacketSocket, pkt interface{}) (err error) {
	// 类型转化
	writer, ok := socket.GetSocket().(io.Writer)
	if !ok || nil == writer {
		return nil // 转换错误，或者连接已经关闭时退出
	}

	// 有写超时时，设置超时
	cntor := socket.GetConnector()
	cntor.SetSocketWriteTimeout(writer.(net.Conn), func() {

	})

	return
}

// 重置数据接收状态
func (this *TcpDataManager) resetRecvStates() {
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
