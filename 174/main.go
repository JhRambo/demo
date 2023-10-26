package main

import (
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func main() {
	// 创建 gRPC 服务器，并注册服务
	grpcServer := grpc.NewServer()
	// 在这里注册你的 gRPC 服务
	// ...
	// 启动 gRPC 服务器
	go func() {
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("Starting gRPC server...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// 创建 gRPC-Web 处理器
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	// 创建 HTTP 服务器
	mux := http.NewServeMux()
	// 在这里添加你的 HTTP 路由和处理函数
	// ...
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(r) {
			// 如果是 gRPC-Web 请求，则交给 gRPC-Web 处理器处理
			wrappedGrpc.ServeHTTP(w, r)
		} else {
			// 否则，交给普通的 HTTP 处理器处理
			// ...
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, World!"))
		}
	})

	// 启动 gRPC 服务器和 HTTP 服务器
	log.Println("Starting gRPC-Web server...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to serve gRPC-Web: %v", err)
	}
}
