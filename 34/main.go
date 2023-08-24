package main

import (
	"fmt"
	"time"
)

func firstRoutine(ch chan int) {
	// 模拟耗时操作
	time.Sleep(time.Second * 1)

	result := 100

	// 将结果发送到通道中
	ch <- result
}

func secondRoutine(ch chan int) {
	// 模拟耗时操作
	time.Sleep(time.Second * 1)

	result := 200

	// 将结果发送到通道中
	ch <- result
}

func thirdRoutine(ch1, ch2 chan int) {
	// 使用select语句从两个通道中接收结果，只会输出一个case，如果需要输出所有的case，则需要for循环来实现，参考demo50
	select {
	case result1 := <-ch1:
		fmt.Println("Third routine received from first:", result1)
	case result2 := <-ch2:
		fmt.Println("Third routine received from second:", result2)
	}

	// 继续执行其他操作
	fmt.Println("Third routine is running")
}

func main() {
	// 创建两个整型通道
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 启动第一个协程
	go firstRoutine(ch1)

	// 启动第二个协程
	go secondRoutine(ch2)

	// 启动第三个协程
	go thirdRoutine(ch1, ch2)

	// 主协程等待一段时间，以便观察输出结果
	time.Sleep(time.Second * 2)
}
