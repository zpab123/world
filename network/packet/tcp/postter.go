// /////////////////////////////////////////////////////////////////////////////
// tcp 消息收发器

package tcp

import (
	"io"
	"net"

	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// TcpPacketPostter 对象

// tcp 消息收发管理
type TcpPacketPostter struct {
}

// 接收消息
func (TcpPacketPostter) RecvPacket(ses worldnet.ISession) (pkt interface{}, err error) {

}
