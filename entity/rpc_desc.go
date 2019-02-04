// /////////////////////////////////////////////////////////////////////////////
// rpc 调用 [代码完整]

package entity

import (
	"reflect"
	"strings"
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

const (
	rfServer      = 1 << iota
	rfOwnClient   = 1 << iota
	rfOtherClient = 1 << iota
)

// /////////////////////////////////////////////////////////////////////////////
// rpcDesc 对象

// rpcDesc 对象
type rpcDesc struct {
	Func       reflect.Value
	Flags      uint
	MethodType reflect.Type
	NumArgs    int
}

// /////////////////////////////////////////////////////////////////////////////
// rpcDescMap 对象

// rpcDescMap 对象
type rpcDescMap map[string]*rpcDesc

// 按照 method 的方法名字，保存 method
func (rdm rpcDescMap) visit(method reflect.Method) {
	methodName := method.Name
	var flag uint
	var rpcName string
	if strings.HasSuffix(methodName, "_Client") {
		flag |= rfServer + rfOwnClient
		rpcName = methodName[:len(methodName)-7]
	} else if strings.HasSuffix(methodName, "_AllClients") {
		flag |= rfServer + rfOwnClient + rfOtherClient
		rpcName = methodName[:len(methodName)-11]
	} else {
		// server method
		flag |= rfServer
		rpcName = methodName
	}

	methodType := method.Type
	rdm[rpcName] = &rpcDesc{
		Func:       method.Func,
		Flags:      flag,
		MethodType: methodType,
		NumArgs:    methodType.NumIn() - 1, // do not count the receiver
	}
}
