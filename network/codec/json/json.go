// /////////////////////////////////////////////////////////////////////////////
// json 方式 编码/解码

package json

import (
	"encoding/json"

	"github.com/zpab123/world/worldnet"       // 网络库
	"github.com/zpab123/world/worldnet/codec" // 注册中心
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

// 注册解码器
func init() {
	codec.RegisterCodec(new(jsonCodec))
}

// /////////////////////////////////////////////////////////////////////////////
// jsonCodec

// json 方式 编码/解码
type jsonCodec struct {
}

// 将结构体编码为 JSON 的字节数组
func (this *jsonCodec) Encode(msgObj interface{}, ctx worldnet.IContextSet) (data interface{}, err error) {
	return json.Marshal(msgObj)
}

// 将 JSON 的字节数组解码为结构体
func (this *jsonCodec) Decode(data interface{}, msgObj interface{}) error {
	return json.Unmarshal(data.([]byte), msgObj)
}

// 获取编码/解码器 名字
func (this *jsonCodec) Name() string {
	return codec.CODEC_TYPE_JSON
}

// 兼容http类型
func (this *jsonCodec) MimeType() string {
	return codec.CODEC_TYPE_JSON_MIME
}
