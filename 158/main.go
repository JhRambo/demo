package main

import "fmt"

// 定义回调函数类型
type CallbackFunc func(string)

// 函数A，接收一个回调函数作为参数
func FunctionA(callback CallbackFunc) {
	// 执行某些操作
	result := "Hello, World!"

	// 调用回调函数，传递结果作为参数
	callback(result)
}

// 函数B，作为回调函数
func FunctionB(result string) {
	fmt.Println("FunctionB:", result)
}

func main() {
	// 调用函数A，传递函数B作为回调函数
	FunctionA(FunctionB)
}
