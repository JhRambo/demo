package middlewares

import (
	"context"
	"demo/docker/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Timeout(ctx *gin.Context) {
	// 创建一个超时上下文
	ctxReq, cancel := context.WithTimeout(ctx.Request.Context(), config.TIMEOUT)
	defer cancel()
	// 将新的上下文对象附加到请求上下文中
	ctx.Request = ctx.Request.WithContext(ctxReq)

	ch := make(chan struct{})
	go func(ctx *gin.Context, ctxReq context.Context, ch chan<- struct{}) {
		defer func() {
			close(ch)
		}()
		ctx.Next()
	}(ctx.Copy(), ctxReq, ch)

	// 等待请求完成或者超时
	select {
	case <-ch:
		return
	case <-ctxReq.Done():
		// 检查是否超时，如果超时则中止请求处理
		err := ctxReq.Err()
		if err == context.DeadlineExceeded {
			resp := &config.GWResponse{
				Code:    -10003,
				Message: "请求超时",
			}
			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusGatewayTimeout, resp)
			return
		}
	}
}
