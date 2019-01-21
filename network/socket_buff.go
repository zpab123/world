// /////////////////////////////////////////////////////////////////////////////
// 附带 读/写 buffer 对象的 socket

package network

import (
	"bufio"

	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
)

// /////////////////////////////////////////////////////////////////////////////
// BufferSocket 对象

// BufferSocket
type BufferSocket struct {
	model.ISocket               // 接口继承： 符合 ISocket 的对象
	bufReader     *bufio.Reader // 读取类 buffer
	bufWriter     *bufio.Writer // 写入类 buffer
}

// 创建1个新的 BufferSocket 对象
func NewBufferSocket(socket model.ISocket, opts *model.TBufferSocketOpts) *BufferSocket {
	// 参数效验
	if nil == opts {
		opts = model.NewTBufferSocketOpts()
	}

	// 创建对象
	bufsocket := &BufferSocket{
		ISocket: socket,
	}

	// 创建 buffer
	bufsocket.bufReader = bufio.NewReaderSize(socket, opts.ReadBufferSize)
	bufsocket.bufWriter = bufio.NewWriterSize(socket, opts.WriteBufferSize)

	return bufsocket
}

// 将 bufReader 中部分读取数据到 p 中
func (this *BufferSocket) Read(p []byte) (int, error) {
	return this.bufReader.Read(p)
}

// 将 p 中部分数据写入 buffer 中
func (this *BufferSocket) Write(p []byte) (int, error) {
	return this.bufWriter.Write(p)
}

// 刷新写入类 buffer 缓冲
func (this *BufferSocket) Flush() error {
	// 刷新 bufWriter
	err := this.bufWriter.Flush()
	if nil != err {
		return err
	}

	// 刷新 socket
	return this.ISocket.Flush()
}
