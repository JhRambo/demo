package main

import (
	"context"
	"demo/gin/config"
	pb_binary "demo/gin/proto/binary"
	pb_hello "demo/gin/proto/hello"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb_binary.UnimplementedBinaryHttpServer
	pb_hello.UnimplementedHelloHttpServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) UploadFile(stream pb_binary.BinaryHttp_UploadFileServer) error {
	return nil
}

func (s *Server) SayHello(ctx context.Context, req *pb_hello.HelloHttpRequest) (*pb_hello.HelloHttpResponse, error) {
	resp := &pb_hello.HelloHttpResponse{
		Code: 200,
		Msg:  req.Name + " hello!",
	}
	return resp, nil
}

func (s *Server) SayGoodbye(ctx context.Context, req *pb_hello.GoodByeHttpRequest) (*pb_hello.GoodByeHttpResponse, error) {
	resp := &pb_hello.GoodByeHttpResponse{
		Code: 200,
		Msg:  req.Name + " goodbye!",
	}
	return resp, nil
}

// grpc-server
func main() {
	log.Println("GRPC-SERVER on http://0.0.0.0:8081")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.SERVER_PORT))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建gRPC-server服务
	s := grpc.NewServer()
	// 注册gRPC-server服务
	pb_binary.RegisterBinaryHttpServer(s, NewServer())
	pb_hello.RegisterHelloHttpServer(s, NewServer())
	// 启动gRPC-Server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
