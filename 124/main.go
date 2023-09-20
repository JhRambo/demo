package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	// 故意创建一个 CPU 密集型任务
	go func() {
		for {
			fibonacci(30)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {}
}
