package main

import (
	"fmt"
	"log"
)

func main() {
	defer_call()
}

// 需要注意的是，即使是在出现异常时，也会保证 defer 语句中的函数能够被完整地执行。这是 defer 的一个重要特性，可以用来进行资源释放、错误处理等操作
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	// 不同位置，打印顺序不一样，有无打印顺序也不一样
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	panic("触发异常")
}
