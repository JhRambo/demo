package middlewares

import (
	"fmt"
	"net/http"

	"demo/docker/config"
	"demo/logs"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter *rate.Limiter

func init() {
	limiter = rate.NewLimiter(config.TOKENS_PER_SECOND, config.TOKENS_BUCKET_SIZE)
}

func Rate(ctx *gin.Context) {
	if !limiter.Allow() {
		resp := &config.GWResponse{
			Code:    -10002,
			Message: "服务器繁忙",
		}
		ctx.Error(fmt.Errorf(logs.GetErrorLocation(resp.Message))) //记录错误发生的位置
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, resp)
		return
	}
	ctx.Next()
}
