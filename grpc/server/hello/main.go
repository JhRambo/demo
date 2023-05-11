package main

import (
	"context"
	pb "demo/grpc/proto/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloHttpServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloHttpRequest) (*pb.HelloHttpResponse, error) {
	return &pb.HelloHttpResponse{
		Name: req.Token,
	}, nil
}

func main() {
	log.Println("server服务启动中...")
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册service到server
	pb.RegisterHelloHttpServer(s, &Server{})
	// 启动gRPC Server
	log.Fatal(s.Serve(lis))
}
