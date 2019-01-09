// /////////////////////////////////////////////////////////////////////////////
// socket 基础参数配置

package connector

// /////////////////////////////////////////////////////////////////////////////
// Property

// socket 基础参数配置
type BaseInfo struct {
	name string // 名字
	addr string // 监听地址
}

// 获取通讯端的名称
func (self *BaseInfo) GetName() string {
	return self.name
}

// 设置名字
func (self *BaseInfo) SetName(v string) {
	self.name = v
}

// 获取监听地址
func (self *BaseInfo) GetAddress() string {
	return self.addr
}

// 设置监听地址
func (self *BaseInfo) SetAddress(v string) {
	self.addr = v
}
