package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AllowedOriginHeader = "Access-Control-Allow-Origin"
	AllowedHeaders      = "Access-Control-Allow-Headers"
	AllowedMethods      = "Access-Control-Allow-Methods"
	ContentTypeHeader   = "Content-Type"
	ContentType         = "application/json"
)

// 允许的域名列表
var allowedOrigins = []string{}

func Cors(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("Origin")
	if len(allowedOrigins) > 0 {
		// 检查请求的来源域名是否在允许的列表中
		if CheckOrigin(allowedOrigins, origin) {
			ctx.Header(AllowedOriginHeader, origin)
		}
	} else {
		ctx.Header(AllowedOriginHeader, "*")
	}
	ctx.Header(AllowedHeaders, "Content-Type, Content-Length, Authorization, Accept, X-Requested-With")
	ctx.Header(AllowedMethods, "GET, OPTIONS, POST, PUT, DELETE, PATCH, HEAD")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusOK)
		return
	}
}

// check origin
func CheckOrigin(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
