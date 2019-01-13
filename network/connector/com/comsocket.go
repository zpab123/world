// /////////////////////////////////////////////////////////////////////////////
// 能够接收/发送 packet 的 socket

package com

import (
	"net"

	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
)

// /////////////////////////////////////////////////////////////////////////////
// comSocket 对象

// comSocket 对象
type comSocket struct {
	net.Conn // 接口继承： 符合 et.Conn 接口的对象
}

// 创建1个新的 comSocket 对象
func newComSocket(conn net.Conn, isWebSocket bool) model.ISocket {
	// 创建 socket
	socket := &comSocket{
		Conn: conn,
	}

	return socket
}

// 刷新缓冲区 [ISocket 接口]
func (this *comSocket) Flush() error {
	return nil
}
