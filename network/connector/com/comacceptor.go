// /////////////////////////////////////////////////////////////////////////////
// 组合模式，支持 tcp websocket 连接

package com

import (
	"net"
	"strings"

	"github.com/zpab123/world/ifs"               // 接口库
	"github.com/zpab123/world/model"             // 常用数据类型
	"github.com/zpab123/world/network/connector" // 连接器
	"github.com/zpab123/world/utils"             // 工具库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

func init() {
	// 注册创建函数
	connector.RegisterAcceptor(connector.CONNECTOR_TYPE_COM, newComAcceptor)
}

// /////////////////////////////////////////////////////////////////////////////
// comAcceptor 对象

// 支持 tcp websocket 连接
type comAcceptor struct {
	connector.AddrManager                    // 对象继承： 监听地址管理
	connector.TCPSocketOption                // 对象继承： socket 基础参数管理
	connector                 ifs.IConnector // connector 对象
	tcpListener               net.Listener   // tcp 侦听器
}

// 创建1个 comAcceptor 对象
func newComAcceptor(cntor ifs.IConnector) ifs.IAcceptor {
	// 创建 comAcceptor
	comaptor := &comAcceptor{
		connector: cntor,
	}

	return comaptor
}

// 启动 Acceptor [IAcceptor 接口]
func (this *comAcceptor) Run() {
	// 启动 tcp 侦听
	this.runTcpListener()

	// 启动 websocket 侦听
}

// 停止 Acceptor [IAcceptor 接口]
func (this *comAcceptor) Stop() {

}

// 获取 侦听成功的 tcp 端口
func (this *comAcceptor) GetTcpPort() int {
	if nil == this.tcpListener {
		return 0
	}

	port := this.tcpListener.Addr().(net.TCPAddr).Port

	return port
}

// 获取 侦听成功的 tcpAddr 端口
func (this *comAcceptor) GetTcpAddr() string {
	// 获取 host
	pos := strings.Index(this.GetAddr().TcpAddr, ":")
	if pos == -1 {
		return this.GetAddr().TcpAddr
	}
	host := this.GetAddr().TcpAddr[:pos]

	// 获取 port
	port := this.GetTcpPort()

	return utils.JoinAddress(host, port)
}

// 启动 tcp 侦听
func (this *comAcceptor) runTcpListener() {
	// 创建侦听器
	f := func(addr *model.Address, port int) (interface{}, error) {
		return net.Listen("tcp", addr.HostPortString(port))
	}
	ln, err := utils.DetectPort(this.GetAddr().TcpAddr, f)

	// 创建失败
	if nil != err {
		zplog.Errorf("启动 tcpAcceptor 失败。错误=%v", err.Error())

		return
	}

	// 创建成功
	this.listener = ln.(net.Listener)
	zplog.Infof("comAcceptor-tcp 启动成功。监听地址=%s", this.GetTcpAddr())

	// 侦听 tcp 连接
	go this.acceptTcpConn()
}

// 接收新的 tcp 连接
func (this *comAcceptor) acceptTcpConn() {
	//  出现错误，关闭监听
	defer this.tcpListener.Close()

	// 监听新连接
	for {
		newConn, err := this.tcpListener.Accept()
		if nil != err {
			if utils.IsTimeoutError(err) {
				continue
			} else {
				return err
			}
		}

		// 开启新线程 处理新 tcp 连接
		go this.onNewTcpConn(newConn)
	}
}

// 收到1个新的 tcp 连接
func (this *comAcceptor) onNewTcpConn(conn net.Conn) {
	// 配置 io 参数
	this.ApplySocketOption(conn)

	// 创建 socket 对象
	newComSocket(conn, false)

	// 通知 Connector 组件
	//this.connector.OnNewSocket(socket)
}

// /////////////////////////////////////////////////////////////////////////////
// private api
