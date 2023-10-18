package main

import (
	"fmt"
	"runtime"
)

func main() {
	A()
}

func A() {
	defer func() {
		if r := recover(); r != nil {
			message := GetPanicMessage()
			fmt.Println(message)
		}
	}()

	a := 1
	b := 0
	fmt.Println(a / b)
}

// 捕获异常信息
func GetPanicMessage() string {
	// 获取调用栈信息
	stack := make([]byte, 1<<16)
	n := runtime.Stack(stack, false)
	// 获取发生 panic 的行号
	pc := make([]uintptr, 10)
	n = runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	var message string
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		message += fmt.Sprintf("panic：%s(%d)\n", frame.File, frame.Line)
	}
	return message
}
