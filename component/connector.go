// /////////////////////////////////////////////////////////////////////////////
// 游戏服务组件

package component

import (
	"github.com/zpab123/syncutil"                // 原子操作工具
	"github.com/zpab123/world/consts"            // 全局常量
	"github.com/zpab123/world/model"             // 全局结构体
	"github.com/zpab123/world/network"           // 网络库
	"github.com/zpab123/world/network/connector" // 网络连接库
	"github.com/zpab123/zplog"                   // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 常量
const (
	_maxConnNum uint32 = 100000 // 默认最大连接数
)

// /////////////////////////////////////////////////////////////////////////////
// connector 对象

// 网络连接对象，支持 websocket tcp
type Connector struct {
	name      string                  // 组件名字
	laddr     *model.Laddr            // 监听地址集合
	connNum   syncutil.AtomicUint32   // 当前连接数
	opt       *connector.ConnectorOpt // 配置参数
	state     syncutil.AtomicInt32    // connector 当前状态
	connector network.IConnector      // 某种类型的 connector 连接器
}

// 新建1个 Connector 对象
func NewConnector(addrs *model.Laddr, param *connector.ConnectorOpt) *Connector {
	// 参数效验
	if nil != parameter.Check() {
		return nil
	}

	// 创建 connector
	cntor := connector.NewConnector(addrs, param)

	// 创建组件
	cpt := &Connector{
		name:  consts.COMPONENT_NAME_CONNECTOR,
		laddr: addrs,
		opt:   param,
	}

	// 保存 Connector
	cpt.connector = cntor

	return cpt
}

// 组件名字 [IComponent 实现]
func (this *Connector) Name() string {
	return this.name
}

// 运行 Connector 组件 [IComponent 实现]
func (this *Connector) Run() {
	// 启动 connector
	this.connector.Run()
}

// 停止运行 [IComponent 实现]
func (this *Connector) Stop() {
	// 停止 connector
	this.connector.Stop()
}
