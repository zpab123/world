// /////////////////////////////////////////////////////////////////////////////
// 游戏服务组件

package connector

import (
	"fmt"

	"github.com/zpab123/syncutil"    // 原子操作工具
	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/zplog"       // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

//  变量
var (
	acceptorMap = map[string]AcceptorCreateFunc{} // typeName->AcceptorCreateFunc 集合
)

// 常量
const (
	_maxConnNum uint32 = 100000 // 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 注册1个 connector 创建函数
func RegisterAcceptor(typeName string, f AcceptorCreateFunc) {
	// 已经存在
	if _, ok := acceptorMap[typeName]; ok {
		panic(fmt.Sprintf("注册 connector 重复，类型=%s", typeName))
	}

	// 保存类型
	acceptorMap[typeName] = f
}

// 根据类型，创建1个 acceptor 对象
//
// typeName 方便自己定义类型，不受 connectorOpt 影响
func NewAcceptor(typeName string, cntor model.IConnector) model.IAcceptor {
	// 类型检查
	creator := acceptorMap[typeName]
	if nil == creator {
		zplog.Panicf("创建 Acceptor 出错：找不到 %s 类型的 Acceptor", typeName)
		panic(fmt.Sprintf("创建 Acceptor 出错：找不到 %s 类型的 Acceptor", typeName))
	}

	// 创建 acceptor
	aptor := creator(cntor)

	// 设置地址参数
	aptor.SetAddr(cntor.GetAddr())

	return aptor
}

// /////////////////////////////////////////////////////////////////////////////
// connector 对象

// 网络连接对象，支持 websocket tcp
type Connector struct {
	name          string                // 组件名字
	laddr         *model.Laddr          // 监听地址集合
	connNum       syncutil.AtomicUint32 // 当前连接数
	opt           *model.ConnectorOpt   // 配置参数
	state         syncutil.AtomicInt32  // connector 当前状态
	acceptor      model.IAcceptor       // 某种类型的 acceptor 连接器
	SockerManager                       // 对象继承： socket 管理
}

// 新建1个 Connector 对象
func NewConnector(addrs *model.Laddr, opts *model.ConnectorOpt) model.IConnector {
	// 参数效验
	if nil != opts.Check() {
		return nil
	}

	// 地址检查？

	// 创建组件
	cntor := &Connector{
		name:  model.C_COMPONENT_NAME_CONNECTOR,
		laddr: addrs,
		opt:   opts,
	}

	// 创建 Acceptor
	aptor := NewAcceptor(opts.TypeName, cntor)

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

// 获取地址集合信息 [IConnector 接口]
func (this *Connector) GetAddr() *model.Laddr {
	return this.laddr
}

// 获取 connector 配置信息 [IConnector 接口]
func (this *Connector) GetConnectorOpt() *model.ConnectorOpt {
	return this.opt
}

// /////////////////////////////////////////////////////////////////////////////
// AcceptorCreateFunc 对象

type AcceptorCreateFunc func(cntor model.IConnector) model.IAcceptor

// /////////////////////////////////////////////////////////////////////////////
// socketManager

// socket 管理
type SockerManager struct {
}

// 收到1个新的 socket 连接 [IConnector] 接口
func (this *SockerManager) OnNewSocket(socket model.IPacketSocket) {

}

// 某个 socket  断开 [IConnector] 接口
func (this *SockerManager) OnSocketClose(socket model.IPacketSocket) {

}

// 某个 socket  收到数据 [IConnector] 接口
func (this *SockerManager) OnSocketMessage(socket model.IPacketSocket) {

}
