package main

import "fmt"

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("发生了 panic：", r)
	// 	}
	// }()

	// 触发 panic
	panic("出现了一个错误")

	// 这行代码不会被执行到
	fmt.Println("这行代码不会被执行到")
}
