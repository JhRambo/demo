package main

import (
	"context"
	"demo/gin/client"
	"demo/gin/config"
	"demo/gin/middlewares"
	"demo/gin/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// gateway
func main() {
	ctx := context.Background()
	client.InitGRPCClients()
	conn, _ := client.GetGRPCClient(ctx, config.SERVER_NAME1)
	r := gin.Default()
	middlewarres := &middlewares.Auth{}
	r.Use(middlewarres.MyAuth) //自定义中间件
	routers.InitRouter(r, conn)
	if err := r.Run(fmt.Sprintf(":%d", config.GW_PORT)); err != nil {
		log.Fatalf("GATEWAY could not run :%v", err)
	}
	log.Println("GATEWAY on http://0.0.0.0:8088")
}
