package hello

import (
	"context"
	pb "demo/grpc/proto/hello"
)

type Server struct {
	pb.UnimplementedHelloHttpServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloHttpRequest) (*pb.HelloHttpResponse, error) {
	return &pb.HelloHttpResponse{
		Name: req.Name,
		Age:  req.Age,
	}, nil
}
