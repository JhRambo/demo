package main

import (
	"fmt"
	"sync"
)

// 错误写法 报 A WaitGroup must not be copied after first use
// 在 WaitGroup 第一次使用后，不能被拷贝
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) { //wg sync.WaitGroup	错误写法
			fmt.Printf("i:%d", i)
			wg.Done()
		}(&wg, i) //(wg, i)	错误写法
	}
	wg.Wait()
	fmt.Println("exit")
}
