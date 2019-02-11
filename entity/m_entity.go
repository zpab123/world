// /////////////////////////////////////////////////////////////////////////////
// 常量-接口-types

package entity

// /////////////////////////////////////////////////////////////////////////////
// 常量

const (
	ENTITY_TYPE_SPACE = "__space__" // space 类型 entity
)

// 实体属性类型
const (
	ATTR_DEFS_CLIENT = "Client"     // 客户端属性
	ATTR_DEFS_ALL    = "AllClients" // 客户端服务器共有
	ATTR_DEFS_PER    = "Persistent" // 需要持久化的属性
)

// /////////////////////////////////////////////////////////////////////////////
// entity 类型实体接口

// entity 类型实体接口
type IEntity interface {
	OnInit()                                // 实体初始化的时候调用
	OnAttrsReady()                          // 实体属性准备好的时候调用
	OnCreated()                             // 实体被创建的时候调用
	OnDestroy()                             // 实体销毁的时候调用
	OnMigrateOut()                          // 迁出的时候调用
	OnMigrateIn()                           // 迁入的时候调用
	OnFreeze()                              // 冻结的时候调用
	OnRestored()                            // 解冻的时候调用
	OnEnterSpace()                          // 进入某个 space 的时候调用
	OnLeaveSpace(space *Space)              // 离开某个 space 的时候调用
	SetEntityTypeDesc(desc *EntityTypeDesc) // 设置实体描述文件属性，注册时调用
}

// /////////////////////////////////////////////////////////////////////////////
// space 类型实体接口

// space 类型实体接口
type ISpace interface {
	IEntity          // 继承 IEntity 方法
	OnSpaceInit()    // 当 space 初始化的时候调用
	OnSpaceCreated() // 当 space 被创建的时候调用
	OnSpaceDestroy() // 当 space 被销毁的时候调用
	OnEntityEnter()  // 当某个 entity 进入的时候调用
	OnEntityLeave()  // 当某个 entity 离开的时候调用
	OnGameReady()    // 仅仅被 nil space 调用
}
