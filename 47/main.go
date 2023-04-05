package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5) //有缓冲区
	for i := 0; i <= 4; i++ {
		ch <- i
		if i == 4 {
			close(ch)
			break
		}
	}
	readNum(ch)
}

// num： 0
// num： 1
// num： 2
// num： 3
// num： 4
// Channel is closed
func readNum(ch chan int) {
	for {
		if num, ok := <-ch; !ok {
			fmt.Println("Channel is closed") //管道关闭时输出
			return
		} else {
			fmt.Println("num：", num)
		}
	}
}
