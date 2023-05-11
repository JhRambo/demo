package main

import (
	"fmt"
	"time"
)

// 带缓冲区的channel:
// 写入阻塞条件：缓冲区满
// 取出阻塞条件：缓冲区没有数据

// 不带缓冲区的channel:
// 写入阻塞条件：同一时间没有另外一个线程对该chan进行取操作
// 取出阻塞条件：同一时间没有另外一个线程对该chan进行写操作
func main() {
	// 1.带缓冲区的channel
	// 这段代码立即就会打印"send over"，3秒后才会打印"receive over"
	// 这是因为带缓冲区的chan在执行发送的操作时只要缓冲区满了就会被阻塞
	// send over
	// receive over
	// 222222
	// a======= 1
	// b======= 2
	// 111111
	ch1 := make(chan int, 2) //可读可写带缓冲区的channel
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("receive over")
		fmt.Println("222222")
		a := <-ch1
		b := <-ch1
		fmt.Println("a=======", a) //先进先出 输出1
		fmt.Println("b=======", b)
	}()
	ch1 <- 1 //写数据
	ch1 <- 2 //写数据
	fmt.Println("send over")
	time.Sleep(time.Second * 5)
	fmt.Println("111111")

	// 2.不带缓冲区的channel
	// 这段代码会先3秒后输出"receive over"，然后才会输出"send over"
	// 这是因为 ch<-1 的操作优先于 <-ch 执行，ch<-1执行的瞬间就被block，直到3秒后<-ch执行完之后ch<-1的操作才会unblock
	// 	receive over
	// 222222
	// send over
	// a======= 1
	// b======= 2
	// 111111
	ch2 := make(chan int, 0) //可读可写不带缓冲区的channel
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("receive over")
		fmt.Println("222222")
		a := <-ch2
		b := <-ch2
		fmt.Println("a=======", a) //先进先出 输出1
		fmt.Println("b=======", b)
	}()
	ch2 <- 1 //写数据
	ch2 <- 2 //写数据
	fmt.Println("send over")
	time.Sleep(time.Second * 5)
	fmt.Println("111111")
}
