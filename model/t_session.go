// /////////////////////////////////////////////////////////////////////////////
// 全局基础 struct -- session 包

package model

import (
	"fmt"
	"time"
)

// /////////////////////////////////////////////////////////////////////////////
// TSessionOpts 对象

// session 配置参数
type TSessionOpts struct {
	TacketType string        // packet 数据结构类型
	Heartbeat  time.Duration // 心跳间隔
	Handshake  func()        // 自定义的握手处理函数
}

// 创建1个新的 TSessionOpts
func NewTSessionOpts() *TSessionOpts {
	// 创建对象
	opts := &TSessionOpts{}
	opts.SetDefaultOpts()

	return opts
}

// 设置默认参数
func (this *TSessionOpts) SetDefaultOpts() {
	this.TacketType = C_PACKET_TYPE_TCP_TLV
}

// 检查 ConnectorConfig 参数是否存在错误
func (this *TSessionOpts) Check() error {
	return nil
}
