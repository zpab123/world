// /////////////////////////////////////////////////////////////////////////////
// World 框架内部通信使用的 socket 对象

package network

import (
	"net"

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
	addr         *TLaddr             // 连接地址
	stateMgr     *state.StateManager // 状态管理
	option       *TWorldSocketOpt    // 配置参数
	packetSocket *PacketSocket       // PacketSocket 对象
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

	this.packetSocket = nil

	// 状态：关闭成功
	this.stateMgr.SetState(C_SOCKET_STATE_CLOSED)

	return
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
		zaplog.Debugf("WorldSocket 发送握手消息失败，状态错误。当前状态=%d，正确状态=%d", s, C_SOCKET_STATE_SHAKE)

		return
	}

	key := config.GetWorldConfig().ShakeKey

	zaplog.Debugf("WorldSocket 发送握手消息。key=%s", key)

	req := &msg.HandshakeReq{
		Key: key,
	}

	buf, err := proto.Marshal(req)

	if nil == err {
		pkt := NewPacket(C_PACKET_ID_HANDSHAKE)
		pkt.AppendBytes(buf)

		this.packetSocket.SendPacket(pkt)
	} else {
		zaplog.Debugf("WorldSocket 发送握手消息失败：protobuf 编码握手消息出错")
	}
}
