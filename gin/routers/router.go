package routers

import (
	router_binary "demo/gin/routers/binary"
	router_hello "demo/gin/routers/hello"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter(r *gin.Engine) {
	router_binary.InitRouter(r)
	router_hello.InitRouter(r)

}
