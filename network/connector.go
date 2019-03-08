// /////////////////////////////////////////////////////////////////////////////
// World 框架内部通信使用的 socket 对象

package network

import (
	"net"

	"github.com/pkg/errors"             // 错误库
	"github.com/vmihailenco/msgpack"    // 消息编码/解码
	"github.com/zpab123/world/config"   // 配置文件
	"github.com/zpab123/world/protocol" // world 内部通信消息
	"github.com/zpab123/world/state"    // 状态管理
	"github.com/zpab123/zaplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// Connector

// World 框架内部通信使用的 socket 对象
type Connector struct {
	addr         *TLaddr             // 连接地址
	stateMgr     *state.StateManager // 状态管理
	option       *TConnectorOpt      // 配置参数
	packetSocket *PacketSocket       // PacketSocket 对象
	heartbeat    uint32              // 心跳周期
}

// 新建1个 Connector
func NewConnector(addr *TLaddr, opt *TConnectorOpt) *Connector {
	// 参数效验
	if nil == opt {
		opt = NewTConnectorOpt()
	}

	// 创建对象
	st := state.NewStateManager()

	// 创建 Connector
	ctor := &Connector{
		addr:     addr,
		stateMgr: st,
		option:   opt,
	}

	// 状态： init
	ctor.stateMgr.SetState(C_CONN_STATE_INIT)

	return ctor
}

// 连接服务器
func (this *Connector) Connect() error {
	var err error

	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_INIT && this.stateMgr.GetState() != C_CONN_STATE_CLOSED {
		err = errors.Errorf("Connector %s 连接失败，状态错误。当前状态=%d", this.stateMgr.GetState())

		return err
	}

	// 根据类型创建连接
	var netType string
	if nil != this.option {
		netType = this.option.NetType
	}

	switch netType {
	case C_CONNECTOR_TCP: // tcp
		err = this.connectTcp()
	default:
		err = this.connectTcp()
	}

	// 发送握手请求
	if nil == err {
		this.stateMgr.SetState(C_CONN_STATE_SHAKE)
		err = this.sendHandshake()

		if nil != err {
			this.Close()

			return err
		}
	}

	return err
}

// 关闭连接
func (this *Connector) Close() (err error) {
	// 状态效验
	if this.stateMgr.GetState() == C_CONN_STATE_INIT {
		err = errors.New("Connector 关闭失败：它处于init状态")

		return
	}

	if this.stateMgr.GetState() == C_CONN_STATE_CLOSED {
		err = errors.New("Connector 关闭失败：它已经处于关闭状态")

		return
	}

	if nil == this.packetSocket {
		err = errors.New("Connector 关闭失败：packetSocket 不存在")

		return
	} else {
		err = this.packetSocket.Close()
	}

	// this.packetSocket = nil

	// 状态：关闭成功
	this.stateMgr.SetState(C_CONN_STATE_CLOSED)

	return
}

// 接收1个 Packet 消息
func (this *Connector) RecvPacket() (*Packet, error) {
	// 接收 packet
	if nil == this.packetSocket {
		return nil, nil
	}

	pkt, err := this.packetSocket.RecvPacket()
	if nil == pkt || nil != err {
		return nil, err
	}

	// 握手消息
	if pkt.pktId == protocol.C_PKT_ID_HANDSHAKE {
		this.handleHandshake(pkt.GetBody())

		return nil, nil
	}

	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_WORKING {
		err = errors.New("Connector 在非 workling 状态下收到数据，关闭 Connector")

		this.Close()

		return nil, err
	}

	return pkt, nil
}

// 发送1个 packet 消息
func (this *Connector) SendPacket(pkt *Packet) error {
	// 状态效验
	return this.packetSocket.SendPacket(pkt)
}

// 刷新缓冲区
func (this *Connector) Flush() error {
	if nil == this.packetSocket {
		return nil
	}

	return this.packetSocket.Flush()
}

// 打印信息
func (this *Connector) String() string {
	if nil == this.packetSocket {
		return ""
	}

	return this.packetSocket.String()
}

// 创建统一 socket
func (this *Connector) createSocket(conn net.Conn) {
	// 创建 packetSocket
	socket := &Socket{
		Conn: conn,
	}
	bufSocket := NewBufferSocket(socket, this.option.BuffSocketOpts)

	this.packetSocket = NewPacketSocket(bufSocket)
}

// 连接 tcp
func (this *Connector) connectTcp() error {
	var err error
	var conn net.Conn

	// 地址效验
	if "" == this.addr.TcpAddr {
		err = errors.New("连接 tcp 服务器失败，TcpAddr 为空")

		return err
	}

	// 连接
	conn, err = net.Dial("tcp", this.addr.TcpAddr)
	if nil != err {
		return err
	}

	tcpConn, ok := conn.(*net.TCPConn)
	if ok {
		tcpConn.SetReadBuffer(this.option.TcpConnOpt.ReadBufferSize)
		tcpConn.SetWriteBuffer(this.option.TcpConnOpt.WriteBufferSize)
		tcpConn.SetNoDelay(this.option.TcpConnOpt.NoDelay)
	}

	zaplog.Debugf("Connector %s 连接tcp服务器成功", this)

	this.createSocket(conn)

	return nil
}

// 发送握手请求
func (this *Connector) sendHandshake() (err error) {
	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_SHAKE {
		err = errors.Errorf("Connector %s 发送握手消息失败，状态错误。当前状态=%d，正确状态=%d", this, this.stateMgr.GetState(), C_CONN_STATE_SHAKE)

		return
	}

	// 发送消息
	shakeKey := config.GetWorldConfig().ShakeKey
	req := &protocol.HandshakeReq{
		Key:      shakeKey,
		Acceptor: 0,
	}

	data, err := msgpack.Marshal(req)
	if nil != err {
		err = errors.Errorf("Connector %s 发送握手消息失败，msgpack 编码错误", this)

		return
	}

	pkt := NewPacket(protocol.C_PKT_ID_HANDSHAKE)
	pkt.AppendBytes(data)

	this.SendPacket(pkt)

	return
}

//  处理握手消息
func (this *Connector) handleHandshake(data []byte) {
	var err error

	// 解码
	res := &protocol.HandshakeRes{}
	err = msgpack.Unmarshal(data, res)
	if nil != err {
		zaplog.Error("Connector 解码握手消息失败，关闭 Connector")

		this.Close()
	}

	// 握手结果
	if res.Code == protocol.OK {
		zaplog.Debugf("Connector %s 握手成功", this)

		this.heartbeat = res.Heartbeat // 保存握手数据
		this.sendAck()                 // 发送 ack
	} else {
		zaplog.Errorf("Connector 握手失败，关闭 Connector。code=%d", res.Code)

		this.Close()
	}
}

// 发送握手ACK
func (this *Connector) sendAck() {
	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_SHAKE {

		return
	}

	pkt := NewPacket(protocol.C_PKT_ID_HANDSHAKE_ACK)
	this.SendPacket(pkt)

	// 状态： 工作中
	this.stateMgr.SetState(C_CONN_STATE_WORKING)
}
