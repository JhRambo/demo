package main

import (
	"demo/docker/middlewares"
	"demo/logs"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 错误日志中间件
	r.Use(middlewares.ErrorLogger, middlewares.Rate, middlewares.Timeout)

	r.POST("/docker/info", func(ctx *gin.Context) {
		time.Sleep(2 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
		})

		// ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		// return

		// // 这里的代码不会被执行
		// ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	// 启动网关
	logs.Infof("GATEWAY on http://0.0.0.0:%d", 8088)
	if err := r.Run(fmt.Sprintf(":%d", 8088)); err != nil {
		logs.Errorf("GATEWAY could not run :%v", err)
	}
}
