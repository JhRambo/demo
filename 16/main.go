package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mutex sync.RWMutex //读写互斥锁

func fn1() {
	mutex.Lock() //写锁	独占
	fmt.Println("写数据")
	time.Sleep(time.Second * 2) //每个goroutine休眠时间
	mutex.Unlock()
	wg.Done()
}

func fn2() {
	mutex.RLock() //读锁	共享
	fmt.Println("....读数据")
	time.Sleep(time.Second * 1) //每个goroutine休眠时间
	mutex.RUnlock()
	wg.Done()
}

// 协程
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go fn1()
		wg.Add(1)
		go fn2()
	}
	wg.Wait()
	println("main函数退出...")
}
