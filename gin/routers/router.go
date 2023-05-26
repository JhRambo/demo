package routers

import (
	router_hello "demo/gin/routers/hello"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	router_hello.InitRouter(r)
}
