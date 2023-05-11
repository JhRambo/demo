package main

import (
	gw "demo/grpc/proto/hello"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	// 将外部RESTful请求转发到GRPC服务的入口处
	echoEndpoint = flag.String("echo_endpoint", "127.0.0.1:8081", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// mux := runtime.NewServeMux() //多路复用器，它将根据 JSON/Restful 请求的路径将请求路由到各种注册服务
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 调用了自动生成代码中的RegisterHelloHttpHandlerFromEndpoint方法完成上下游调用的绑定
	err := gw.RegisterHelloHttpHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":8088", mux) //这是对外提供RESTful服务的端口
}

/*
	网关入口

HTTP GRPC监听不同的端口
*/
func main() {
	log.Println("网关启动中...")
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}
