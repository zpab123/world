// /////////////////////////////////////////////////////////////////////////////
// DispatcherClient 连接器

package dispatcher

import (
	"github.com/zpab123/world/network" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// Connector

// 连接对象
type Connector struct {
	addr      *network.TLaddr       // 连接地址集合
	option    *TDispatcherClientOpt // 配置参数
	connector *network.Connector    // 连接器
}

// 新建1个 Connector
func NewConnector(addr *network.TLaddr, opt *TDispatcherClientOpt) *Connector {
	// 参数效验
	if nil == opt {
		opt = NewTDispatcherClientOpt(nil)
	}
	// 创建对象
	ct := network.NewConnector(addr, opt.ConnectorOpt)

	// 创建 Connector
	dct := &Connector{
		addr:      addr,
		option:    opt,
		connector: ct,
	}

	return dct
}

// 连接服务器
func (this *Connector) Connect() {
	this.connector.Connect()

	go this.recvLoop()

	go this.sendLoop()
}

// 接收线程
func (this *Connector) recvLoop() {
	for {
		// 接收消息
		pkt, err := this.connector.RecvPacket()

		if nil != err {
			break
		}

		if nil == pkt {
			continue
		}
	}
}

// 发送线程
func (this *Connector) sendLoop() {
	var err error
	for {
		// 刷新缓冲区
		err = this.connector.Flush()

		if nil != err {
			break
		}
	}
}
