package main

import (
	grpc_client "demo/gin/client"
	"demo/gin/config"
	"demo/gin/middlewares"
	"demo/gin/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// gateway
func main() {
	grpc_client.InitGRPCClients()
	r := gin.Default()
	middlewarres := &middlewares.Auth{}
	r.Use(middlewarres.MyAuth) //自定义中间件
	routers.InitRouters(r)     //gin 路由
	log.Println("GATEWAY on http://0.0.0.0:8088")
	if err := r.Run(fmt.Sprintf(":%d", config.GW_PORT)); err != nil {
		log.Fatalf("GATEWAY could not run :%v", err)
	}
}
