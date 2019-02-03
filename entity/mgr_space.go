// /////////////////////////////////////////////////////////////////////////////
// space entity 管理

package entity

import (
	"reflect"
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	spaceType reflect.Type // space 实体类型
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 注册1个 space 类型实体
//
// spacePtr=符合 ISpace 的指针
func RegisterSpace(spacePtr ISpace) {
	// 反映射实体类型
	spaceVal := reflect.Indirect(reflect.ValueOf(spacePtr))
	spaceType = spaceVal.Type()

	// 注册实体
	RegisterEntity(ENTITY_TYPE_SPACE, spacePtr, false)
}
