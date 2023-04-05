package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 生产者
func fn1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i //给管道赋值
		fmt.Printf("%v--【写入数据】%v成功\n", time.Now().Unix(), i)
		time.Sleep(time.Microsecond * 10)
	}
	wg.Done()
	close(ch) // 用for循环遍历数据，需要关闭管道
}

// 消费者
func fn2(ch chan int) {
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("【读取数据】%v成功\n",i)
	// }
	for v := range ch { //用for range 循环遍历读取数据，不需要close管道
		fmt.Printf("%v--【读取数据】%v成功\n", time.Now().Unix(), v)
		time.Sleep(time.Microsecond * 1)
	}
	wg.Done()
}

// 协程结合管道，管道用于协程间的通信
func main() {
	var ch = make(chan int) //这里为什么不用分配内存空间也能写入数据，无缓存区的通道请参考demo41
	// ch <- 1
	// b := <-ch
	// fmt.Println(b) //all goroutines are asleep - deadlock!
	// fmt.Println(len(ch), cap(ch))	//0 0
	wg.Add(1)
	go fn1(ch) //协程1
	wg.Add(1)
	go fn2(ch) //协程2
	wg.Wait()
	println("main函数退出...")
}
