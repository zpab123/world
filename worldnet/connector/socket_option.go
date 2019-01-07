// /////////////////////////////////////////////////////////////////////////////
// socket io 参数配置

package connector

import (
	"net"
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// TcpSocketOption 对象

// TcpSocketOption 配置
type TcpSocketOption struct {
	readBufferSize  int           // 读取 buffer 大小
	writeBufferSize int           // 写入 buffer 大小
	noDelay         bool          // 写入数据后，是否立即发送
	maxPacketSize   int           // 单个packet 数据包大小
	readTimeout     time.Duration // 读数据超时时间
	writeTimeout    time.Duration // 写数据超时时间
}

// 初始化 TcpSocketOption
//
// 读/写 buferr 尺寸均设置为-1
func (self *TcpSocketOption) Init() {
	self.readBufferSize = -1
	self.writeBufferSize = -1
}

// 设置 socket readBufferSize 字节数大小
func (self *TcpSocketOption) SetReadBufferSize(readBufferSize int) {
	self.readBufferSize = readBufferSize
}

// 设置 socket writeBufferSize 字节数大小
func (self *TcpSocketOption) SetWriteBufferSize(writeBufferSize int) {
	self.writeBufferSize = writeBufferSize
}

// 设置 socket noDelay
func (self *TcpSocketOption) SetNoDelay(noDelay bool) {
	self.noDelay = noDelay
}

// 设置 socket buffer 参数
func (self *TcpSocketOption) SetSocketBuffer(readBufferSize int, writeBufferSize int, noDelay bool) {
	self.SetReadBufferSize(readBufferSize)
	self.SetWriteBufferSize(writeBufferSize)
	self.SetNoDelay(noDelay)
}

// 设置 Packet 最大字节
func (self *TcpSocketOption) SetMaxPacketSize(maxSize int) {
	self.maxPacketSize = maxSize
}

// 获取 Packet 最大字节
func (self *TcpSocketOption) GetMaxPacketSize() int {
	return self.maxPacketSize
}

// 设置 net.Conn 连接对象基础参数
//
// conn 符合 net.TCPConn 接口, 才会成功
func (self *TcpSocketOption) SetSocketOption(conn net.Conn) {
	if cc, ok := conn.(*net.TCPConn); ok {
		if self.readBufferSize >= 0 {
			cc.SetReadBuffer(self.readBufferSize)
		}

		if self.writeBufferSize >= 0 {
			cc.SetWriteBuffer(self.writeBufferSize)
		}

		cc.SetNoDelay(self.noDelay)
	}
}

// 设置 net.Conn 连接对象读取超时
func (self *TcpSocketOption) SetReadTimeout(conn net.Conn, callback func()) {
	if self.readTimeout > 0 {
		// issue: http://blog.sina.com.cn/s/blog_9be3b8f10101lhiq.html
		conn.SetReadDeadline(time.Now().Add(self.readTimeout))
		conn.SetReadDeadline(time.Time{})
	}

	// 回调函数
	if nil != callback {
		callback()
	}
}

// 设置 net.Conn 连接对象写入超时
func (self *TcpSocketOption) SetWriteTimeout(conn net.Conn, callback func()) {
	if self.writeTimeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(self.writeTimeout))
		conn.SetWriteDeadline(time.Time{})
	}

	// 回调函数
	if nil != callback {
		callback()
	}
}
