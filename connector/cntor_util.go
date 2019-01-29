// /////////////////////////////////////////////////////////////////////////////
// 支持格式配置的连接器

package connector

import (
	"github.com/zpab123/world/network" // 网络库
)

// 创建1个新的 Acceptor
func newAcceptor(name string, addr *network.TLaddr, cntor *Connector) (network.IAcceptor, error) {
	// 根据名字创建 Acceptor。
	// 因为不同类型的连接器，参数不同，难以用同1个接口实现
	// 所以逐个创建

	// 接收器
	var aptor network.IAcceptor

	switch name {
	case network.C_ACCEPTOR_NAME_TCP: // tcp
		aptor = network.NewTcpAcceptor(addr, cntor)
		break
	case network.C_ACCEPTOR_NAME_WEBSOCKET: // websocket
		aptor = network.NewWsAcceptor(addr, cntor)
		break
	case network.C_ACCEPTOR_NAME_MUL: // tcp ws 混合模式
		aptor = network.NewMulAcceptor(addr, cntor)
		break
	case network.C_ACCEPTOR_NAME_COM: // tcp ws 组合模式
		aptor = network.NewComAcceptor(addr, cntor)
		break
	default:
		break
	}

	// 创建错误
	if nil == aptor {
		return nil, nil
	}

	return aptor, nil
}
