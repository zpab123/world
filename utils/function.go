// /////////////////////////////////////////////////////////////////////////////
// 函数工具

package utils

import "github.com/zpab123/zaplog"

// /////////////////////////////////////////////////////////////////////////////
// 对外api

// CatchPanic calls a function and returns the error if function paniced
func CatchPanic(f func()) (err interface{}) {
	defer func() {
		err = recover()
		if err != nil {
			zaplog.TraceError("%s 系统恐慌: %s", f, err)
		}
	}()

	f()

	return
}

// RunPanicless calls a function panic-freely
// 自由调用函数，并捕获恐慌
func RunPanicless(f func()) (panicless bool) {
	// 系统恐慌捕获
	defer func() {
		err := recover()
		panicless = err == nil
		if err != nil {
			zaplog.TraceError("%s 系统恐慌: %s", f, err)
		}
	}()

	// 运行函数
	f()

	return
}

// RepeatUntilPanicless runs the function repeatly until there is no panic
// 循环调用 f 函数，直到出现 系统恐慌
func RepeatFunc(f func()) {
	for !RunPanicless(f) {
	}
}

// NextLargerKey finds the next key that is larger than the specified key,
// but smaller than any other keys that is larger than the specified key
func NextLargerKey(key string) string {
	return key + "\x00" // the next string that is larger than key, but smaller than any other keys > key
}
