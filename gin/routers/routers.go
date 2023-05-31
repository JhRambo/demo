package routers

import (
	handler_binary "demo/gin/handlers/binary"
	handler_hello "demo/gin/handlers/hello"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouters(r *gin.Engine) {
	r.POST("/binary/uploadfile", handler_binary.UploadFile)
	r.POST("/hello/sayhello", handler_hello.SayHello)
	r.POST("/hello/saygoodbye", handler_hello.SayGoodbye)

}
