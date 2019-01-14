// /////////////////////////////////////////////////////////////////////////////
// io 异常捕获 [代码完整]

package network

// /////////////////////////////////////////////////////////////////////////////
// RecoverIoPanic 对象

// 是否捕获 Io 异常
type RecoverIoPanic struct {
	recoverIoPanic bool // 是否捕获异常
}

// 设置是否捕获 Io 异常 [IRecoverIoPanic 接口]
func (self *RecoverIoPanic) SetRecoverIoPanic(v bool) {
	self.recoverIoPanic = v
}

// 获取是否捕获 Io 异常 [IRecoverIoPanic 接口]
func (self *RecoverIoPanic) GetRecoverIoPanic() bool {
	return self.recoverIoPanic
}
