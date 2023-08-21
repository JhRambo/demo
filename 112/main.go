package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("发生了1 panic：", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("发生了2 panic：", r)
		}
	}()

	// 当发生 panic 时，程序会立即停止当前函数的执行，并依次向上层函数传播，直到遇到恢复（recover）或达到最外层函数。如果没有进行恢复处理，程序将会崩溃并退出，并且会返回一个非零的退出码
	panic("出现了一个错误") //

	// 这行代码不会被执行到
	fmt.Println("这行代码不会被执行到")
}
