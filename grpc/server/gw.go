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
	port         = 8088
)

type Token struct {
	Token string
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// mux := runtime.NewServeMux() //多路复用器，它将根据 JSON/Restful 请求的路径将请求路由到各种注册服务
	m := &runtime.JSONPb{} //定义以哪种数据格式返回给客户端
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, m))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 调用了自动生成代码中的RegisterHelloHttpHandlerFromEndpoint方法完成上下游调用的绑定
	err := gw.RegisterHelloHttpHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux) //这是对外提供RESTful服务的端口
}

/*
	网关入口

HTTP GRPC监听不同的端口
*/
func main() {
	log.Println(fmt.Sprintf("gw网关启动中，监听%d端口...", port))
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}
