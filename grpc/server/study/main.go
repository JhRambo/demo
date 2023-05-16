package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	"demo/grpc/proto/study"
	pb "demo/grpc/proto/study"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	// pb.go中自动生成的，是个空结构体
	pb.UnimplementedStudyHttpServer
}

const IP = "127.0.0.1"
const PORT = ":9099"

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayStudy(ctx context.Context, req *study.StudyRequest) (*study.StudyResponse, error) {
	return &pb.StudyResponse{
		Code:    200,
		Message: "成功",
		Name:    req.Name,
	}, nil
}

// http grpc监听同一个端口
func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册StudyHttp服务到server层
	pb.RegisterStudyHttpServer(s, &Server{})

	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// http转grpc	通过context去连接, 注册在runtime.NewServeMux()上面
	err = pb.RegisterStudyHttpHandlerFromEndpoint(context.Background(), gwmux, IP+PORT, dops)
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
