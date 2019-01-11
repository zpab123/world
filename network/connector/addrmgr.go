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
	addr *model.Laddr
}

// 获取地址 [IAcceptor 接口]
func (this *AddrManager) GetAddr() *model.Laddr {
	return this.addr
}

// 设置地址 [IAcceptor 接口]
func (this *AddrManager) SetAddr(addrs *model.Laddr) {
	this.addr = addrs
}
