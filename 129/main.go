package main

import (
	"demo/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	logs.Init("logs", "error", 3, false)
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logs.Errorf("Panic occurred - Path: %s, Panic: %s", c.Request.URL.Path, err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Internal Server Error",
				})
			} else if c.Writer.Status() >= 400 {
				// 记录非成功请求的错误信息到日志文件
				logs.Errorf("Error occurred - Path: %s, Status: %d", c.Request.URL.Path, c.Writer.Status())
			}
		}()

		c.Next()
	}
}
