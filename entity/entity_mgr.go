// /////////////////////////////////////////////////////////////////////////////
// entity 管理

package entity

import (
	"reflect"

	"github.com/zpab123/world/ids" // id 库
	"github.com/zpab123/zaplog"    // log 库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

var (
	registeredEntityTypes = map[string]*EntityTypeDesc{} // 实体 type->*EntityTypeDesc 类型集合
	entityManager         = newEntityManager()           // _EntityManager 实例
)

// /////////////////////////////////////////////////////////////////////////////
// 对外 api

// 根据 实体类型 注册实体，返回实体描述对象指针 *EntityTypeDesc
//
// typeName=实体类型；entity=符合 IEntity 接口的指针；isService=是否是服务类实体
func RegisterEntity(typeName string, entity IEntity, isService bool) *EntityTypeDesc {
	// 已经注册警告
	if _, ok := registeredEntityTypes[typeName]; ok {
		zaplog.Fatalf("实体类型=%s，已经注册过", typeName)
	}

	// 类型推断
	entityVal := reflect.ValueOf(entity)
	entityType := entityVal.Type()
	if entityType.Kind() == reflect.Ptr {
		entityType = entityType.Elem()
	}

	// 创建描述文件
	rpcDescs := rpcDescMap{} // 实体 rpc 方法集合
	entityTypeDesc := &EntityTypeDesc{
		isService: isService,
		useAOI:    false,
		rpcDescs:  rpcDescs,
	}

	// 保存描述文件
	registeredEntityTypes[typeName] = entityTypeDesc

	// 实体方法
	entityPtrType := reflect.PtrTo(entityType)
	numMethods := entityPtrType.NumMethod()
	for i := 0; i < numMethods; i++ {
		method := entityPtrType.Method(i)
		rpcDescs.visit(method)
	}

	// 返回注册信息
	zaplog.Infof(">>> 注册实体 %s => %s <<<", typeName, entityType.Name())
	// entity.DescribeEntityType(entityTypeDesc)
	return entityTypeDesc
}

// /////////////////////////////////////////////////////////////////////////////
// 私有 api

func createEntity(typeName string, entityID ids.EntityID, space *Space, pos Vector3) *Entity {
	// 注册效验
	entityTypeDesc, ok := registeredEntityTypes[typeName]
	if !ok {
		zaplog.Panicf("实体创建失败：类型未注册。typeName=%s", typeName)
	}

	// ID 效验
	if entityID == "" {
		entityID = ids.GenEntityID()
	}

	// 创建实体
	var entity *Entity
	var entityInstance reflect.Value

	entityInstance = reflect.New(entityTypeDesc.entityType)
	entity = reflect.Indirect(entityInstance).FieldByName("Entity").Addr().Interface().(*Entity)
	entity.init(typeName, entityID, entityInstance)
	entity.Space = nilSpace

	entityManager.put(entity)

	// 进入空间
	if space != nil {
		space.enter(entity, pos, false)
	}

	return entity
}

// /////////////////////////////////////////////////////////////////////////////
// _EntityManager 对象

// 实体管理对象
type _EntityManager struct {
	entities EntityMap            // entityID - > *Entity 类型集合
	typeMap  map[string]EntityMap // type -> EntityMap 类型集合
}

// 创建1个新的 _EntityManager 对象
func newEntityManager() *_EntityManager {
	// 创建对象
	eMgr := &_EntityManager{
		entities: EntityMap{},
		typeMap:  map[string]EntityMap{},
	}

	return eMgr
}

// 添加1个 Entity
func (em *_EntityManager) put(e *Entity) {
	// 添加id 集合
	em.entities.Add(e)

	// 添加类型集合
	etype := e.TypeName
	eid := e.ID
	if entitys, ok := em.typeMap[etype]; ok {
		entitys.Add(e)
	} else {
		em.typeMap[etype] = EntityMap{eid: entity}
	}
}

// 删除1个 Entity
func (em *_EntityManager) del(e *Entity) {
	// 删除 id 集合
	eid := e.ID
	em.entities.Del(eid)

	// 删除类型集合
	if entitys, ok := em.typeMap[e.TypeName]; ok {
		entitys.Del(eid)
	}
}

// 根据ID 获取1个 Entity
func (em *_EntityManager) get(id ids.EntityID) *Entity {
	return em.entities.Get(id)
}

// /////////////////////////////////////////////////////////////////////////////
// EntityTypeDesc 对象

// 实体描述对象
type EntityTypeDesc struct {
	isService  bool         // 是否是服务类型实体
	useAOI     bool         // 是否使用 AOI
	rpcDescs   rpcDescMap   // 实体 rpc 方法集合
	entityType reflect.Type // 反射
}
