// /////////////////////////////////////////////////////////////////////////////
// 网络库

package network

// /////////////////////////////////////////////////////////////////////////////
// public api

// 根据名字创建1个新的 Acceptor
func NewAcceptor(name string, addr *TLaddr, mgr IComConnManager) IAcceptor {
	var aptor IAcceptor

	switch name {
	case C_ACCEPTOR_TYPE_TCP: // tcp
		aptor = NewTcpAcceptor(addr, mgr)
		break
	case C_ACCEPTOR_TYPE_WEBSOCKET: // websocket
		aptor = NewWsAcceptor(addr, mgr)
		break
	case C_ACCEPTOR_TYPE_MUL: // tcp ws 混合模式
		aptor = NewMulAcceptor(addr, mgr)
		break
	case C_ACCEPTOR_TYPE_COM: // tcp ws 组合模式
		aptor = NewComAcceptor(addr, mgr)
		break
	default:
		break
	}

	if nil == aptor {
		return nil
	}

	return aptor
}
