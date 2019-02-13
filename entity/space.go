// /////////////////////////////////////////////////////////////////////////////
// space 类实体

package entity

import (
	"github.com/zpab123/aoi" // aoi
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	nilSpace *Space // 空 Space
)

// /////////////////////////////////////////////////////////////////////////////
// space 对象

// space 类实体
type Space struct {
	Entity                 // 继承 Entity 对象
	Kind   int             // space 种类
	Self   ISpace          // 未知
	aoiMgr aoi.IAoiManager // aoi 管理
}

// 当 space 初始化的时候调用
func (this *Space) OnSpaceInit() {

}

// 当 space 被创建的时候调用
func (this *Space) OnSpaceCreated() {

}

// 当 space 被销毁的时候调用
func (this *Space) OnSpaceDestroy() {

}

// 当某个 entity 进入的时候调用
func (this *Space) OnEntityEnter() {

}

// 当某个 entity 离开的时候调用
func (this *Space) OnEntityLeave() {

}

// 仅仅被 nil space 调用
func (this *Space) OnGameReady() {

}

// 某个实体进入
func (this *Space) enter(entity *Entity, pos Vector3, isRestore bool) {

}
