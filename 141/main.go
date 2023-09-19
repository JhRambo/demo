package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var ipLock sync.Map
var requestsPerMinute = 10 // 每分钟允许的最大请求次数

func main() {
	r := gin.Default()
	r.Use(HandlePanic, IPRateLimiter)

	r.GET("/hello", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Hello, World!",
		})
	})

	r.Run(":8088")
}

func IPRateLimiter(ctx *gin.Context) {
	ip := ctx.ClientIP()
	fmt.Println(ip)
	currentTime := time.Now()

	// 检查是否超过请求限制
	if !checkRequestLimit(ip, currentTime) {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"code":    -1,
			"message": "Too many requests",
		})
		return
	}

	ctx.Next()
}

func checkRequestLimit(ip string, currentTime time.Time) bool {
	// 获取该IP的锁
	lock, _ := ipLock.LoadOrStore(ip, &sync.Mutex{})
	mutex := lock.(*sync.Mutex)

	// 对该IP的请求进行加锁，保证同一时间只能有一个请求处理
	mutex.Lock()
	defer mutex.Unlock()

	// 检查上次请求的时间
	lastRequestTime, exists := ipLock.Load(ip + "_time")
	fmt.Println(lastRequestTime, exists)
	if exists {
		if lastRequestTime.(time.Time).Add(time.Minute).After(currentTime) {
			// 如果距离上次请求不足1分钟，则判断是否超过请求限制
			count, _ := ipLock.Load(ip + "_count")
			fmt.Println(count)
			if count.(int) > requestsPerMinute {
				// 请求次数超过限制
				return false
			}

			// 更新请求次数
			ipLock.Store(ip+"_count", count.(int)+1)
			return true
		}
	}

	// 如果不存在上次请求的时间或者距离上次请求超过1分钟，则重置计数器
	ipLock.Store(ip+"_time", currentTime)
	ipLock.Store(ip+"_count", 1)

	return true
}

func HandlePanic(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
			return
		}
	}()
	ctx.Next()
}
