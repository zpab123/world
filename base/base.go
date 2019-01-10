// /////////////////////////////////////////////////////////////////////////////
// 服务器公用的一些基础 信息

package base

// /////////////////////////////////////////////////////////////////////////////
// Cmd 对象

// 启动信息
type BaseInfo struct {
	runer      string // 启动者
	serverType string // 服务器类型
	mainPath   string // main 程序所在路径
}

// 获取 Runer
func (this *BaseInfo) GetRuner() string {
	return this.runer
}

// 设置 Runer
func (this *BaseInfo) SetRuner(v string) {
	this.runer = v
}

// 获取服务器类型
func (this *BaseInfo) GetServerType() string {
	return this.serverType
}

// 设置服务器类型
func (this *BaseInfo) SetServerType(v string) {
	this.serverType = v
}

// 获取 main 程序所在路径
func (this *BaseInfo) GetMainPath() string {
	return this.mainPath
}

// 设置 main 程序所在路径
func (this *BaseInfo) SetMainPath(v string) {
	this.mainPath = v
}
