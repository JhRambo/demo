package binary

import (
	handler_binary "demo/gin/handlers/binary"
	"github.com/gin-gonic/gin"
)

// 路由分组
func InitRouter(r *gin.Engine) {
	g := r.Group("/binary/")
	g.POST("uploadfile", handler_binary.UploadFile)

}
