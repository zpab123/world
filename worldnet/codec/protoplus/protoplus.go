// /////////////////////////////////////////////////////////////////////////////
// protobuff 编码/解码

package protoplus

import (
	"github.com/davyxu/protoplus/proto"       // protobuff 编码/解码
	"github.com/zpab123/world/worldnet"       // 网络库
	"github.com/zpab123/world/worldnet/codec" // 注册中心
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 注册解码器
func init() {
	codec.RegisterCodec(new(protoplus))
}

// /////////////////////////////////////////////////////////////////////////////
// protoplus

type protoplus struct {
}

// 将结构体编码为 JSON 的字节数组
func (this *protoplus) Encode(msgObj interface{}, ctx worldnet.IContextSet) (data interface{}, err error) {
	return proto.Marshal(msgObj)
}

// 将 JSON 的字节数组解码为结构体
func (this *protoplus) Decode(data interface{}, msgObj interface{}) error {
	return proto.Unmarshal(data.([]byte), msgObj)
}

// 获取编码/解码器 名字
func (this *protoplus) Name() string {
	return codec.CODEC_TYPE_PROTOPLUS
}

// 兼容http类型
func (this *protoplus) MimeType() string {
	return codec.CODEC_TYPE_PROTOPLUS_MIME
}
