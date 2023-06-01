package middlewares

import (
	"context"
	"demo/gin/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Auth struct{}

func (x *Auth) MyAuth(ctx *gin.Context) {
	// 创建一个取消上下文
	ctxReq, cancel := context.WithTimeout(ctx.Request.Context(), time.Second*10)
	// ctxReq, cancel := context.WithCancel(ctx.Request.Context())
	defer cancel()

	// 将新的上下文对象附加到请求上下文中
	ctx.Request = ctx.Request.WithContext(ctxReq)

	ch := make(chan int8)

	go func() {
		ctx.Next()
		ch <- 1
	}()

	// 等待请求完成或者超时
	select {
	case <-ch:
		return
	// case <-time.After(time.Second * 5):
	// 	// 如果已经超时，则返回错误响应
	// 	ctx.AbortWithStatusJSON(http.StatusGatewayTimeout, &config.GWResponse{
	// 		Code:    -1,
	// 		Message: "请求超时",
	// 	})
	// 	return
	case <-ctxReq.Done():
		err := ctxReq.Err()
		// 检查上下文超时是由于取消还是时间超时引起的
		if err == context.DeadlineExceeded {
			ctx.AbortWithStatusJSON(http.StatusGatewayTimeout, &config.GWResponse{
				Code:    -1,
				Message: "请求超时",
			})
		} else {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
}
