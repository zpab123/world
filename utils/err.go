// /////////////////////////////////////////////////////////////////////////////
// 错误检查

package utils

import (
	"io"
	"net"

	"github.com/pkg/errors" // 错误检查工具
)

// /////////////////////////////////////////////////////////////////////////////
// 数据和接口

// 超时检查接口
type timeoutError interface {
	Timeout() bool // Is it a timeout error
}

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// err 是否是1个超时错误
func IsTimeoutError(err error) bool {
	if err == nil {
		return false
	}

	err = errors.Cause(err)
	ne, ok := err.(timeoutError)
	return ok && ne.Timeout()
}

// io 错误检查
func IsEOFOrNetReadError(err error) bool {
	if err == io.EOF {
		return true
	}
	ne, ok := err.(*net.OpError)
	return ok && ne.Op == "read"
}
