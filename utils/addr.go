// /////////////////////////////////////////////////////////////////////////////
// 监听地址工具

package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/zpab123/world/model" // 常用数据类型
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 变量
var (
	ErrInvalidPortRange = errors.New("无效的端口范围")
)

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// 将 addr 解析成 *model.Address 对象
//
// 参数 addr 的格式为 scheme://host:minPort~maxPort/path
func ParseAddress(addr string) (addrObj *model.Address, err error) {
	// 创建对象指针
	addrObj = new(model.Address)

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
func DetectPort(addr string, fn func(a *model.Address, port int) (interface{}, error)) (interface{}, error) {
	// 将 addr 解析为 *model.Address 对象
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
