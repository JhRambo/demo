package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件函数1
func Middleware1(c *gin.Context) {
	fmt.Println("执行中间件函数1 - 开始")

	// 执行下一个中间件或路由处理函数
	c.Next()

	fmt.Println("执行中间件函数1 - 结束")
}

// 中间件函数2
func Middleware2(c *gin.Context) {
	fmt.Println("执行中间件函数2 - 开始")

	// 模拟耗时操作
	time.Sleep(2 * time.Second)

	// 执行下一个中间件或路由处理函数
	c.Next()

	fmt.Println("执行中间件函数2 - 结束")
}

// 中间件函数3
func Middleware3(c *gin.Context) {
	fmt.Println("执行中间件函数3 - 开始")

	// 模拟耗时操作
	time.Sleep(6 * time.Second)

	// 执行下一个中间件或路由处理函数
	c.Next()

	fmt.Println("执行中间件函数3 - 结束")
}

// 路由处理函数
func Handler(c *gin.Context) {
	fmt.Println("执行路由处理函数")

	// 返回响应
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func main() {
	router := gin.Default()

	// 注册中间件函数
	router.Use(Middleware1, Middleware2, Middleware3)

	// 注册路由及对应的处理函数
	router.GET("/check/health", Handler)

	// 启动服务器
	router.Run(":8080")
}
