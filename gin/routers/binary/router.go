package hello

import (
	handler_binary "demo/gin/handlers/binary"

	"github.com/gin-gonic/gin"
)

// 路由分组
func InitRouter(r *gin.Engine) {
	g := r.Group("/hello")
	g.POST("/sayhello", handler_binary.UploadFile)
}
