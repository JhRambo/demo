package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func ErrorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := c.Writer.Status()
		if status >= http.StatusInternalServerError {
			errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
			logError(status, c.Request.Method, c.Request.URL.Path, errorMessage)
		}
	}
}

func logError(status int, method string, path string, errorMessage string) {
	now := time.Now()
	logDir := "logs"
	logFileName := fmt.Sprintf("error-%s.log", now.Format("20060102"))
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

	// 记录错误的相关信息
	log.Printf("Timestamp: %s", now.Format("2006-01-02 15:04:05"))
	log.Printf("Status: %d", status)
	log.Printf("Method: %s", method)
	log.Printf("Path: %s", path)
	log.Printf("Error Message: %s", errorMessage)
}

func main() {
	router := gin.Default()
	router.Use(ErrorLogger())

	router.GET("/hello", func(c *gin.Context) {
		// 模拟内部服务器错误
		err := fmt.Errorf("Internal Server Error")
		c.Error(err) // 将错误添加到 Gin 上下文的错误列表中
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": -100002, "message": "系统内部错误"})
	})

	router.Run(":8080")
}
