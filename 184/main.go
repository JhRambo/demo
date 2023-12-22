package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	m := &sync.Map{}

	// 启动 10 个 goroutine 并发写入键值对
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Store(n, n*2)
			fmt.Printf("store key=%d value=%d\n", n, n*2)
		}(i)
	}

	// 启动 10 个 goroutine 并发读取键值对
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if v, ok := m.Load(n); ok {
				fmt.Printf("load key=%d value=%d\n", n, v.(int))
			} else {
				fmt.Printf("key=%d not found\n", n)
			}
		}(i)
	}

	wg.Wait()

	// 输出当前 map 中的所有键值对
	fmt.Println("all items in map:")
	m.Range(func(k, v interface{}) bool {
		fmt.Printf("key=%d value=%d\n", k.(int), v.(int))
		return true
	})
}
