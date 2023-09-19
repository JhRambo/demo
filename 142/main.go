package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var taskPool chan struct{} // Goroutine 连接池

func main() {
	// 初始化 Goroutine 连接池
	maxConcurrency := 1
	taskPool = make(chan struct{}, maxConcurrency)

	r := gin.Default()
	r.GET("/hello", handleRequest)
	r.Run(":8088")
}

func handleRequest(c *gin.Context) {
	select {
	case taskPool <- struct{}{}: // 尝试获取一个任务槽
		go processRequest(c) // 启动一个 Goroutine 处理请求
	default:
		c.JSON(http.StatusTooManyRequests, gin.H{"message": "Too many requests"})
		return // 无法获取任务槽时直接返回
	}
}

func processRequest(c *gin.Context) {
	// 将连接放回连接池
	defer func() {
		<-taskPool // 任务处理完成后释放任务槽
	}()

	// 在这里处理具体的请求逻辑
	// ...
	// log.Println("Processing request...")
	time.Sleep(1 * time.Second)
}
