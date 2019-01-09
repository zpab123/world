// /////////////////////////////////////////////////////////////////////////////
// 消息元信息

package connector

import (
	"errors"
	"fmt"
	"path"
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	addError = errors.New("Address 数据错误")
)

// /////////////////////////////////////////////////////////////////////////////
// Address 对象

// 支持地址范围的格式
type Address struct {
	Scheme  string // 协议类型
	Host    string // ip 地址
	MinPort int    // 最小端口
	MaxPort int    // 最大端口
	Path    string // 包路径
}

// 参数检查,正确返回 true 错误 返回 fasle
func (this *Address) Check() (bool, error) {
	// 地址效验 -- 正则是否是ip 地址

	// 端口效验

	// 最大端口效验
	if this.MaxPort < this.MinPort {
		this.MaxPort = this.MinPort
	}

	return true, nil
}

// 获取带范围的 addr 格式
//
// 例如：scheme://host:minPort~maxPort/path
func (this *Address) GetAddrRange() (string, error) {
	// 参数检查
	ok := this.Check()
	if !ok {
		return nil, addError
	}

	// 获取地址
	var addr string
	if "" == this.Scheme {
		addr = fmt.Sprintf("%s:%d~%d", this.Host, this.MinPort, this.MaxPort)
	} else {
		addr = fmt.Sprintf("%s://%s:%d~%d/%s", this.Scheme, this.Host, this.MinPort, this.MaxPort, this.Path)
	}

	return addr
}

// 将 port 与 Host 组合成 192.168.1.101:6500 类似的格式
func (this *Address) HostPortString(port int) string {
	return fmt.Sprintf("%s:%d", this.Host, port)
}

// 获取完整的地址格式：格式 scheme://host:minPort~maxPort/path
func (this *Address) String(port int) string {
	if this.Scheme == "" {
		return this.HostPortString(port)
	}

	return fmt.Sprintf("%s://%s:%d%s", this.Scheme, this.Host, port, this.Path)
}
