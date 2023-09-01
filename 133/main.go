package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func SlowRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		responseTime := time.Since(startTime)

		if responseTime > 1*time.Second {
			logSlowRequest(c.Request.URL.Path, responseTime)
		}
	}
}

func logSlowRequest(uri string, responseTime time.Duration) {
	now := time.Now()
	logDir := "logs"
	logFileName := fmt.Sprintf("slow-%s.log", now.Format("20060102"))
	logFilePath := filepath.Join(logDir, logFileName)

	// 检查 logs 文件夹是否存在，不存在则创建
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			log.Printf("Failed to create log directory: %v", err)
			return
		}
	}

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		return
	}
	defer file.Close()

	log.SetOutput(file)

	// 记录请求时间大于10秒的相关信息
	log.Printf("URI: %s, Response Time: %s", uri, responseTime.String())
}

func main() {
	router := gin.Default()
	router.Use(SlowRequestLogger())

	// 添加其他路由和处理程序
	// ...
	router.GET("/hello", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		c.String(200, "Hello, Secure!")
	})

	router.Run(":8080")
}
