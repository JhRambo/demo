package middlewares

import (
	"fmt"
	"net/http"

	"demo/docker/config"

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
		ctx.Error(fmt.Errorf(resp.Message))
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, resp)
		return
	}
	ctx.Next()
}
