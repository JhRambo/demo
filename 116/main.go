package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	doneChannel := make(chan struct{})

	go func() {
		sig := <-signalChannel
		fmt.Println("接收到信号:", sig)

		// 在这里执行异常关闭的逻辑
		fmt.Println("程序异常关闭")

		close(doneChannel) // 关闭doneChannel
	}()

	// 这里可以执行你的主要业务逻辑
	// 模拟业务逻辑运行
	select {
	case <-doneChannel:
		fmt.Println("程序退出")
		return
	case <-time.After(time.Minute * 10):
		fmt.Println("业务逻辑执行完毕")
		return
	}
}
