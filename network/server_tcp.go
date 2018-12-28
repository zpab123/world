// /////////////////////////////////////////////////////////////////////////////
// tcp 网络服务

package network

import (
	"net"
	"time"

	"github.com/zpab123/world/consts" // 全局常量
	"github.com/zpab123/world/utils"  // 工具类
	"github.com/zpab123/zplog"        // log 日志库
)

// /////////////////////////////////////////////////////////////////////////////
// TcpServer 对象

// TcpServer 对象
type TcpServer struct {
	listenAddr string      // 监听地址(格式 -> 127.0.0.1:6532)
	service    ITcpService // 符合 ITcpService 接口的对象
}

// 新建1个 TcpServer 对象
func NewTcpServer(laddr string, svc ITcpService) *TcpServer {
	// 创建对象
	ts := &TcpServer{
		listenAddr: laddr,
		service:    svc,
	}

	return ts
}

// 运行游戏服务器
func (ts *TcpServer) Run() error {
	// 错误变量
	var err error

	// 开启 tcp 服务器
	startTcpServerForever(ts.listenAddr, ts.service)

	return err
}

// 停止服务器
func (ts *TcpServer) Stop() error {
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// 私有 api

// 尝试开启 tcp 服务器（出现错误后 会重新开启）
//
// listenAddr=监听地址(格式 -> 127.0.0.1:6532);service=符合 ITcpService 接口对象
func startTcpServerForever(listenAddr string, service ITcpService) {
	for {
		err := startTcpServerOnce(listenAddr, service)
		zplog.Errorf("开启tcp服务器错误，监听地址=%s 错误=%s，%d 秒后重新开启", listenAddr, err, consts.TCP_SERVER_RECONNECT_TIME)
		time.Sleep(consts.TCP_SERVER_RECONNECT_TIME)
	}
}

// 尝试开启 tcp 服务器（出现错误后，不会重新开启）
//
// listenAddr=监听地址(格式 -> 127.0.0.1:6532);service=符合 ITcpService 接口对象
func startTcpServerOnce(listenAddr string, service ITcpService) error {
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			zplog.TraceError("tcp 服务器出现 paniced，错误原因：%s", err)
		}
	}()

	return runTcpServer(listenAddr, service)
}

// 开启 tcp 服务器
//
// listenAddr=监听地址(格式 -> 127.0.0.1:6532);service=符合 ITcpService 接口对象
func runTcpServer(listenAddr string, service ITcpService) error {
	// 开启连接监听
	ln, err := net.Listen("tcp", listenAddr)
	if nil == err {
		zplog.Infof("tcp 服务开启成功，ip=%s", listenAddr)
	} else {
		zplog.Fatalf("tcp 服务开启失败，ip=%s，错误原因=s", listenAddr, err.Error())
		return err
	}

	//  出现错误，关闭监听
	defer ln.Close()

	// 监听新连接
	for {
		newConn, err := ln.Accept()
		if err != nil {
			if utils.IsTimeoutError(err) {
				continue
			} else {
				return err
			}
		}

		// 开启新线程 处理新连接
		go service.OnTcpConn(newConn)
	}
}
