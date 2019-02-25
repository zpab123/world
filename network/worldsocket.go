// /////////////////////////////////////////////////////////////////////////////
// World 框架内部通信使用的 socket 对象

package network

import (
	"net"
	"time"

	"github.com/gogo/protobuf/proto"  // protobuf 库
	"github.com/pkg/errors"           // 错误库
	"github.com/zpab123/world/config" // 配置文件
	"github.com/zpab123/world/msg"    // world 内部通信消息
	"github.com/zpab123/world/state"  // 状态管理
	"github.com/zpab123/zaplog"       // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// WorldSocket

// World 框架内部通信使用的 socket 对象
type WorldSocket struct {
	addr          *TLaddr             // 连接地址
	stateMgr      *state.StateManager // 状态管理
	option        *TWorldSocketOpt    // 配置参数
	packetSocket  *PacketSocket       // PacketSocket 对象
	timeOut       int64               // 心跳超时时间，单位：秒
	localTimeOut  int64               // 本地心跳超时时间点，精确到秒
	remoteTimeOut int64               // 远端心跳超时时间点，精确到秒
}

// 新建1个 WorldSocket
func NewWorldSocket(addr *TLaddr, opt *TWorldSocketOpt) *WorldSocket {
	// 创建对象
	st := state.NewStateManager()

	// 创建 WorldSocket
	ws := &WorldSocket{
		addr:     addr,
		stateMgr: st,
		option:   opt,
	}

	// 状态： init
	ws.stateMgr.SetState(C_SOCKET_STATE_INIT)

	return ws
}

// 连接服务器
func (this *WorldSocket) Connect() error {
	var err error

	// 状态效验
	s := this.stateMgr.GetState()
	if s != C_SOCKET_STATE_INIT && s != C_SOCKET_STATE_CLOSED {
		err = errors.Errorf("WorldSocket 连接失败，状态错误。当前状态=%d", s)

		return err
	}

	// 根据类型创建连接
	var netType string
	if nil != this.option {
		netType = this.option.NetType
	}

	switch netType {
	case C_NET_TYPE_TCP: // tcp
		err = this.connectTcp()
		break
	default:
		err = this.connectTcp()
		break
	}

	// 发送握手请求
	if nil == err {
		this.stateMgr.SetState(C_SOCKET_STATE_SHAKE)
		this.sendHandshake()
	}

	return err
}

// 关闭连接
func (this *WorldSocket) Close() (err error) {
	// 状态效验
	s := this.stateMgr.GetState()
	if s == C_SOCKET_STATE_INIT {
		err = errors.New("WorldSocket 关闭失败：它处于init状态")

		return
	}

	if s == C_SOCKET_STATE_CLOSED {
		err = errors.New("WorldSocket 关闭失败：它已经处于关闭状态")

		return
	}

	if nil == this.packetSocket {
		err = errors.New("WorldSocket 关闭失败：packetSocket 不存在")

		return
	} else {
		err = this.packetSocket.Close()
	}

	// this.packetSocket = nil

	// 状态：关闭成功
	this.stateMgr.SetState(C_SOCKET_STATE_CLOSED)

	return
}

// 接收1个 Packet 消息
func (this *WorldSocket) RecvPacket() (*Packet, error) {
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
		this.handleHandshake(pkt.GetBody())

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
	if s != C_SOCKET_STATE_WORKING {
		this.Close()

		err = errors.Errorf("WorldSocket 状态错误。当前状态=%d，正确状态=%d", s, C_SOCKET_STATE_WORKING)

		return nil, err
	}

	return pkt, nil
}

// 发送1个 packet 消息
func (this *WorldSocket) SendPacket(pkt *Packet) error {
	// 状态效验

	// 记录前端超时
	if this.timeOut > 0 {
		this.localTimeOut = time.Now().Unix() + this.timeOut
	}

	return this.packetSocket.SendPacket(pkt)
}

// 刷新缓冲区
func (this *WorldSocket) Flush() error {
	if nil == this.packetSocket {
		return nil
	}

	return this.packetSocket.Flush()
}

// 检查本地心跳
func (this *WorldSocket) CheckLocalHeartbeat() {
	if this.stateMgr.GetState() != C_SOCKET_STATE_WORKING {
		return
	}

	if this.timeOut > 0 {
		if time.Now().Unix() >= this.localTimeOut {

			this.sendHeartbeat()
		}
	}
}

// 检查远端心跳
func (this *WorldSocket) CheckRemoteHeartbeat() error {
	if this.stateMgr.GetState() != C_SOCKET_STATE_WORKING {
		return nil
	}

	if this.timeOut > 0 {
		if time.Now().Unix() >= this.remoteTimeOut {
			zaplog.Warnf("WorldSocket 后端心跳超时，断开连接")

			return this.Close()
		}
	}

	return nil
}

// 创建统一 socket
func (this *WorldSocket) createSocket(conn net.Conn) {
	// 创建 packetSocket
	socket := &Socket{
		Conn: conn,
	}
	bufSocket := NewBufferSocket(socket, this.option.BuffSocketOpts)

	this.packetSocket = NewPacketSocket(bufSocket)
}

// 连接 tcp
func (this *WorldSocket) connectTcp() error {
	conn, err := net.Dial("tcp", this.addr.TcpAddr)
	if nil != err {
		zaplog.Errorf("WorldSocket 连接tcp服务器失败。ip=%s", this.addr.TcpAddr)

		return err
	}

	tcpConn, ok := conn.(*net.TCPConn)
	if ok {
		tcpConn.SetReadBuffer(this.option.TcpConnOpt.ReadBufferSize)
		tcpConn.SetWriteBuffer(this.option.TcpConnOpt.WriteBufferSize)
		tcpConn.SetNoDelay(this.option.TcpConnOpt.NoDelay)
	}

	zaplog.Debugf("WorldSocket 连接tcp服务器成功。ip=%s", this.addr.TcpAddr)

	this.createSocket(conn)

	return nil
}

// 发送握手请求
func (this *WorldSocket) sendHandshake() {
	// 状态效验
	s := this.stateMgr.GetState()
	if s != C_SOCKET_STATE_SHAKE {
		zaplog.Errorf("WorldSocket 发送握手消息失败，状态错误。当前状态=%d，正确状态=%d", s, C_SOCKET_STATE_SHAKE)

		return
	}

	key := config.GetWorldConfig().ShakeKey

	req := &msg.HandshakeReq{
		Key: key,
	}

	buf, err := proto.Marshal(req)

	if nil == err {
		pkt := NewPacket(C_PACKET_ID_HANDSHAKE)
		pkt.AppendBytes(buf)

		this.SendPacket(pkt)
	} else {
		zaplog.Errorf("WorldSocket 发送握手消息失败：protobuf 编码握手消息出错")
	}
}

//  处理握手消息
func (this *WorldSocket) handleHandshake(data []byte) {
	// 解码
	res := &msg.HandshakeRes{}
	err := proto.Unmarshal(data, res)
	if nil != err {
		zaplog.Error("WorldSocket: protobuf 解码握手消息出错。关闭连接")

		this.Close()

		return
	}

	// 握手结果
	if res.Code == msg.OK {
		// 保存握手数据
		t := int64(res.Heartbeat * 2)
		if t > 0 {
			this.remoteTimeOut = time.Now().Unix() + t
			this.timeOut = t
		}

		// 发送 ack
		this.sendAck()
	} else {
		zaplog.Error("WorldSocket 握手失败。code=", res.Code)

		this.Close()
	}
}

// 发送握手ACK
func (this *WorldSocket) sendAck() {
	// 状态效验
	if this.stateMgr.GetState() != C_SOCKET_STATE_SHAKE {

		return
	}

	pkt := NewPacket(C_PACKET_ID_HANDSHAKE_ACK)

	this.SendPacket(pkt)

	// 状态： 工作中
	this.stateMgr.SetState(C_SOCKET_STATE_WORKING)
}

//  发送心跳数据
func (this *WorldSocket) sendHeartbeat() {
	// 状态效验
	if this.stateMgr.GetState() != C_SOCKET_STATE_WORKING {

		return
	}

	zaplog.Debugf("client 发送心跳")

	// 发送心跳数据
	pkt := NewPacket(C_PACKET_ID_HEARTBEAT)
	this.SendPacket(pkt)
}

//  处理心跳消息
func (this *WorldSocket) handleHeartbeat() {

}
