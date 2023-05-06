package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"demo/grpc/proto/hello"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 定义结构体，在调用注册api的时候作为入参，
// 该结构体会带上SayHello方法，里面是业务代码
// 这样远程调用时就执行了业务代码了
type Server struct {
	// pb.go中自动生成的，是个空结构体
	hello.UnimplementedHelloHttpServer
}

const IP = "127.0.0.1"
const PORT = ":9099"

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: in.Name + "，hello world"}, nil
}

// http grpc监听同一个端口
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册HelloHttp service到server
	hello.RegisterHelloHttpServer(s, &Server{})

	// gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = hello.RegisterHelloHttpHandlerFromEndpoint(context.Background(), gwmux, IP+PORT, dops)
	if err != nil {
		log.Fatalln("Failed to register gwmux:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// 定义HTTP server配置
	gwServer := &http.Server{
		Addr:    IP + PORT,
		Handler: grpcHandlerFunc(s, mux), // 请求的统一入口
	}
	log.Println("Serving on http://" + IP + PORT)
	log.Fatalln(gwServer.Serve(lis)) // 启动HTTP服务
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
