// /////////////////////////////////////////////////////////////////////////////
// 实体基础 [代码完整]

package entity

import (
	"fmt"

	"github.com/zpab123/world/ids" // id 类
)

// /////////////////////////////////////////////////////////////////////////////
// Entity 对象

// Entity 对象
type Entity struct {
	ID       ids.EntityID // 实体ID
	TypeName string       // 实体类型
}

// String 接口
func (e *Entity) String() string {
	return fmt.Sprintf("%s<%s>", e.TypeName, e.ID)
}

// 实体初始化的时候调用
func (e *Entity) OnInit() {

}

// 实体属性准备好的时候调用
func (e *Entity) OnAttrsReady() {

}

// 实体被创建的时候调用
func (e *Entity) OnCreated() {
}

// 实体销毁的时候调用
func (e *Entity) OnDestroy() {
}

// 迁出的时候调用
func (e *Entity) OnMigrateOut() {
}

// 迁入的时候调用
func (e *Entity) OnMigrateIn() {
}

// 冻结的时候调用
func (e *Entity) OnFreeze() {
}

// 解冻的时候调用
func (e *Entity) OnRestored() {
}

// 进入某个 space 的时候调用
func (e *Entity) OnEnterSpace() {
}

// 离开某个 space 的时候调用
func (e *Entity) OnLeaveSpace() {
}
