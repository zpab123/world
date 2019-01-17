// /////////////////////////////////////////////////////////////////////////////
// session 组件

package session

import (
	"net"

	"github.com/zpab123/world/model"          // 全局模型
	"github.com/zpab123/world/network/packet" // packet 消息包
	"github.com/zpab123/world/network/socket" // socket
	"github.com/zpab123/zplog"                // 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 包变量
var (
	sesCreateFunMap map[string]SesCreateFun // sessionType->SesCreateFun 创建函数
)

// 初始化函数
func init() {
	sesCreateFunMap[model.C_SES_TYPE_CLINET] = NewClientSession
}

// /////////////////////////////////////////////////////////////////////////////
// session 组件

// 创建 session

func NewSession(opts *model.TSessionOpts) model.ISession {
	// 查找 func
	typeName := opts.SessionType
	sesCreateFun, ok := sesCreateFunMap[typeName]
	if !ok {
		zplog.Errorf("创建 session 失败。不存在的的 session 类型=%s", typeName)
		return nil
	} else {
		return sesCreateFun(opts * model.TSessionOpts)
	}
}

// /////////////////////////////////////////////////////////////////////////////
// session 组件
type SesCreateFun func(opts *model.TSessionOpts) model.ISession
