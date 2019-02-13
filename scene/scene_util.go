// /////////////////////////////////////////////////////////////////////////////
// scene 辅助函数

package scene

import (
	"flag"

	"github.com/zpab123/zaplog" // log 库
)

// 完成 scene 的默认设置
func defaultConfig(scene *Scene) {
	// 解析命令行参数
	parseArgs(scene)

	// 获取服务器信息
	getServerInfo(scene)

	// 设置 log 信息
	configLogger(scene)
}

// 解析 命令行参数
func parseArgs(scene *Scene) {
	// 参数定义
	//serverType := flag.String("type", "serverType", "server type") // 服务器类型，例如 gate connect area ...等
	//gid := flag.Uint("gid", 0, "gid")                                       // 服务器进程id
	name := flag.String("name", "master", "server name") // 服务器名字
	//frontend := flag.Bool("frontend", false, "is frontend server")          // 是否是前端服务器
	//host := flag.String("host", "127.0.0.1", "server host")                 // 服务器IP地址
	//port := flag.Uint("port", 0, "server port")                             // 服务器端口
	//clientHost := flag.String("clientHost", "127.0.0.1", "for client host") // 面向客户端的 IP地址
	//cTcpPort := flag.Uint("cTcpPort", 0, "tcp port")                        // tcp 连接端口
	//cWsPort := flag.Uint("cWsPort", 0, "websocket port")                    // websocket 连接端口

	// 解析参数
	flag.Parse()

	// 赋值
	//cmdParam := &cmd.CmdParam{
	//ServerType: *serverType,
	//Gid:        *gid,
	//Name: *name,
	//Frontend:   *frontend,
	//Host:       *host,
	//Port:       *port,
	//ClientHost: *clientHost,
	//CTcpPort:   *cTcpPort,
	//CWsPort:    *cWsPort,
	//}

	// 设置 app 名字
	scene.baseInfo.Name = *name
}

// 获取服务器信息
func getServerInfo(scene *Scene) {

}

// 设置 log 信息
func configLogger(scene *Scene) {

}
