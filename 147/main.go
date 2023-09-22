package main

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	mutex           = &sync.Mutex{}
	condition       = sync.NewCond(mutex)
	tokenBucket     = 1               // 令牌桶容量
	tokensPerSecond = 1               // 每秒生成的令牌数
	isPaused        = false           // 是否暂停生成令牌
	requestQueue    = make(chan bool) // 请求队列，用于请求处理线程的阻塞和唤醒
	responseQueue   = make(chan bool) // 响应队列，用于令牌生成线程的阻塞和唤醒
)

func main() {
	r := gin.Default()

	r.GET("/api", handleRequest)

	go generateTokens() // 启动令牌生成线程

	r.Run(":8080")
}

func handleRequest(c *gin.Context) {
	// 检查是否有可用的令牌
	mutex.Lock()
	if tokenBucket <= 0 {
		isPaused = true  // 暂停生成令牌
		condition.Wait() // 阻塞等待，直到有新的令牌生成
		isPaused = false // 唤醒后恢复生成令牌
	}
	tokenBucket-- // 消耗一个令牌
	mutex.Unlock()

	c.JSON(200, gin.H{"message": "Request processed successfully"})
}

func generateTokens() {
	for {
		time.Sleep(time.Second / time.Duration(tokensPerSecond))

		mutex.Lock()
		if !isPaused && tokenBucket < 1 {
			tokenBucket++      // 生成一个新令牌
			condition.Signal() // 唤醒一个等待的请求处理线程
		}
		mutex.Unlock()
	}
}
