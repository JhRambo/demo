package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 跨域中间件
	r.Use(Cors())

	// 添加其他路由和处理程序
	r.GET("/api/resource", getResourceHandler)
	r.POST("/api/resource", postResourceHandler)

	r.Run(":8088")
}

// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// 处理 GET 请求的处理程序示例
func getResourceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GET Resource"})
}

// 处理 POST 请求的处理程序示例
func postResourceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "POST Resource"})
}
