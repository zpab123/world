// /////////////////////////////////////////////////////////////////////////////
// socket 基础参数配置

package connector

// /////////////////////////////////////////////////////////////////////////////
// BaseInfo 对象

// socket 基础信息
type BaseInfo struct {
	addr string // 监听地址
}

// 获取监听地址 [IConnector 接口]
func (this *BaseInfo) GetAddress() string {
	return self.addr
}

// 设置监听地址 [IConnector 接口]
func (this *BaseInfo) SetAddress(v string) {
	self.addr = v
}
