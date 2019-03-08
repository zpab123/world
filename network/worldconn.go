// /////////////////////////////////////////////////////////////////////////////
// 对 PacketSocket 的封装，面向前端的连接对象

package network

import (
	"github.com/pkg/errors"             // 异常库
	"github.com/vmihailenco/msgpack"    // 二进制结构体数据转化
	"github.com/zpab123/world/protocol" // world 内部通信协议
	"github.com/zpab123/world/state"    // 状态管理
	"github.com/zpab123/zaplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// WorldConnection 对象

// world 框架内部需要用到的一些常用网络消息
type WorldConnection struct {
	stateMgr     *state.StateManager // 状态管理
	option       *TWorldConnOpt      // 配置参数
	packetSocket *PacketSocket       // 接口继承： 符合 IPacketSocket 的对象
}

// 新建1个 WorldConnection 对象
func NewWorldConnection(socket ISocket, opt *TWorldConnOpt) *WorldConnection {
	// 参数效验
	if nil == opt {
		opt = NewTWorldConnOpt()
	}

	// 创建状态管理
	st := state.NewStateManager()

	// 创建 packetSocket
	bufSocket := NewBufferSocket(socket, opt.BuffSocketOpts)
	pktSocket := NewPacketSocket(bufSocket)

	// 创建对象
	wc := &WorldConnection{
		stateMgr:     st,
		packetSocket: pktSocket,
		option:       opt,
	}

	// 设置为初始化状态
	wc.stateMgr.SetState(C_CONN_STATE_INIT)

	return wc
}

// 接收1个 Packet 消息
func (this *WorldConnection) RecvPacket() (*Packet, error) {
	// 接收 packet
	pkt, err := this.packetSocket.RecvPacket()
	if nil == pkt || nil != err {
		return nil, err
	}

	// 内部 packet
	if pkt.pktId < protocol.C_PKT_ID_WORLD {
		this.handlePacket(pkt)

		return nil, err
	}

	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_WORKING {
		this.Close()

		err = errors.Errorf("WorldConnection %s 收到数据，但是状态错误。当前状态=%d，正确状态=%s", this, this.stateMgr.GetState(), C_CONN_STATE_WORKING)

		return nil, err
	}

	return pkt, nil
}

// 关闭 WorldConnection
func (this *WorldConnection) Close() error {
	var err error
	s := this.stateMgr.GetState()

	if s == C_CONN_STATE_CLOSED {
		err = errors.New("WorldConnection 关闭失败：它已经处于关闭状态")

		return err
	}

	err = this.packetSocket.Close()

	this.stateMgr.SetState(C_CONN_STATE_CLOSED)

	return err
}

// 发送1个 packet 消息
func (this *WorldConnection) SendPacket(pkt *Packet) error {
	var err error

	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_WORKING {
		err = errors.Errorf("WorldConnection %s 发送 Packet 数据失败：状态不在 working 中", this)

		return err
	}

	return this.packetSocket.SendPacket(pkt)
}

// 发送1个 packet 消息，然后将 packet 放回对象池
func (this *WorldConnection) SendPacketRelease(pkt *Packet) error {
	err := this.SendPacket(pkt)
	pkt.Release()

	return err
}

//  发送心跳数据
func (this *WorldConnection) SendHeartbeat() {
	zaplog.Debugf("WorldConnection %s 发送心跳", this)

	// 发送心跳数据
	pkt := NewPacket(protocol.C_PKT_ID_HEARTBEAT)
	this.SendPacket(pkt)
}

// 发送通用数据
func (this *WorldConnection) SendData(data []byte) {
	pkt := NewPacket(protocol.C_PKT_ID_DATA)
	pkt.AppendBytes(data)

	this.SendPacket(pkt)
}

// 刷新缓冲区
func (this *WorldConnection) Flush() error {
	return this.packetSocket.Flush()
}

// 打印信息
func (this *WorldConnection) String() string {
	return this.packetSocket.String()
}

// 处理 Packet 消息
func (this *WorldConnection) handlePacket(pkt *Packet) {
	// 根据类型处理数据
	switch pkt.pktId {
	case protocol.C_PKT_ID_HANDSHAKE: // 客户端握手请求
		this.handleHandshake(pkt.GetBody())
	case protocol.C_PKT_ID_HANDSHAKE_ACK: // 客户端握手 ACK
		this.handleHandshakeAck()
	default:
		zaplog.Errorf("WorldConnection &s 收到无效消息类型=%d，关闭连接", this, pkt.pktId)

		this.Close()
	}
}

//  处理握手消息
func (this *WorldConnection) handleHandshake(data []byte) {
	var err error

	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_INIT {
		return
	}

	// 消息解码
	req := &protocol.HandshakeReq{}
	err = msgpack.Unmarshal(data, req)
	if nil != err {
		zaplog.Errorf("WorldConnection %s 解码握手消息出错，关闭该连接", this)

		this.Close()
	}

	// 回复消息
	res := &protocol.HandshakeRes{
		Code:      protocol.OK,
		Heartbeat: this.option.Heartbeat,
	}
	var buf []byte
	var sucess bool = true

	// 版本验证
	if this.option.ShakeKey != "" && req.Key != this.option.ShakeKey {
		res.Code = protocol.SHAKE_KEY_ERROR
		sucess = false
	}

	// 通信方式验证,后续添加

	// 回复处理结果
	buf, err = msgpack.Marshal(res)
	if nil != err {
		zaplog.Errorf("WorldConnection %s 返回握手消息失败，编码握手消息出错", this)
	} else {
		this.handshakeResponse(sucess, buf)
	}

	// 握手失败，关闭连接
	if sucess == false {
		this.Close()
	}
}

//  返回握手消息
func (this *WorldConnection) handshakeResponse(sucess bool, data []byte) {
	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_INIT {
		return
	}

	// 返回数据
	pkt := NewPacket(protocol.C_PKT_ID_HANDSHAKE)
	pkt.AppendBytes(data)
	this.packetSocket.SendPacket(pkt) // 越过工作状态发送消息

	// 状态： 等待握手 ack
	if sucess {
		this.stateMgr.SetState(C_CONN_STATE_WAIT_ACK)
	}
}

//  处理握手ACK
func (this *WorldConnection) handleHandshakeAck() {
	// 状态：工作中
	if !this.stateMgr.SwapState(C_CONN_STATE_WAIT_ACK, C_CONN_STATE_WORKING) {

		return
	}

	// 发送心跳数据
	this.SendHeartbeat()
}
