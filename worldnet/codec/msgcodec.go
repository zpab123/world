// /////////////////////////////////////////////////////////////////////////////
// 从 pkt 中解码 Message

package codec

import (
	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 编码消息, 在使用了带内存池的 codec 中，可以传入 session 或 connector 的 ContextSet，保存内存池上下文，默认ctx传nil
func EncodeMessage(msg interface{}, ctx worldnet.IContextSet) (data []byte, meta *worldnet.MessageMeta, err error) {
	// 获取消息元信息
	meta = worldnet.GetMetaByMsg(msg)

	// 消息没有注册
	if nil == meta {
		return nil, nil, worldnet.NewErrorContext("该 msg 不存在 MessageMeta", msg)
	}

	// 将消息编码为字节数组
	var raw interface{}
	raw, err = meta.Codec.Encode(msg, ctx)
	if nil != err {
		return
	}
	data = raw.([]byte)

	return
}

// 解码消息
func DecodeMessage(msgid int, data []byte) (interface{}, *worldnet.MessageMeta, error) {
	// 获取消息元信息
	meta := worldnet.GetMetaByID(msgid)

	// 消息没有注册
	if nil == meta {
		return nil, nil, worldnet.NewErrorContext("该 msg 不存在 MessageMeta", msgid)
	}

	// 创建消息
	msg := meta.NewType()

	// 从字节数组转换为消息
	err := meta.Codec.Decode(data, msg)
	if nil != err {
		return nil, meta, err
	}

	return msg, meta, nil
}
