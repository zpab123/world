// /////////////////////////////////////////////////////////////////////////////
// StringSet

package collection

// /////////////////////////////////////////////////////////////////////////////
// StringSet

// 字符串容器
type StringSet map[string]struct{}

// 检查是否含有 elem 属性
func (ss StringSet) Contains(elem string) bool {
	_, ok := ss[elem]

	return ok
}

// 添加1个 elem 属性
func (ss StringSet) Add(elem string) {
	ss[elem] = struct{}{}
}

// 移除1个 elem 属性
func (ss StringSet) Remove(elem string) {

	delete(ss, elem)
}

// 将所有属性转化成 []string
func (ss StringSet) ToList() []string {
	keys := make([]string, 0, len(ss))

	for s := range ss {
		keys = append(keys, s)
	}

	return keys
}
