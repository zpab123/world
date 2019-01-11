// /////////////////////////////////////////////////////////////////////////////
// socket 地址参数

package connector

// /////////////////////////////////////////////////////////////////////////////
// Address 对象

// 监听地址
type AddrManager struct {
	addr *Laddr
}

// 获取地址 [IAcceptor 接口]
func (this *AddrManager) GetAddr() *Laddr {
	return this.addr
}

// 设置地址 [IAcceptor 接口]
func (this *AddrManager) SetAddr(addrs *Laddr) {
	this.addr = addrs
}
