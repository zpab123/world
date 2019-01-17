// /////////////////////////////////////////////////////////////////////////////
// socket 基础参数管理

package network

import (
	"net"
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// TCPSocketOption 对象

// 监听地址集合
type TCPSocketOption struct {
	readBufferSize  int           // 读取 buffer 字节大小
	writeBufferSize int           // 写入 buffer 字节大小
	noDelay         bool          // 写入数据后，是否立即发送
	maxPacketSize   int           // 单个 packet 最大字节数
	readTimeout     time.Duration // 读数据超时时间
	writeTimeout    time.Duration // 写数据超时时间
}

// 初始化 TCPSocketOption
//
// 读/写 buferr 尺寸均设置为 默认参数
func (this *TCPSocketOption) Init() {
	this.readBufferSize = TCP_BUFFER_READ_SIZE   // 张鹏：原先是-1，这里被修改了
	this.writeBufferSize = TCP_BUFFER_WRITE_SIZE // 张鹏：原先是-1，这里被修改了
	this.noDelay = TCP_NO_DELAY                  // 张鹏：原先没有这个设置项，这里被修改了
}

// 设置 socket readBufferSize 字节数大小
func (this *TCPSocketOption) SetSocketReadBufferSize(readBufferSize int) {
	this.readBufferSize = readBufferSize
}

// 设置 socket writeBufferSize 字节数大小
func (this *TCPSocketOption) SetSocketWriteBufferSize(writeBufferSize int) {
	this.writeBufferSize = writeBufferSize
}

// 设置 socket noDelay
func (this *TCPSocketOption) SetSocketNoDelay(noDelay bool) {
	this.noDelay = noDelay
}

// 设置 socket buffer io 参数
func (this *TCPSocketOption) SetSocketBuffer(readBufferSize int, writeBufferSize int, noDelay bool) {
	this.SetSocketReadBufferSize(readBufferSize)
	this.SetSocketWriteBufferSize(writeBufferSize)
	this.SetSocketNoDelay(noDelay)
}

// 设置 读/写 buffer 超时时间
func (this *TCPSocketOption) SetSocketDeadline(read, write time.Duration) {
	this.readTimeout = read
	this.writeTimeout = write
}

// 设置 net.Conn 连接对象基础参数
//
// conn 符合 *net.TCPConn 接口, 才会成功
func (this *TCPSocketOption) ApplySocketOption(conn net.Conn) {
	if cc, ok := conn.(*net.TCPConn); ok {
		if this.readBufferSize >= 0 {
			cc.SetReadBuffer(this.readBufferSize)
		}

		if this.writeBufferSize >= 0 {
			cc.SetWriteBuffer(this.writeBufferSize)
		}

		cc.SetNoDelay(this.noDelay)
	}
}

// 设置 Packet 最大字节
func (this *TCPSocketOption) SetMaxPacketSize(maxSize int) {
	this.maxPacketSize = maxSize
}

// 获取 Packet 最大字节 [iSocketOpt 接口]
func (this *TCPSocketOption) GetMaxPacketSize() int {
	return this.maxPacketSize
}

// 设置 net.Conn 连接对象读取超时 [iSocketOpt 接口]
func (this *TCPSocketOption) ApplySocketReadTimeout(conn net.Conn, callback func()) {
	if this.readTimeout > 0 {
		// issue: http://blog.sina.com.cn/s/blog_9be3b8f10101lhiq.html
		conn.SetReadDeadline(time.Now().Add(this.readTimeout))
		//callback() 原来的
		conn.SetReadDeadline(time.Time{})
	}

	// 回调函数
	if nil != callback {
		callback()
	}
}

// 设置 net.Conn 连接对象写入超时 [iSocketOpt 接口]
func (this *TCPSocketOption) ApplySocketWriteTimeout(conn net.Conn, callback func()) {
	if this.writeTimeout > 0 {
		conn.SetWriteDeadline(time.Now().Add(this.writeTimeout))
		//callback() 原来的
		conn.SetWriteDeadline(time.Time{})
	}

	// 回调函数
	if nil != callback {
		callback()
	}
}
