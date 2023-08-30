package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.Default()

	// 初始化 Secure 中间件并进行配置
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect:      false, // 强制使用 HTTPS
		BrowserXssFilter: true,  // 启用浏览器的 XSS 过滤器	纯api接口的项目，没有实际意义，可不用该中间件
	})

	// 将 Secure 中间件添加到全局中间件链中
	router.Use(func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			// 处理 Secure 中间件返回的错误
			c.AbortWithStatus(400)
			return
		}
		c.Next()
	})

	// 定义路由和处理函数
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, Secure!")
	})

	// 启动 Gin 服务器
	router.Run(":8080")
}
