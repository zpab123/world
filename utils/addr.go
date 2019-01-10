// /////////////////////////////////////////////////////////////////////////////
// 监听地址工具

package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 变量
var (
	ErrInvalidPortRange = errors.New("无效的端口范围")
)

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// 将 addr 解析成 *Address 对象
//
// 参数 addr 的格式为 scheme://host:minPort~maxPort/path
func ParseAddress(addr string) (addrObj *Address, err error) {
	// 创建对象指针
	addrObj = new(Address)

	// 解析并保存 scheme
	schemePos := strings.Index(addr, "://")
	if -1 != schemePos {
		addrObj.Scheme = addr[:schemePos]
		addr = addr[schemePos+3:]
	}

	// 解析并保存 Host
	colonPos := strings.Index(addr, ":")
	if -1 != colonPos {
		addrObj.Host = addr[:colonPos]
	}
	addr = addr[colonPos+1:]

	// 解析 MinPort ~ MaxPort
	rangePos := strings.Index(addr, "~") // ~ 分隔符号为止
	var minStr, maxStr string            // 字符串形式的 port
	if rangePos != -1 {
		minStr = addr[:rangePos]
		slashPos := strings.Index(addr, "/")

		if slashPos != -1 {
			maxStr = addr[rangePos+1 : slashPos]
			addrObj.Path = addr[slashPos:]
		} else {
			maxStr = addr[rangePos+1:]
		}
	} else {
		slashPos := strings.Index(addr, "/")

		if slashPos != -1 {
			addrObj.Path = addr[slashPos:]
			minStr = addr[rangePos+1 : slashPos]
		} else {
			minStr = addr[rangePos+1:]
		}
	}

	// 解析 MinPort
	addrObj.MinPort, err = strconv.Atoi(minStr)
	if nil != err {
		return nil, ErrInvalidPortRange
	}

	// 解析 MaxPort
	if maxStr != "" {
		addrObj.MaxPort, err = strconv.Atoi(maxStr)
		if err != nil {
			return nil, ErrInvalidPortRange
		}
	} else {
		addrObj.MaxPort = addrObj.MinPort
	}

	return
}

// 在给定的端口范围内找到一个能用的端口 格式:
func DetectPort(addr string, fn func(a *Address, port int) (interface{}, error)) (interface{}, error) {
	// 将 addr 解析为 *Address 对象
	addrObj, err := ParseAddress(addr)
	if nil != err {
		return nil, err
	}

	// 查找端口
	for port := addrObj.MinPort; port <= addrObj.MaxPort; port++ {
		// 使用回调侦听
		ln, err := fn(addrObj, port)
		if nil == err {
			return ln, nil
		}

		// 达到最大端口
		if port == addrObj.MaxPort {
			return nil, err
		}
	}

	return nil, fmt.Errorf("绑定监听地址=%s；失败", addr)
}

// 将ip和端口合并为地址
func JoinAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// /////////////////////////////////////////////////////////////////////////////
// Address 对象

// 支持地址范围的格式
type Address struct {
	Scheme  string
	Host    string
	MinPort int
	MaxPort int
	Path    string
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

// 根据 port 参数，与 Address 对象的 host 组成1个 addr 字符
//
// 返回格式： 192.168.1.1:6002
func (this *Address) HostPortString(port int) string {
	return fmt.Sprintf("%s:%d", this.Host, port)
}

// 根据 port 参数，与 Address 对象的 Scheme host Path 组成1个完整 addr 字符
//
// 返回格式： http://192.168.1.1:6002/romte
func (this *Address) String(port int) string {
	if this.Scheme == "" {
		return this.HostPortString(port)
	}

	return fmt.Sprintf("%s://%s:%d%s", this.Scheme, this.Host, port, this.Path)
}
