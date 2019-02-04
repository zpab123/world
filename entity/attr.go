// /////////////////////////////////////////////////////////////////////////////
// 实体属性

package entity

// 实体属性
type Attr struct {
	entity *Entity                // 实体对象
	attrs  map[string]interface{} // 属性集合
}

// 新建1个 Attr 对象
func NewAttr() *Attr {
	a := &Attr{
		attrs: make(map[string]interface{}),
	}

	return a
}
