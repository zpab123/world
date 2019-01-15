// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- app 包

package model

import (
	"fmt"
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// Address 对象

// 支持地址范围的格式
type Address struct {
	Scheme  string
	Host    string
	MinPort int
	MaxPort int
	Path    string
}

// 参数检查,正确返回 true 错误 返回 fasle
func (this *Address) Check() (bool, error) {
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
func (this *Address) GetAddrRange() (string, error) {
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

// 根据 port 参数，与 Address 对象的 host 组成1个 addr 字符
//
// 返回格式： 192.168.1.1:6002
func (this *Address) HostPortString(port int) string {
	return fmt.Sprintf("%s:%d", this.Host, port)
}

// 根据 port 参数，与 Address 对象的 Scheme host Path 组成1个完整 addr 字符
//
// 返回格式： http://192.168.1.1:6002/romte
func (this *Address) String(port int) string {
	if this.Scheme == "" {
		return this.HostPortString(port)
	}

	return fmt.Sprintf("%s://%s:%d%s", this.Scheme, this.Host, port, this.Path)
}

// /////////////////////////////////////////////////////////////////////////////
// Laddr 对象

// 监听地址集合
type Laddr struct {
	TcpAddr string // Tcp 监听地址：格式 192.168.1.1:8600
	WsAddr  string // websocket 监听地址: 格式 192.168.1.1:8600
	UdpAddr string // udp 监听地址: 格式 192.168.1.1:8600
	KcpAddr string // kcp 监听地址: 格式 192.168.1.1:8600
}

// /////////////////////////////////////////////////////////////////////////////
// ConnectorOpt 对象

// connector 组件配置参数
type ConnectorOpt struct {
	TypeName  string        // connector 类型
	Heartbeat time.Duration // 心跳间隔
	Handshake func()        // 自定义的握手处理函数
	MaxConn   uint32        // 最大连接数量，超过此数值后，不再接收新连接
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *ConnectorOpt) Check() error {
	return nil
}

// /////////////////////////////////////////////////////////////////////////////
// TConnectorOpt 对象

// connector 组件配置参数
type TConnectorOpt struct {
	AcceptorType string        // 接收器 类型
	PktType      string        // packet 数据结构类型
	Heartbeat    time.Duration // 心跳间隔
	Handshake    func()        // 自定义的握手处理函数
	MaxConn      uint32        // 最大连接数量，超过此数值后，不再接收新连接
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TConnectorOpt) Check() error {
	return nil
}

// 设置默认参数
func (this *TConnectorOpt) SetDefault() {
	this.AcceptorType = C_ACCEPTOR_TYPE_COM
	this.PktType = C_PACKET_TYPE_TCP_TLV
}

// /////////////////////////////////////////////////////////////////////////////
// TConnectorOpt 对象

type TDataMgrCreator func(pm IPacketManager, handler IPacketHandler)
