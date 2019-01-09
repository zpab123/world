// /////////////////////////////////////////////////////////////////////////////
// connector 构造参数

package connector

import (
	"fmt"
)

// /////////////////////////////////////////////////////////////////////////////
// 构造参数

// connector 构造参数
type Param struct {
	Host    string // IP地址
	MinPort uint32 // 最小端口
	MaxPort uint32 // 最大端口
}

// 参数检查,正确返回 true 错误 返回 fasle
func (this *Param) Check() (bool, error) {
	// 地址效验 -- 正则是否是ip 地址

	// 端口效验

	// 最大端口效验
	if this.MaxPort < this.MinPort {
		this.MaxPort = this.MinPort
	}

}

// 获取监听地址
//
// 返回格式为 scheme://host:minPort~maxPort/path
func (this *Param) GetAddrRange() (string, error) {
	// 参数检查
	_, err := this.Check()
	if nil != err {
		return nil, err
	}

	// 创建地址
	var addr string
	if this.MinPort == this.MaxPort {
		addr = fmt.Sprintf("%s:%d", this.Host, this.MinPort)
	}

	return addr, nil
}
