package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	最终执行顺序：

111
333
222
*/
func main() {
	router := gin.Default()

	// 定义一个中间件
	router.Use(func(ctx *gin.Context) {
		// 在这个中间件中，我们可以执行一些预处理的逻辑
		fmt.Println("111")

		// 将请求传递给下一个中间件或处理器函数
		ctx.Next()

		// 这里可以处理请求完成后的一些后续逻辑
		fmt.Println("222")

	})

	// 定义一个处理器函数
	router.GET("/", func(ctx *gin.Context) {
		// 处理器函数的逻辑

		// 在处理器函数中，我们可以根据需要对请求进行处理，并最终生成响应

		// 这里不需要调用 ctx.Next()，因为处理器函数是处理请求的最后一环，到这里就不会再传递给下一个处理器函数了

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})

		fmt.Println("333")
	})

	// 启动服务器
	router.Run(":8080")
}
