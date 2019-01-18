// /////////////////////////////////////////////////////////////////////////////
// session 消息处理工具

package session

import (
	"github.com/zpab123/world/model"          // 全局模型
	"github.com/zpab123/world/network/packet" // packet 消息包
)

// 处理 packet 消息
func handlePacket(ses model.ISession, pkt *packet.Packet) {
	// 获取类型
	pktType := pkt.GetType()

	// 请求握手
	if model.C_PACKET_TYPE_HANDSHAKE == pktType {
		if model.C_SES_STATE_INITED != ses.GetState() {
			zplog.Debugf("客户端消息错误。非初始化状态下收到握手请求消息")

			return
		}

		handleHandshake(ses, pkt)
	}

	// 握手ATK
	if model.C_PACKET_TYPE_HANDSHAKE_ACK == pktType {
		if model.C_SES_STATE_WAIT_ACK != ses.GetState() {
			zplog.Debugf("客户端消息错误。非等待握手ACK状态下收到握手ACK消息")

			return
		}

		handleHandshakeAck()
	}

	// 心跳消息

}

// 处理握手消息
func handleHandshake(ses model.ISession, pkt *packet.Packet) {

}

// 处理握手ACK
func handleHandshakeAck(ses model.ISession, pkt *packet.Packet) {

}

// 处理心跳消息
func handleHeartbeat() {

}
