// /////////////////////////////////////////////////////////////////////////////
// socket 错误工具 [代码完整]

package utils

import (
	"io"
	"net"
)

// io 错误检查
func IsEOFOrNetReadError(err error) bool {
	if err == io.EOF {
		return true
	}
	ne, ok := err.(*net.OpError)
	return ok && ne.Op == "read"
}
