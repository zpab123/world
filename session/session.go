// /////////////////////////////////////////////////////////////////////////////
// session 组件

package session

import (
	"net"

	"github.com/zpab123/world/model"          // 全局模型
	"github.com/zpab123/world/network/packet" // packet 消息包
	"github.com/zpab123/world/network/socket" // socket
)

// /////////////////////////////////////////////////////////////////////////////
//

// 包变量
var (
	sessionMap map[string]SesCreateFun // sessionType->SesCreateFun 创建函数
)

// /////////////////////////////////////////////////////////////////////////////
// session 组件

// 创建 session

func NewSession(sessionType string) model.ISession {

}

// /////////////////////////////////////////////////////////////////////////////
// session 组件
type SesCreateFun func(st model.ISocket, mgr model.ISessionManage, handler model.ICilentPktHandler) model.ISession
