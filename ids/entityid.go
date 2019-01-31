// /////////////////////////////////////////////////////////////////////////////
// entityID 对象 [代码齐全]

package ids

import (
	"github.com/zpab123/zaplog" // log 工具
)

// /////////////////////////////////////////////////////////////////////////////
// 包变量

// 实体ID 长度
const ENTITYID_LENGTH = UUID_LENGTH

// /////////////////////////////////////////////////////////////////////////////
// entityID 对象

// EntityID 对象
type EntityID string

// 生成1个实体ID
func GenEntityID() EntityID {
	return EntityID(GenUUID())
}

// 实体ID 是否 == ""
func (id EntityID) IsNil() bool {
	return id == ""
}

// 将1个 字符串id 转化为 EntityID
func MustEntityID(id string) EntityID {
	if len(id) != ENTITYID_LENGTH {
		zaplog.Panicf("%s 的长度=%d，是1个无效的实体长度(正确长度=%d)", id, len(id), ENTITYID_LENGTH)
	}

	return EntityID(id)
}
