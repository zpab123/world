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
	// 获取读取 io.Reader
	reader, ok := ses.GetSocket().(io.Reader)

	// 转换错误，或者连接已经关闭时退出
	if !ok || nil == reader {
		return nil, nil
	}

	// 将 ses 所属 Connector 转化为 符合 IsocketOpt 的对象
	opt := ses.GetConnector().(iSocketOpt)

	if conn, ok := reader.(net.Conn); ok {
		// 有读超时时，设置超时
		opt.SetReadTimeout(conn, func() {
			// 读取数据
		})
	}

	return
}
