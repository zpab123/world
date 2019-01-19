// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- app 包

package model

import (
	"fmt"
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// TAddress 对象

// 支持地址范围的格式
type TAddress struct {
	Scheme  string
	Host    string
	MinPort int
	MaxPort int
	Path    string
}

// 参数检查,正确返回 true 错误 返回 fasle
func (this *TAddress) Check() (bool, error) {
	// 地址效验 -- 正则是否是ip 地址

	// 端口效验

	// 最大端口效验
	if this.MaxPort < this.MinPort {
		this.MaxPort = this.MinPort
	}

	return true, nil
}

// 获取带范围的 addr 格式
//
// 例如：scheme://host:minPort~maxPort/path
func (this *TAddress) GetAddrRange() (string, error) {
	// 参数检查
	ok, err := this.Check()
	if !ok {
		return "", err
	}

	// 获取地址
	var addr string
	if "" == this.Scheme {
		addr = fmt.Sprintf("%s:%d~%d", this.Host, this.MinPort, this.MaxPort)
	} else {
		addr = fmt.Sprintf("%s://%s:%d~%d/%s", this.Scheme, this.Host, this.MinPort, this.MaxPort, this.Path)
	}

	return addr, nil
}

// 根据 port 参数，与 TAddress 对象的 host 组成1个 addr 字符
//
// 返回格式： 192.168.1.1:6002
func (this *TAddress) HostPortString(port int) string {
	return fmt.Sprintf("%s:%d", this.Host, port)
}

// 根据 port 参数，与 TAddress 对象的 Scheme host Path 组成1个完整 addr 字符
//
// 返回格式： http://192.168.1.1:6002/romte
func (this *TAddress) String(port int) string {
	if this.Scheme == "" {
		return this.HostPortString(port)
	}

	return fmt.Sprintf("%s://%s:%d%s", this.Scheme, this.Host, port, this.Path)
}

// /////////////////////////////////////////////////////////////////////////////
// Laddr 对象

// 监听地址集合
type TLaddr struct {
	TcpAddr string // Tcp 监听地址：格式 192.168.1.1:8600
	WsAddr  string // websocket 监听地址: 格式 192.168.1.1:8600
	UdpAddr string // udp 监听地址: 格式 192.168.1.1:8600
	KcpAddr string // kcp 监听地址: 格式 192.168.1.1:8600
}

// /////////////////////////////////////////////////////////////////////////////
// TTcpConnOpts 对象

// TcpSocket 配置参数
type TTcpConnOpts struct {
	ReadBufferSize  int           // 读取 buffer 字节大小
	WriteBufferSize int           // 写入 buffer 字节大小
	NoDelay         bool          // 写入数据后，是否立即发送
	MaxPacketSize   int           // 单个 packet 最大字节数
	ReadTimeout     time.Duration // 读数据超时时间
	WriteTimeout    time.Duration // 写数据超时时间
}

// 创建1个新的 TTcpConnOpts 对象
func NewTTcpConnOpts() *TTcpConnOpts {
	// 创建对象
	tcpOpts := &TTcpConnOpts{}
	tcpOpts.SetDefaultOpts()

	return tcpOpts
}

// 设置默认参数
func (this *TTcpConnOpts) SetDefaultOpts() {
	this.ReadBufferSize = C_TCP_BUFFER_READ_SIZE   // 张鹏：原先是-1，这里被修改了
	this.WriteBufferSize = C_TCP_BUFFER_WRITE_SIZE // 张鹏：原先是-1，这里被修改了
	this.NoDelay = C_TCP_NO_DELAY                  // 张鹏：原先没有这个设置项，这里被修改了
}
