package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 定义路由和处理函数 uri path 的区别
	router.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.Request.RequestURI) // /hello?a=1&b=2
		fmt.Println(c.Request.URL.Path)   // /hello
	})

	// 定义路由和处理函数 uri path 的区别
	router.POST("/golang", func(c *gin.Context) {
		fmt.Println(c.Request.RequestURI) // /golang?a=1&b=2
		fmt.Println(c.Request.URL.Path)   // /golang
	})

	// 启动 Gin 服务器
	router.Run(":8080")
}
