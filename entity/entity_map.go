// /////////////////////////////////////////////////////////////////////////////
// entityID - > Entity 类型集合 [代码完整]

package entity

import (
	"bytes"

	"github.com/zpab123/world/ids" // id 类库
)

// /////////////////////////////////////////////////////////////////////////////
// EntityMap 对象

// entityID - > *Entity 类型集合
type EntityMap map[ids.EntityID]*Entity

// 添加1个 *Entity
func (em EntityMap) Add(entity *Entity) {
	em[entity.ID] = entity
}

// 删除1个 *Entity
func (em EntityMap) Del(id ids.EntityID) {
	delete(em, id)
}

// 根据实体ID，获取 *Entity
func (em EntityMap) Get(id ids.EntityID) *Entity {
	return em[id]
}

// 获取 EntityMap 所有 key 组合而成的 切片
func (em EntityMap) Keys() (keys []ids.EntityID) {
	for eid := range em {
		keys = append(keys, eid)
	}

	return
}

// 获取 EntityMap 所有 values 组合而成的 切片
func (em EntityMap) Values() (vals []*Entity) {
	for _, e := range em {
		vals = append(vals, e)
	}

	return
}

// /////////////////////////////////////////////////////////////////////////////
// EntitySet 对象

// EntitySet is the data structure for a set of entities
type EntitySet map[*Entity]struct{}

// Add adds an entity to the EntitySet
func (es EntitySet) Add(entity *Entity) {
	es[entity] = struct{}{}
}

// Del deletes an entity from the EntitySet
func (es EntitySet) Del(entity *Entity) {
	delete(es, entity)
}

// Contains returns if the entity is in the EntitySet
func (es EntitySet) Contains(entity *Entity) bool {
	_, ok := es[entity]
	return ok
}

func (es EntitySet) ForEach(f func(e *Entity)) {
	for e := range es {
		f(e)
	}
}

func (es EntitySet) String() string {
	b := bytes.Buffer{}
	b.WriteString("{")
	first := true
	for entity := range es {
		if !first {
			b.WriteString(", ")
		} else {
			first = false
		}
		b.WriteString(entity.String())
	}
	b.WriteString("}")
	return b.String()
}
