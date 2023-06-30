package main

import (
	"fmt"
	"log"
	"time"
)

func strWorker(ch chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Println("do something with strWorker...")
		ch <- "str"
		if i == 5 {
			close(ch)
		}
	}
}

func intWorker(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Microsecond)
		fmt.Println("do something with intWorker...")
		ch <- 1
		if i == 5 {
			close(ch)
		}
	}
}

// select语句只能用于channel信道的IO操作，每个case都必须是一个信道；
// 如果有一个或多个IO操作发生时，Go运行时会随机选择一个case执行，但此时将无法保证执行顺序；
// 对于case语句，如果存在信道值为nil的读写操作，则该分支将被忽略，可以理解为相当于从select语句中删除了这个case；
// 总的来说，如果希望 select 语句在成功完成一个操作后立即返回，则需要使用 return 或 break。但是，如果需要继续等待，或者希望处理所有 case 分支，则不需要使用它们；
func main() {
	chStr := make(chan string, 5)
	chInt := make(chan int, 5)
	strWorker(chStr)
	intWorker(chInt)
	for {
		select {
		case <-chStr:
			fmt.Println("get value from strWorker")
			// return
		case <-chInt:
			fmt.Println("get value from intWorker")
			// return
		default:
			log.Fatalln("没有数据了")
		}
	}
}
