package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

const MaxLogFileSize = 10 * 1024 * 1024 // 10MB

var currentLogFile *os.File

func main() {
	// 创建一个新的 Gin 路由器
	router := gin.Default()

	// 使用自定义的 ErrorLogger 中间件
	router.Use(errorLogger())

	// 定义路由和处理程序
	router.GET("/panic", func(c *gin.Context) {
		panic("panic error 123")
	})

	// 启动服务器
	router.Run(":8080")
}

// 自定义 ErrorLogger 中间件
func errorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录 panic 信息到日志文件
				log.Printf("Panic occurred - Path: %s, Panic: %s", c.Request.URL.Path, err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Internal Server Error",
				})
			} else if c.Writer.Status() >= 400 {
				// 记录非成功请求的错误信息到日志文件
				log.Printf("Error occurred - Path: %s, Status: %d", c.Request.URL.Path, c.Writer.Status())
			}
		}()

		// 检查当前日志文件大小
		if currentLogFile == nil || isFileSizeExceeded(currentLogFile) {
			// 创建新的日志文件
			createNewLogFile()
		}

		log.SetOutput(currentLogFile)

		c.Next()
	}
}

// 检查当前日志文件大小是否超过阈值
func isFileSizeExceeded(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Size() > MaxLogFileSize
}

// 创建新的日志文件
func createNewLogFile() {
	now := time.Now()
	logFileName := "error-" + now.Format("20060102150405") + ".log"
	logFilePath := filepath.Join("logs", logFileName)

	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Println("Failed to create logs directory:", err)
		return
	}

	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Println("Failed to create new log file:", err)
		return
	}

	if currentLogFile != nil {
		currentLogFile.Close()
	}
	currentLogFile = logFile
}
