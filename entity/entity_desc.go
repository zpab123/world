// /////////////////////////////////////////////////////////////////////////////
// 实体注册对象

package entity

import (
	"reflect"
	"strings"

	"github.com/zpab123/world/collection" // 容器库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	_VALID_ATTR_DEFS = collection.StringSet{} // 有效属性类型
)

func init() {
	_VALID_ATTR_DEFS.Add(strings.ToLower(ATTR_DEFS_CLIENT))
	_VALID_ATTR_DEFS.Add(strings.ToLower(ATTR_DEFS_ALL))
	_VALID_ATTR_DEFS.Add(strings.ToLower(ATTR_DEFS_PER))
}

// /////////////////////////////////////////////////////////////////////////////
// EntityTypeDesc 对象

// 实体描述对象
type EntityTypeDesc struct {
	isService    bool         // 是否是服务类型实体
	isPersistent bool         // 是否持久化
	useAOI       bool         // 是否使用 AOI
	aoiDistance  Coord        // aoi 距离
	entityType   reflect.Type // 实体类型
	rpcDescs     rpcDescMap   // 实体 rpc 方法集合
}

// 设置是否持久化
func (this *EntityTypeDesc) SetPersistent(v bool) {
	this.isPersistent = v
}

// 设置是否使用 AOI
func (this *EntityTypeDesc) SetUseAOI(v bool) {
	this.useAOI = v
}

// 定义属性
func (this *EntityTypeDesc) DefineAttr(attr string, defs ...string) {

}
