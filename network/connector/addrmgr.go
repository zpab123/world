// /////////////////////////////////////////////////////////////////////////////
// socket 地址参数

package connector

import (
	"github.com/zpab123/world/model" // 全局常量
)

// /////////////////////////////////////////////////////////////////////////////
// Address 对象

// 监听地址
type AddrManager struct {
	addr *model.TLaddr
}

// 获取地址 [IAcceptor 接口]
func (this *AddrManager) GetAddr() *model.TLaddr {
	return this.addr
}

// 设置地址 [IAcceptor 接口]
func (this *AddrManager) SetAddr(addrs *model.TLaddr) {
	this.addr = addrs
}
