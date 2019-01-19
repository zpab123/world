// /////////////////////////////////////////////////////////////////////////////
// 支持格式配置的连接器

package connector

import (
	"fmt"

	"github.com/zpab123/syncutil"      // 原子操作工具
	"github.com/zpab123/world/model"   // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/world/network" // 网络库
	"github.com/zpab123/world/session" // 会话组件
	"github.com/zpab123/zplog"         // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 常量
const (
	_maxConnNum uint32 = 100000 // 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// /////////////////////////////////////////////////////////////////////////////
// connector 对象

// 网络连接对象，支持 websocket tcp
type Connector struct {
	laddr       *model.TLaddr         // 监听地址集合
	opts        *model.TConnectorOpt  // 配置参数
	acceptor    model.IAcceptor       // 某种类型的 acceptor 连接器
	connNum     syncutil.AtomicUint32 // 当前连接数
	state       syncutil.AtomicInt32  // connector 当前状态
	tcpAcceptor *network.TcpAcceptor  // tcpAcceptor 连接器
	wsAcceptor  *network.WsAcceptor   // wsAcceptor 连接器
	mulAcceptor *network.MulAcceptor  // mulAcceptor 连接器
	comAcceptor *network.ComAcceptor  // comAcceptor 连接器
}

// 新建1个 Connector 对象
func NewConnector(addrs *model.TLaddr, opt *model.TConnectorOpt) model.IConnector {
	// 参数效验
	if nil != opt.Check() {
		return nil
	}

	// 地址检查？

	// 创建组件
	cntor := &Connector{
		laddr: addrs,
		opts:  opt,
	}

	// 创建 Acceptor
	aptor := NewAcceptor(opts.AcceptorType, cntor)
	cntor.acceptor = aptor

	return cntor
}

// /////////////////////////////////////////////////////////////////////////////
// private api

// 创建1个新的 Acceptor
func newAcceptor(name string) model.IAcceptor {

}
