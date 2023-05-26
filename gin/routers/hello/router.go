package hello

import (
	handler_hello "demo/gin/handlers/hello"

	"github.com/gin-gonic/gin"
)

// 路由分组
func InitRouter(r *gin.Engine) {
	g := r.Group("/hello")
	g.POST("/sayhello", handler_hello.SayHello)
}
