package main

import (
	"context"
	"demo/grpc-jwt/api"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

const (
	// grpc-gateway服务 和 grpc服务监听同一个端口
	port = ":8088"
)

func main() {
	log.Println("网关，server启动中...")
	// 1.创建grpc-gateway服务，转发到grpc的port端口	【网关gw】
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// 调用了自动生成代码中的RegisterPingHandlerFromEndpoint方法完成上下游调用的绑定
	err := api.RegisterPingHandlerFromEndpoint(context.Background(), gwmux, "localhost"+port, opts)
	if err != nil {
		log.Fatal(err)
	}

	// 2.创建grpc服务 【服务层server】
	rpcServer := grpc.NewServer()
	// 注册服务
	api.RegisterPingServer(rpcServer, new(api.Server))

	// 创建http服务，监听port端口，并调用上面的两个服务来处理请求
	http.ListenAndServe(
		port,
		grpcHandlerFunc(rpcServer, gwmux),
	)
}

// grpcHandlerFunc 根据请求头判断是grpc请求还是grpc-gateway请求
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			allowCORS(otherHandler).ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	fmt.Println("preflight request for:", r.URL.Path)
	return
}
