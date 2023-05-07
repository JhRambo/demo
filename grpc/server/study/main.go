package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"demo/grpc/proto/study"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	// pb.go中自动生成的，是个空结构体
	study.UnimplementedStudyHttpServer
}

const IP = "127.0.0.1"
const PORT = ":9099"

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayStudy(ctx context.Context, in *study.StudyRequest) (*study.StudyResponse, error) {
	var infos []*study.Info
	info1 := &study.Info{
		DevId: "111",
		Date:  "2026-03-01",
		Time:  "01:01:01",
	}
	info2 := &study.Info{
		DevId: "111",
		Date:  "2026-03-01",
		Time:  "01:02:02",
	}
	info3 := &study.Info{
		DevId: "222",
		Date:  "2026-03-01",
		Time:  "01:01:01",
	}
	infos = append(infos, info1, info2, info3)

	ss := make(map[string][]*study.Info)
	for _, v := range infos {
		ss[v.DevId] = append(ss[v.DevId], v) //同一个key，append
	}
	fmt.Println("==============", ss)
	return &study.StudyResponse{
		Code:    200,
		Message: "成功",
		List:    infos,
	}, nil
}

// http grpc监听同一个端口
func main2() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册StudyHttp service到server
	study.RegisterStudyHttpServer(s, &Server{})
	// gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = study.RegisterStudyHttpHandlerFromEndpoint(context.Background(), gwmux, IP+PORT, dops)
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
