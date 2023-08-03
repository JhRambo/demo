package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func main() {
	r := gin.Default()
	r.Use(AddMetadataToContext) // 使用中间件
	r.GET("/", func(c *gin.Context) {
		md, _ := metadata.FromIncomingContext(c.Request.Context())
		fmt.Println(md)
	})
	r.Run(":8080")
}

func AddMetadataToContext(c *gin.Context) {
	// 提取请求头中的自定义字段，并创建 metadata
	md := metadata.Pairs("key1", "1111", "key2", "222")

	// 将 metadata 添加到请求的上下文中
	ctx := metadata.NewIncomingContext(c.Request.Context(), md)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
