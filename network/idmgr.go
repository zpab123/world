// /////////////////////////////////////////////////////////////////////////////
// id 管理

package network

// /////////////////////////////////////////////////////////////////////////////
// IdMananger 对象

// socket id 管理
type IdMananger struct {
	id int64 // socket id 标识
}

// 获取ID
func (this *IdMananger) GetId() int64 {
	return this.id
}

// 设置ID
func (this *IdMananger) SetId(v int64) {
	this.id = v
}