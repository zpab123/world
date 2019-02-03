// /////////////////////////////////////////////////////////////////////////////
// go 空文件模板

package dispatcher

// 分发客户端
type DispatcherClient struct {
}

// 新建1个 DispatcherClient
func NewDispatcherClient() *DispatcherClient {
	dc := &DispatcherClient{}

	return dc
}

// 启动
func (this *DispatcherClient) Run() {
	// 连接所有 DispatcherServer 服务器
}

// 关闭
func (this *DispatcherClient) Stop() {
	// 关闭所有 DispatcherServer 连接
}
