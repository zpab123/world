// /////////////////////////////////////////////////////////////////////////////
// space 类实体

package entity

// /////////////////////////////////////////////////////////////////////////////
// space 对象

// space 类实体
type Space struct {
	Entity // 继承 Entity 对象
}

// 当 space 初始化的时候调用
func (sp *Space) OnSpaceInit() {

}

// 当 space 被创建的时候调用
func (sp *Space) OnSpaceCreated() {

}

// 当 space 被销毁的时候调用
func (sp *Space) OnSpaceDestroy() {

}

// 当某个 entity 进入的时候调用
func (sp *Space) OnEntityEnter() {

}

// 当某个 entity 离开的时候调用
func (sp *Space) OnEntityLeave() {

}

// 仅仅被 nil space 调用
func (sp *Space) OnGameReady() {

}
