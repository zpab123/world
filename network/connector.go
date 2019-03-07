// /////////////////////////////////////////////////////////////////////////////
// World 框架内部通信使用的 socket 对象

package network

import (
	"net"
	"time"

	"github.com/pkg/errors"             // 错误库
	"github.com/zpab123/world/config"   // 配置文件
	"github.com/zpab123/world/protocol" // world 内部通信消息
	"github.com/zpab123/world/state"    // 状态管理
	"github.com/zpab123/zaplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// Connector

// World 框架内部通信使用的 socket 对象
type Connector struct {
	addr          *TLaddr             // 连接地址
	stateMgr      *state.StateManager // 状态管理
	option        *TConnectorOpt      // 配置参数
	packetSocket  *PacketSocket       // PacketSocket 对象
	timeOut       int64               // 心跳超时时间，单位：秒
	localTimeOut  int64               // 本地心跳超时时间点，精确到秒
	remoteTimeOut int64               // 远端心跳超时时间点，精确到秒
}

// 新建1个 Connector
func NewConnector(addr *TLaddr, opt *TConnectorOpt) *Connector {
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
	s := this.stateMgr.GetState()
	if s != C_CONN_STATE_INIT && s != C_CONN_STATE_CLOSED {
		err = errors.Errorf("Connector 连接失败，状态错误。当前状态=%d", s)

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
	}

	return err
}

// 关闭连接
func (this *Connector) Close() (err error) {
	// 状态效验
	s := this.stateMgr.GetState()
	if s == C_CONN_STATE_INIT {
		err = errors.New("Connector 关闭失败：它处于init状态")

		return
	}

	if s == C_CONN_STATE_CLOSED {
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

	// 记录后端超时
	if this.timeOut > 0 {
		this.remoteTimeOut = time.Now().Unix() + this.timeOut
	}

	// 握手消息
	if pkt.pktId == C_PACKET_ID_HANDSHAKE {
		this.handleHandshake(pkt)

		return nil, err
	}

	// 心跳消息
	if pkt.pktId == C_PACKET_ID_HEARTBEAT {
		zaplog.Debugf("收到 server 心跳消息")
		//this.handleHeartbeat()

		return nil, err
	}

	// 状态效验
	s := this.stateMgr.GetState()
	if s != C_CONN_STATE_WORKING {
		this.Close()

		err = errors.Errorf("Connector 状态错误。当前状态=%d，正确状态=%d", s, C_CONN_STATE_WORKING)

		return nil, err
	}

	return pkt, nil
}

// 发送1个 packet 消息
func (this *Connector) SendPacket(pkt *Packet) error {
	// 状态效验

	// 记录前端超时
	if this.timeOut > 0 {
		this.localTimeOut = time.Now().Unix() + this.timeOut
	}

	return this.packetSocket.SendPacket(pkt)
}

// 刷新缓冲区
func (this *Connector) Flush() error {
	if nil == this.packetSocket {
		return nil
	}

	return this.packetSocket.Flush()
}

// 检查本地心跳
func (this *Connector) CheckLocalHeartbeat() {
	if this.stateMgr.GetState() != C_CONN_STATE_WORKING {
		return
	}

	if this.timeOut > 0 {
		if time.Now().Unix() >= this.localTimeOut {

			this.sendHeartbeat()
		}
	}
}

// 检查远端心跳
func (this *Connector) CheckRemoteHeartbeat() error {
	if this.stateMgr.GetState() != C_CONN_STATE_WORKING {
		return nil
	}

	if this.timeOut > 0 {
		if time.Now().Unix() >= this.remoteTimeOut {
			zaplog.Warnf("Connector 后端心跳超时，断开连接")

			return this.Close()
		}
	}

	return nil
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
	conn, err := net.Dial("tcp", this.addr.TcpAddr)

	if nil != err {
		return err
	}

	tcpConn, ok := conn.(*net.TCPConn)
	if ok {
		tcpConn.SetReadBuffer(this.option.TcpConnOpt.ReadBufferSize)
		tcpConn.SetWriteBuffer(this.option.TcpConnOpt.WriteBufferSize)
		tcpConn.SetNoDelay(this.option.TcpConnOpt.NoDelay)
	}

	zaplog.Debugf("Connector 连接tcp服务器成功。ip=%s", this.addr.TcpAddr)

	this.createSocket(conn)

	return nil
}

// 发送握手请求
func (this *Connector) sendHandshake() (err error) {
	// 状态效验
	s := this.stateMgr.GetState()
	if s != C_CONN_STATE_SHAKE {
		err = errors.Errorf("Connector 发送握手消息失败，状态错误。当前状态=%d，正确状态=%d", s, C_CONN_STATE_SHAKE)

		return
	}

	key := config.GetWorldConfig().ShakeKey

	pkt := NewPacket(C_PACKET_ID_HANDSHAKE)
	pkt.AppendString(key)

	this.SendPacket(pkt)

	return
}

//  处理握手消息
func (this *Connector) handleHandshake(pkt *Packet) {
	// 解码
	code := pkt.ReadUint32()
	heartbeat := pkt.ReadUint32()

	// 握手结果
	if code == protocol.OK {
		// 保存握手数据
		t := int64(heartbeat * 2)
		if t > 0 {
			this.remoteTimeOut = time.Now().Unix() + t
			this.timeOut = t
		}

		// 发送 ack
		this.sendAck()
	} else {
		zaplog.Error("Connector 握手失败。code=%d", code)

		this.Close()
	}
}

// 发送握手ACK
func (this *Connector) sendAck() {
	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_SHAKE {

		return
	}

	pkt := NewPacket(C_PACKET_ID_HANDSHAKE_ACK)

	this.SendPacket(pkt)

	// 状态： 工作中
	this.stateMgr.SetState(C_CONN_STATE_WORKING)
}

//  发送心跳数据
func (this *Connector) sendHeartbeat() {
	// 状态效验
	if this.stateMgr.GetState() != C_CONN_STATE_WORKING {

		return
	}

	zaplog.Debugf("client 发送心跳")

	// 发送心跳数据
	pkt := NewPacket(C_PACKET_ID_HEARTBEAT)
	this.SendPacket(pkt)
}

//  处理心跳消息
func (this *Connector) handleHeartbeat() {

}
