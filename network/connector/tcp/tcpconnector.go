// /////////////////////////////////////////////////////////////////////////////
// tcp 连接器

package tcp

import (
	"net"

	"github.com/zpab123/world/consts"            // 全局常量
	"github.com/zpab123/world/network"           // 网络库
	"github.com/zpab123/world/network/connector" // 连接器
	"github.com/zpab123/world/utils"             // 工具库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

func init() {
	// 注册创建函数
	connector.RegisterCreator(newTcpConnector)
}

// /////////////////////////////////////////////////////////////////////////////
// tcpConnector 对象

// tcp 接收器
type tcpConnector struct {
}
