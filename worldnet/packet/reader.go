// /////////////////////////////////////////////////////////////////////////////
// 网络数据读取工具

package packet

import (
	"encoding/binary"
	"io"

	"github.com/zpab123/world/worldnet" // 网络库
)

// /////////////////////////////////////////////////////////////////////////////
// public api

// 接收 Length-Type-Value 格式的封包流程
func RecvLTVPacket(reader io.Reader, maxSize int) (pkt interface{}, err error) {
	// Size 为 uint16，占2字节
	var sizeBuffer = make([]byte, bodySize)

	// 持续读取Size直到读到为止
	_, err = io.ReadFull(reader, sizeBuffer)

	// 发生错误时返回
	if err != nil {
		return
	}

	// 长度错误
	if len(sizeBuffer) < bodySize {
		return nil, ErrMinPacket
	}

	// 用小端格式读取 Size
	size := binary.LittleEndian.Uint16(sizeBuffer)

	// 数据包超过最大值
	if maxSize > 0 && size >= uint16(maxSize) {
		return nil, ErrMaxPacket
	}

	// 分配包体大小
	body := make([]byte, size)

	// 读取包体数据
	_, err = io.ReadFull(reader, body)

	// 发生错误时返回
	if err != nil {
		return
	}

	// type 数据丢失
	if len(body) < msgIDSize {
		return nil, ErrShortMsgID
	}

	// 读取消息id
	msgid := binary.LittleEndian.Uint16(body)

	// 读取数据
	msgData := body[msgIDSize:]

	// 将字节数组和消息ID用户解出消息

	return
}
