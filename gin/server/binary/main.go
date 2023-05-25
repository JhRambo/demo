package main

import (
	"demo/gin/config"
	pb "demo/gin/proto/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBinaryHttpServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) UploadFile(stream pb.BinaryHttp_UploadFileServer) error {
	var fileData []byte
	for {
		// 从客户端流中接收数据
		chunk, err := stream.Recv()
		if err == io.EOF { //数据传输结束
			stream.SendAndClose(&empty.Empty{})
			break
		}
		if err != nil {
			return err
		}
		fileData = append(fileData, chunk.Data...)
	}
	log.Println("fileData============", fileData)
	return nil
}

// grpc-server
func main() {
	log.Println("GRPC-SERVER on http://0.0.0.0:8081")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.SERVER_PORT))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建一个gRPC-server服务
	s := grpc.NewServer()
	// 注册gRPC-server服务
	pb.RegisterBinaryHttpServer(s, NewServer())
	// 启动gRPC-Server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
