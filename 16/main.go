package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mutex sync.RWMutex //读写互斥锁

func fn1() {
	mutex.Lock()
	fmt.Println("写数据")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

func fn2() {
	mutex.RLock()
	fmt.Println("....读数据")
	time.Sleep(time.Second * 1)
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
