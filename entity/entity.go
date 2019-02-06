// /////////////////////////////////////////////////////////////////////////////
// 实体基础

package entity

import (
	"fmt"
	"reflect"

	"github.com/zpab123/world/ids" // id 类
)

// /////////////////////////////////////////////////////////////////////////////
// Entity 对象

// Entity 对象
type Entity struct {
	Id       ids.EntityID    // 实体ID
	TypeName string          // 实体类型
	Value    reflect.Value   // 通过 reflect 创建的对象
	I        IEntity         // 符合 IEntity 接口的对象
	typeDesc *EntityTypeDesc // 实体描述文件
	Attrs    *Attr           // 属性集合
	Space    *Space          // 实体所在空间
}

// String 接口
func (this *Entity) String() string {
	return fmt.Sprintf("%s<%s>", this.TypeName, this.Id)
}

// 实体初始化的时候调用
func (this *Entity) OnInit() {

}

// 实体属性准备好的时候调用
func (this *Entity) OnAttrsReady() {

}

// 实体被创建的时候调用
func (this *Entity) OnCreated() {
}

// 实体销毁的时候调用
func (this *Entity) OnDestroy() {
}

// 迁出的时候调用
func (this *Entity) OnMigrateOut() {
}

// 迁入的时候调用
func (this *Entity) OnMigrateIn() {
}

// 冻结的时候调用
func (this *Entity) OnFreeze() {
}

// 解冻的时候调用
func (this *Entity) OnRestored() {
}

// 进入某个 space 的时候调用
func (this *Entity) OnEnterSpace() {
}

// 离开某个 space 的时候调用
func (this *Entity) OnLeaveSpace() {
}

// 初始化1个 Entity
func (this *Entity) init(typeName string, entityid ids.EntityID, entityInstance reflect.Value) {
	this.Id = entityid
	this.Value = entityInstance
	this.I = entityInstance.Interface().(IEntity)
	this.TypeName = typeName

	this.typeDesc = registeredEntityTypes[typeName]

	attrs := NewAttr()
	attrs.entity = this
	this.Attrs = attrs

	this.I.OnInit()
}
