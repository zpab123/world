// /////////////////////////////////////////////////////////////////////////////
// buffer 对象池

package buffer

import (
	"sync"
)

// /////////////////////////////////////////////////////////////////////////////
// 包初始化

const (
	MIN_LEN         uint32 = 128              // buffer 最小长度
	MAX_LEN         uint32 = 25 * 1024 * 1024 // buffer 最大长度
	_CAP_GROW_SHIFT uint   = 2                // 二进制数据 位计算变量
)

var (
	buffLenSlice []uint32                                        // buffer 长度切片，比如 [32 64 128 256 512 2048 ...] 长度
	bufferPools  map[uint32]*sync.Pool = map[uint32]*sync.Pool{} // buffer 对象池 buffLen -> *sync.Pool
)

func init() {
	// 添加 256 512 2048 ... 等长度数组
	length := MIN_LEN << _CAP_GROW_SHIFT
	for length < MAX_LEN {
		buffLenSlice = append(buffLenSlice, length)
		length <<= _CAP_GROW_SHIFT
	}

	// 添加最大长度
	buffLenSlice = append(buffLenSlice, MAX_LEN)

	// 创建 buffer 对象池
	for _, bufLen := range buffLenSlice {
		// 创建函数
		newBuff := func() interface{} {
			return make([]byte, bufLen)
		}

		// 创建对象池
		bufferPools[bufLen] = &sync.Pool{
			New: newBuff,
		}
	}
}

// /////////////////////////////////////////////////////////////////////////////
// public api
