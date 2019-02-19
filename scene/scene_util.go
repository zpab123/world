// /////////////////////////////////////////////////////////////////////////////
// scene 辅助函数

package scene

import (
	"flag"
	"fmt"
	"os"

	"github.com/zpab123/world/config" // 配置工具库
	//"github.com/zpab123/world/dispatcher" // 消息分发库
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
	// 获取服务器类型
	serverType := scene.baseInfo.Type

	// 获取服务器名字
	name := scene.baseInfo.Name

	// 获取类型列表
	list := config.GetServerMap()[serverType]
	if nil == list || len(list) <= 0 {
		zaplog.Error("Scene 获取 serverType 信息失败。 serverType=%s", serverType)

		os.Exit(1)
	}

	// 获取服务器信息
	for _, info := range list {
		if info.Name == name {
			scene.serverInfo = info
			break
		}
	}

	if scene.serverInfo == nil {
		zaplog.Errorf("Scene 获取 server.json 信息失败。 serverName=%s", scene.baseInfo.Name)

		os.Exit(1)
	}
}

// 设置 log 信息
func configLogger(scene *Scene) {
	// 模块名字
	zaplog.SetSource(scene.baseInfo.Name)

	// 输出等级
	lv := config.GetWorldConfig().LogLevel
	zaplog.SetLevel(zaplog.ParseLevel(lv))

	// 输出文件
	logFile := fmt.Sprintf("%s.log", scene.baseInfo.Name)
	var outputs []string
	stdErr := config.GetWorldConfig().LogStderr
	if stdErr {
		outputs = append(outputs, "stderr")
	}
	outputs = append(outputs, logFile)
	zaplog.SetOutput(outputs)
}

// 创建默认组件
func createComponent(scene *Scene) {
	// 创建分发客户端

}

// 创建分发服务器
func newDispatcherClient(scene *Scene) {
	// 配置参数
	//opt := scene.componentMgr.

	//dc := dispatcher.NewDispatcherClient()
}
