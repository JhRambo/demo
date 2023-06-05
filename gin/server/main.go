package main

import (
	"context"
	"demo/gin/config"
	pb_account "demo/gin/utils/proto/account"
	pb_binary "demo/gin/utils/proto/binary"
	pb_hello "demo/gin/utils/proto/hello"
	pb_msgpack "demo/gin/utils/proto/msgpack"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/grpc"
)

type Server struct {
	pb_binary.UnimplementedBinaryHttpServer
	pb_hello.UnimplementedHelloHttpServer
	pb_msgpack.UnimplementedMsgpackHttpServer
	pb_account.UnimplementedAccountHttpServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) AccountInfo(ctx context.Context, req *pb_account.AccountInfoRequest) (*pb_account.AccountInfoResponse, error) {
	resp := &pb_account.AccountInfoResponse{
		Code:    0,
		Message: " AccountInfo!",
		Info:    &pb_account.AccountInfo{},
	}
	return resp, nil
}

// 使用msgpack协议
func (s *Server) MsgPackProtocol(ctx context.Context, req *pb_msgpack.MsgpackHttpRequest) (*pb_msgpack.MsgpackHttpResponse, error) {
	var bys []byte
	//根据不同key跳转到不同的服务去执行
	switch req.Key {
	case "/HELLO/SAYHELLO": //统一转大写处理
		r := &pb_hello.HelloHttpRequest{}
		msgpack.Unmarshal(req.Val, r)
		w, _ := s.SayHello(ctx, r)
		bys, _ = msgpack.Marshal(w)
		//这里追加其他的case=====================TODO
	}
	return &pb_msgpack.MsgpackHttpResponse{
		Data: bys,
	}, nil
}

func (s *Server) UploadFile(stream pb_binary.BinaryHttp_UploadFileServer) error {
	var fileData []byte
	for {
		// 从流中读取文件数据
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fileData = append(fileData, chunk.Data...)
	}
	stream.SendAndClose(&pb_binary.BinaryHttpResponse{
		Code:    0,
		Message: "ok",
	})
	return nil
}

func (s *Server) SayHello(ctx context.Context, req *pb_hello.HelloHttpRequest) (*pb_hello.HelloHttpResponse, error) {
	resp := &pb_hello.HelloHttpResponse{
		Code:    0,
		Message: req.Name + " hello!",
	}
	return resp, nil
}

func (s *Server) SayGoodbye(ctx context.Context, req *pb_hello.GoodByeHttpRequest) (*pb_hello.GoodByeHttpResponse, error) {
	resp := &pb_hello.GoodByeHttpResponse{
		Code: 0,
	}
	if req.Name == "" {
		resp.Code = -101
		resp.Message = "参数不能为空！"
		return resp, nil
	}
	resp.Message = req.Name + " goodbye!"
	return resp, nil
}

// grpc-server
func main() {
	log.Println("GRPC-SERVER on http://0.0.0.0:8081")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.SERVER_PORT1))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// 创建gRPC-server服务
	s := grpc.NewServer()
	// 注册gRPC-server服务
	pb_binary.RegisterBinaryHttpServer(s, NewServer())
	pb_hello.RegisterHelloHttpServer(s, NewServer())
	pb_msgpack.RegisterMsgpackHttpServer(s, NewServer())
	pb_account.RegisterAccountHttpServer(s, NewServer())
	// 启动gRPC-Server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
