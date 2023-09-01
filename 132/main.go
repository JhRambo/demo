package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 定义路由和处理函数
	router.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.Request.RequestURI) // /hello?a=1&b=2
		fmt.Println(c.Request.URL.Path)   // /hello
	})

	// 启动 Gin 服务器
	router.Run(":8080")
}
