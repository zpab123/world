// /////////////////////////////////////////////////////////////////////////////
// 编码/解码器 注册中心

package codec

import (
	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 变量
var (
	registedCodecs []worldnet.ICodec // 编码/解码器 切片
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 注册编码器
func RegisterCodec(c worldnet.ICodec) {
	// 重复注册
	if GetCodec(c.Name()) != nil {
		panic("Panic：Codec 重复注册，name=" + c.Name())
	}

	// 添加到切片
	registedCodecs = append(registedCodecs, c)
}

// 获取编码器
//
// 返回：nil=编码/解码器 不存在
func GetCodec(name string) worldnet.ICodec {
	for _, c := range registedCodecs {
		if c.Name() == name {
			return c
		}
	}

	return nil
}

// 指定编码器不存在时，报错
func MustGetCodec(name string) worldnet.ICodec {
	// 获取 codec
	codec := GetCodec(name)

	// 异常
	if nil == codec {
		panic("Panic：Codec 未注册，name=" + name)
	}

	return codec
}
