// /////////////////////////////////////////////////////////////////////////////
// 二进制解码/编码

package binary

import (
	"github.com/davyxu/goobjfmt"              // 二进制编码/解码
	"github.com/zpab123/world/worldnet"       // 网络库
	"github.com/zpab123/world/worldnet/codec" // 注册中心
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 注册解码器
func init() {
	codec.RegisterCodec(new(binaryCodec))
}

// /////////////////////////////////////////////////////////////////////////////
// binaryCodec

// 二进制解码/编码
type binaryCodec struct {
}

// 将数据转换为字节数组
func (this *binaryCodec) Encode(msgObj interface{}, ctx worldnet.IContextSet) (data interface{}, err error) {
	return goobjfmt.BinaryWrite(msgObj)
}

// 将字节数组转换为数据
func (this *binaryCodec) Decode(data interface{}, msgObj interface{}) error {
	return goobjfmt.BinaryRead(data.([]byte), msgObj)
}

// 获取编码/解码器 名字
func (this *binaryCodec) Name() string {
	return codec.CODEC_TYPE_BINARY
}

// 兼容http类型
func (this *binaryCodec) MimeType() string {
	return codec.CODEC_TYPE_BINARY_MIME
}
