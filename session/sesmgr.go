// /////////////////////////////////////////////////////////////////////////////
// Session 管理对象 [代码完整]

package session

import (
	"sync"
	"sync/atomic"

	"github.com/zpab123/syncutil"      // 同步变量
	"github.com/zpab123/world/model"   // 全局 [常量-基础数据类型-接口] 集合
	"github.com/zpab123/world/network" // 网路库
)

// /////////////////////////////////////////////////////////////////////////////
// SessionManager

// Session 管理对象
type SessionManager struct {
	sesMap   sync.Map             // id -> session 对象集合
	sesIDGen syncutil.AtomicInt64 // session ID生成器
	count    syncutil.AtomicInt64 // 记录当前在使用的会话数量
}

// 添加1个符合 ISession 接口的对象
//
// [ISessionManager 接口]
func (self *SessionManager) Add(ses model.ISession) {
	// id +1
	id := self.sesIDGen.Add(1)

	// 计数 +1
	self.count.Add(1)

	// 设置 session ID
	//ses.(interface {
	//SetID(int64)
	//}).SetID(id)
	ses.SetId(id)

	// 保存
	self.sesMap.Store(id, ses)
}

// 移除1个符合 ISession 接口的对象
//
// [ISessionManager 接口]
func (self *SessionManager) Remove(ses model.ISession) {
	// 移除
	self.sesMap.Delete(ses.GetId())

	// 计数 -1
	self.count.Add(-1)
}

// 获取当前 ISession 数量 [ISessionManager 接口]
func (self *SessionManager) GetCount() int {
	return int(self.count.Load())
}

// 设置ID开始的号 [ISessionManager 接口]
func (self *SessionManager) SetIDStart(start int64) {
	self.sesIDGen.Store(start)
}

// 从 session 存取器中获取一个连接 [ISessionAccessor 接口]
//
// 返回 nil=不存在
func (self *SessionManager) GetSession(id int64) model.ISession {
	// 遍历查找
	if ses, ok := self.sesMap.Load(id); ok {
		return ses.(model.ISession)
	}

	return nil
}

// 遍历连接 [ISessionAccessor 接口]
func (self *SessionManager) VisitSession(callback func(model.ISession) bool) {
	self.sesMap.Range(func(key, value interface{}) bool {
		return callback(value.(model.ISession))
	})
}

// 活跃的会话数量 [ISessionAccessor 接口]
func (self *SessionManager) SessionCount() int {
	v := self.count.Load()

	return int(v)
}

// 关闭所有连接 [ISessionAccessor 接口]
func (self *SessionManager) CloseAllSession() {
	// 处理函数
	f := func(ses model.ISession) bool {
		ses.Close()

		return true
	}

	// 遍历
	self.VisitSession(f)
}