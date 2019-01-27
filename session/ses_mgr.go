// /////////////////////////////////////////////////////////////////////////////
// Session 管理对象 [代码完整]

package session

import (
	"sync"
	"sync/atomic"

	"github.com/zpab123/syncutil"      // 同步变量
	"github.com/zpab123/world/model"   // 全局模型
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

// 收到1个新的 session [ISessionManager 接口]
func (this *SessionManager) OnNewSession(ses model.ISession) {
	this.Add(ses)
}

// 某个 session 关闭 [ISessionManager 接口]
func (this *SessionManager) OnSessionClose(ses model.ISession) {
	this.Remove(this)
}

// 添加1个符合 ISession 接口的对象
func (this *SessionManager) Add(ses model.ISession) {
	// id +1
	id := this.sesIDGen.Add(1)

	// 计数 +1
	this.count.Add(1)

	// 设置 session ID
	ses.SetId(id)

	// 保存
	this.sesMap.Store(id, ses)
}

// 移除1个符合 ISession 接口的对象
func (this *SessionManager) Remove(ses model.ISession) {
	// 移除
	this.sesMap.Delete(ses.GetId())

	// 计数 -1
	this.count.Add(-1)
}

// 获取当前 ISession 数量
func (this *SessionManager) GetCount() int {
	return int(this.count.Load())
}

// 设置ID开始的号
func (this *SessionManager) SetIDStart(start int64) {
	this.sesIDGen.Store(start)
}

// 从 session 存取器中获取一个连接
//
// 返回 nil=不存在
func (this *SessionManager) GetSession(id int64) model.ISession {
	// 遍历查找
	if ses, ok := this.sesMap.Load(id); ok {
		return ses.(model.ISession)
	}

	return nil
}

// 遍历连接
func (this *SessionManager) VisitSession(callback func(model.ISession) bool) {
	this.sesMap.Range(func(key, value interface{}) bool {
		return callback(value.(model.ISession))
	})
}

// 活跃的会话数量
func (this *SessionManager) SessionCount() int {
	v := this.count.Load()

	return int(v)
}

// 关闭所有连接
func (this *SessionManager) CloseAllSession() {
	// 处理函数
	f := func(ses model.ISession) bool {
		ses.Close()

		return true
	}

	// 遍历
	this.VisitSession(f)
}
