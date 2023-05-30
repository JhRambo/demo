package binary

import (
	handler_binary "demo/gin/handlers/binary"
	"github.com/gin-gonic/gin"
)

// 路由分组
func InitRouter(r *gin.Engine) {
	g := r.Group("/binary")
	//这里要通过正则，循环遍历api方法：
	g.POST("/sayhello", handler_binary.SayHello)
}
