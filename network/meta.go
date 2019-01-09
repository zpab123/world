// /////////////////////////////////////////////////////////////////////////////
// 消息元信息

package network

import (
	"fmt"
	"path"
	"reflect"
	"regexp"
	"strings"
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 变量
var (
	metaMapById       = map[int]*MessageMeta{}          // id -> *MessageMeta 映射集合
	metaMapByType     = map[reflect.Type]*MessageMeta{} // reflect.Type -> *MessageMeta 映射集合
	metaMapByFullName = map[string]*MessageMeta{}       // fullname -> *MessageMeta 映射集合
)

// /////////////////////////////////////////////////////////////////////////////
// public api

/*
http消息
Method URL -> Meta
Type -> Meta

非http消息
ID -> Meta
Type -> Meta

*/

// 注册消息元信息
//
// 非http类,才需要包装,Type必须唯一
func RegisterMessageMeta(meta *MessageMeta) *MessageMeta {
	// 根据类型保存
	if _, ok := metaMapByType[meta.Type]; ok {
		panic(fmt.Sprintf("注册 MessageMeta 消息出错：类型重复，id=%d，name=%s", meta.ID, meta.Type.Name()))
	} else {
		metaMapByType[meta.Type] = meta
	}

	// 根据名字保存
	if _, ok := metaMapByFullName[meta.FullName()]; ok {
		panic(fmt.Sprintf("注册 MessageMeta 消息出错：全名重复，id=%d，name=%s", meta.ID, meta.FullName()))
	} else {
		metaMapByFullName[meta.FullName()] = meta
	}

	// ID 效验
	if 0 == meta.ID {
		panic(fmt.Sprintf("注册 MessageMeta 消息出错：ID=0，type=%s", meta.TypeName()))
	}

	// 根据ID保存
	if prev, ok := metaMapById[meta.ID]; ok {
		panic(fmt.Sprintf("注册 MessageMeta 消息出错：ID重复，id=%d，meta_type=%s，指针type=", meta.ID, meta.TypeName(), prev.TypeName()))
	} else {
		metaMapById[meta.ID] = meta
	}

	return meta
}

// 根据名字查找消息元信息
func GetMetaByFullName(name string) *MessageMeta {
	if v, ok := metaMapByFullName[name]; ok {
		return v
	}

	return nil
}

// 根据类型查找消息元信息
func GetMetaByType(t reflect.Type) *MessageMeta {
	// 类型效验
	if t == nil {
		return nil
	}

	// 获取指针类型指向值的类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 查找
	if v, ok := metaMapByType[t]; ok {
		return v
	}

	return nil
}

// 根据id查找消息元信息
func GetMetaByID(id int) *MessageMeta {
	if v, ok := metaMapById[id]; ok {
		return v
	}

	return nil
}

// 根据消息对象获得消息元信息
func GetMetaByMsg(msg interface{}) *MessageMeta {
	if msg == nil {
		return nil
	}

	return GetMetaByType(reflect.TypeOf(msg))
}

// 根据名字遍历 Meta
func MessageMetaVisit(nameRule string, callback func(meta *MessageMeta) bool) error {
	// 创建正则
	exp, err := regexp.Compile(nameRule)
	if err != nil {
		return err
	}

	// 遍历
	for name, meta := range metaMapByFullName {
		if exp.MatchString(name) {
			if !callback(meta) {
				return nil
			}
		}
	}

	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// MessageMeta 对象

// 消息元信息
type MessageMeta struct {
	Codec   ICodec       // 消息用到的编码/解码格式 （符合 ICodec 类型接口）
	Type    reflect.Type // 消息类型
	ID      int          // 消息ID (二进制协议中使用，每个 MessageMeta 的id都不能重复)
	ctxList []*context   // 上下文指针切片
}

// 获取 全名
func (this *MessageMeta) FullName() string {
	// 空
	if nil == this {
		return ""
	}

	// 获取元素类型
	rtype := this.Type
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}

	// 获取包路径名？
	var sb strings.Builder
	sb.WriteString(path.Base(rtype.PkgPath()))
	sb.WriteString(".")
	sb.WriteString(rtype.Name())

	return sb.String()
}

// 获取类型名字
func (this *MessageMeta) TypeName() string {
	if nil == this {
		return ""
	}

	// 指针类型
	if this.Type.Kind() == reflect.Ptr {
		return this.Type.Elem().Name()
	}

	return this.Type.Name()
}

// 创建 meta 类型的实例
func (this *MessageMeta) NewType() interface{} {
	if nil == this.Type {
		return nil
	}

	return reflect.New(this.Type).Interface()
}

// /////////////////////////////////////////////////////////////////////////////
// context 对象

// 自定义 context
type context struct {
	name string      // 名字
	data interface{} // 数据
}
