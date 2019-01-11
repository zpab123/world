// /////////////////////////////////////////////////////////////////////////////
// 游戏服务组件

package connector

import (
	"fmt"

	"github.com/zpab123/syncutil"     // 原子操作工具
	"github.com/zpab123/world/consts" // 全局常量
	"github.com/zpab123/world/ifs"    // 顶级接口库
	"github.com/zpab123/zplog"        // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

//  变量
var (
	creatorMap = map[string]CreateFunc{} // typeName->CreateFunc 集合
)

// 常量
const (
	_maxConnNum uint32 = 100000 // 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 注册1个 connector 创建函数
func RegisterCreator(f CreateFunc) {
	// 临时实例化一个，获取类型
	cntor := f()

	// 已经存在
	if _, ok := creatorMap[cntor.GetType()]; ok {
		panic(fmt.Sprintf("注册 connector 重复，类型=%s", cntor.GetType()))
	}

	// 保存类型
	creatorMap[cntor.GetType()] = f
}

// 根据类型，创建1个 acceptor 对象
func NewAcceptor(addr *Laddr, opts *ConnectorOpt) IAcceptor {
	// 获取类型
	typeName := opts.TypeName

	// 类型检查
	creator := creatorMap[typeName]
	if nil == creator {
		zplog.Panicf("创建 Acceptor 出错：找不到 %s 类型的 Acceptor", typeName)
		panic(fmt.Sprintf("创建 Acceptor 出错：找不到 %s 类型的 Acceptor", typeName))
	}

	// 地址检查

	// 参数检查
	opts.Check()

	// 创建 acceptor
	aptor := creator()

	// 设置地址参数
	aptor.SetAddr(addr)

	return aptor
}

// /////////////////////////////////////////////////////////////////////////////
// connector 对象

// 网络连接对象，支持 websocket tcp
type Connector struct {
	name     string                // 组件名字
	laddr    *Laddr                // 监听地址集合
	connNum  syncutil.AtomicUint32 // 当前连接数
	opt      *ConnectorOpt         // 配置参数
	state    syncutil.AtomicInt32  // connector 当前状态
	acceptor IAcceptor             // 某种类型的 acceptor 连接器
}

// 新建1个 Connector 对象
func NewConnector(addrs *Laddr, param *ConnectorOpt) ifs.IComponent {
	// 参数效验
	if nil != param.Check() {
		return nil
	}

	// 创建 Acceptor
	aptor := NewAcceptor(addrs, param)

	// 创建组件
	cntor := &Connector{
		name:  consts.COMPONENT_NAME_CONNECTOR,
		laddr: addrs,
		opt:   param,
	}

	// 保存 Acceptor
	cntor.acceptor = aptor

	return cntor
}

// 组件名字 [IComponent 实现]
func (this *Connector) Name() string {
	return this.name
}

// 运行 Connector 组件 [IComponent 实现]
func (this *Connector) Run() {
	// 启动 acceptor
	this.acceptor.Run()
}

// 停止运行 [IComponent 实现]
func (this *Connector) Stop() {
	// 停止 acceptor
	this.acceptor.Stop()
}

// /////////////////////////////////////////////////////////////////////////////
// CreateFunc 对象

type CreateFunc func() IAcceptor

// /////////////////////////////////////////////////////////////////////////////
// socketManager

// socket 管理
type SockerManager struct {
}

// 收到1个新的 socket 连接
func (this *SockerManager) onNewSocket() {

}

// 某个 socket  断开
func (this *SockerManager) onSocketClose() {

}

// 某个 socket  收到数据
func (this *SockerManager) onSocketMessage() {

}