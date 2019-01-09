// /////////////////////////////////////////////////////////////////////////////
// 同时支持 tcp websocket 的 连接器

package hybrid

import (
	"github.com/zpab123/world/worldnet"           // 网络库
	"github.com/zpab123/world/worldnet/connector" // 连接器
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

func init() {
	// 注册创建函数
	connector.RegisterCreator(newHybridConnector)
}

// /////////////////////////////////////////////////////////////////////////////
// hybridConnector 对象

// hybrid 接收器
type hybridConnector struct {
	parm *HybridParam // 创建参数
}

// 创建1个新的 hybridConnector 对象
//
// 返回 nil=创建失败
func newHybridConnector(opt *HybridParam) worldnet.IConnector {
	// 参数检查 (这里可先不用)
	//if nil != opt.Check() {
	//return nil
	//}

	// 创建对象
	cntor := &hybridConnector{
		parm: opt,
	}

	return cntor
}

// 启动 hybridConnector [IConnector 接口]
func (this *hybridConnector) Run() {

}

// 停止 hybridConnector [IConnector 接口]
func (this *hybridConnector) Stop() {

}

// 获取 connector 类型 [IConnector 接口]
func (this *hybridConnector) GetType() string {
	return connector.TYPE_HYBRID_CONNECTOR
}
