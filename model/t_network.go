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
// 网络接收器配置参数

// 监听地址集合
type TAcceptorOpts struct {
	SessionType string // 创建 session 类型
}

// 新建1个 TAcceptorOpts 对象
func NewTAcceptorOpts() *TAcceptorOpts {
	// 创建对象
	aptor := &TAcceptorOpts{}
	aptor.SetDefaultOpts()

	return aptor
}

// 设置默认参数
func (this *TAcceptorOpts) SetDefaultOpts() {
	this.SessionType = C_SES_TYPE_CLINET
}

// /////////////////////////////////////////////////////////////////////////////
// TConnectorOpt 对象

// connector 组件配置参数
type TConnectorOpt struct {
	SessionOpts  *TSessionOpts // 对象继承： session 设置
	AcceptorType string        // 接收器 类型
	MaxConn      uint32        // 最大连接数量，超过此数值后，不再接收新连接
}

// 创建1个新的 TConnectorOpt
func NewTConnectorOpt() *TConnectorOpt {
	// 创建对象
	opts := &TConnectorOpt{
		SessionOpts: NewTSessionOpts(),
	}

	opts.SetDefaultOpts()

	return opts
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TConnectorOpt) Check() error {
	return nil
}

// 设置默认参数
func (this *TConnectorOpt) SetDefaultOpts() {
	this.AcceptorType = C_ACCEPTOR_TYPE_COM
	this.SessionOpts.SetDefaultOpts()
}

// /////////////////////////////////////////////////////////////////////////////
// TConnectorOpt 对象

type TDataMgrCreator func(pm IPacketManager, handler IPacketHandler)
