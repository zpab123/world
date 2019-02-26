// /////////////////////////////////////////////////////////////////////////////
// 函数工具

package utils

import "github.com/zpab123/zaplog"

// /////////////////////////////////////////////////////////////////////////////
// public api

// 运行函数f，并返回f造成的panic错误
func CatchPanic(f func()) interface{} {
	var err error

	defer func() {
		err = recover()

		if err != nil {
			zaplog.TraceError("函数 %s 引发系统恐慌: %s", f, err)
		}
	}()

	// 运行函数
	f()

	return err
}

// 运行函数f，并返回该函数是否存在panic错误
func RunPanicless(f func()) bool {
	// 结果
	var panicless bool

	// 系统恐慌捕获
	defer func() {
		err := recover()
		panicless = (err == nil)

		if err != nil {
			zaplog.TraceError("函数 %s 引发系统恐慌: %s", f, err)
		}
	}()

	// 运行函数
	f()

	return panicless
}

// for循环调用f函数，直到f出现系统恐慌
func RepeatFunc(f func()) {
	var haveErr bool = false

	for !haveErr {
		haveErr = RunPanicless(f)
	}
}

// NextLargerKey finds the next key that is larger than the specified key,
// but smaller than any other keys that is larger than the specified key
func NextLargerKey(key string) string {
	return key + "\x00" // the next string that is larger than key, but smaller than any other keys > key
}
