package routers

import (
	handler_binary "demo/gin/handlers/binary"
	handler_hello "demo/gin/handlers/hello"
	handler_msgpack "demo/gin/handlers/msgpack"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouters(r *gin.Engine) {
	r.POST("/binary/uploadfile", handler_binary.UploadFile)
	r.POST("/hello/sayhello", handler_msgpack.MsgPackProtocol)
	r.POST("/hello/saygoodbye", handler_hello.SayGoodbye)
	r.POST("/msgpack/protocol", handler_msgpack.MsgPackProtocol)

}
