package main

import (
	"fmt"
	"time"
)

func strWorker(ch chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("do something with strWorker...")
		ch <- "str"
		if i == 5 {
			close(ch)
		}
	}
}

func intWorker(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("do something with intWorker...")
		ch <- 1
		if i == 5 {
			close(ch)
		}
	}
}

// select语句只能用于channel信道的IO操作，每个case都必须是一个信道。
// 如果不设置 default条件，当没有IO操作发生时，select语句就会一直阻塞；
// 如果有一个或多个IO操作发生时，Go运行时会随机选择一个case执行，但此时将无法保证执行顺序；
// 对于case语句，如果存在信道值为nil的读写操作，则该分支将被忽略，可以理解为相当于从select语句中删除了这个case；
// 对于空的 select语句，会引起死锁；
// 对于在 for中的select语句，不能添加 default，否则会引起cpu占用过高的问题；
func main() {
	chStr := make(chan string, 5)
	chInt := make(chan int, 5)
	strWorker(chStr)
	intWorker(chInt)
	for {
		select {
		case <-chStr:
			fmt.Println("get value from strWorker")
		case <-chInt:
			fmt.Println("get value from intWorker")
		default:
			fmt.Println("没有数据了")
			return
		}
	}
}
