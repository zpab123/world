// /////////////////////////////////////////////////////////////////////////////
// codec 包常量

package codec

// /////////////////////////////////////////////////////////////////////////////
// 编码方式

// 编码方式
const (
	CODEC_TYPE_BINARY         = "binary"                // 二进制 解码/编码
	CODEC_TYPE_BINARY_MIME    = "application/binary"    // 兼容 http
	CODEC_TYPE_JSON           = "json"                  // json 解码/编码
	CODEC_TYPE_JSON_MIME      = "application/json"      // 兼容 http
	CODEC_TYPE_PROTOPLUS      = "protoplus"             // protobuff 解码/编码
	CODEC_TYPE_PROTOPLUS_MIME = "application/protoplus" // 兼容 http
)
