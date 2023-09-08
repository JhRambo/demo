package main

import (
	"demo/docker/middlewares"
	"demo/docker/utils"
	"demo/logs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 跨域处理
	r.Use(middlewares.ErrorLogger)

	r.POST("/logs", func(ctx *gin.Context) {
		utils.CreateDir("logs")
		ctx.Error(fmt.Errorf("测试运行过程中，出错记录日志，如果是在容器中运行，则日志只会记录到容器中，不会同步到宿主机，不过可以通过挂载路径实现同步"))
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "POST Resource"})
	})

	// 启动网关
	logs.Infof("GATEWAY on http://0.0.0.0:%d", 8088)
	if err := r.Run(fmt.Sprintf(":%d", 8088)); err != nil {
		logs.Errorf("GATEWAY could not run :%v", err)
	}
}
