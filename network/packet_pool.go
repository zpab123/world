// /////////////////////////////////////////////////////////////////////////////
// pcket 需要用到的对象池

package network

import (
	"sync"

	"github.com/zpab123/world/model" // 全局 [常量-基础数据类型-接口] 集合
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化 -- 创建 buffer 对象池 和 packet 对象池

const (
	_CAP_GROW_SHIFT = uint(2)                                          // 二进制数据 位计算变量
	_MAX_BODY_LEN   = model.C_PACKET_MAX_LEN - model.C_PACKET_HEAD_LEN // body 最大长度
)

var (
	// buffer 长度切片，比如 [256 512 2048 ...] 长度
	buffLenSlice []uint32

	// Packet 对象池
	packetPool sync.Pool = sync.Pool{
		New: newPacket,
	}

	// buffer 对象池 buffLen -> *sync.Pool
	bufferPools map[uint32]*sync.Pool = map[uint32]*sync.Pool{}
)

// network 初始化函数
func init() {
	// 添加 256 512 2048 ... 等各个body长度数组
	length := (_MIN_PAYLOAD_LEN << _CAP_GROW_SHIFT)
	for length < _MAX_BODY_LEN {
		buffLenSlice = append(buffLenSlice, length)
		length <<= _CAP_GROW_SHIFT
	}

	// 添加最大 body 长度
	buffLenSlice = append(buffLenSlice, _MAX_BODY_LEN)

	// 创建 buffer 对象池
	for _, bufLen := range buffLenSlice {
		// 创建函数
		newBuf := func() interface{} {
			return make([]byte, _HEAD_LEN+bufLen) // 消息头 + 有效载荷
		}

		// 创建对象池
		bufferPools[ln] = &sync.Pool{
			New: newBuf,
		}
	}
}

// /////////////////////////////////////////////////////////////////////////////
// 私有 api

// 从对象池中获取1个 packet
func getPacketFromPool() *Packet {
	// 获取 *Packet
	pkt := packetPool.Get().(*Packet)

	return pkt
}

// 根据 need ，计算需要从 bufferPools 中取出哪个对象池
func getPoolKey(need uint32) uint32 {
	for _, ln := range buffLenSlice {
		if ln >= need {
			return ln
		}
	}

	return _MAX_BODY_LEN
}
