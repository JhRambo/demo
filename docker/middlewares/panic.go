package middlewares

import (
	"demo/docker/config"
	"demo/docker/utils"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandlePanic(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// message := fmt.Sprintf("panic：%v", r)
			message := fmt.Sprintf("panic：%s", GetPanicMessage(r))
			ctx.Error(utils.GetError(message))

			resp := &config.GWResponse{
				Code:    -10004,
				Message: "系统错误",
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			return
		}
	}()
	ctx.Next()
}

// 捕获异常详细信息
func GetPanicMessage(r interface{}) string {
	// 获取调用栈信息
	stack := make([]byte, 1<<16)
	n := runtime.Stack(stack, false)
	// 获取发生 panic 的行号
	pc := make([]uintptr, 10)
	n = runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	message := fmt.Sprintf("%s\n", r)
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		message += fmt.Sprintf("%s(%d)\n", frame.File, frame.Line)
	}
	message = strings.TrimRight(message, "\n") // 去掉末尾换行符
	return message
}
