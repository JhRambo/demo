package main

import (
	"context"
	pb "demo/grpc/proto/hello"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

// 定义结构体，在调用注册api的时候作为入参，
// 该结构体会带上SayHello方法，里面是业务代码
// 这样远程调用时就执行了业务代码了
type server struct {
	// pb.go中自动生成的，是个空结构体
	pb.UnimplementedHelloHttpServer
}

// 定义切片[]类型的map：
// var ms = make(map[string][]map[string]string)
// uu := []map[string]string{}
//
//	u1 := map[string]string{
//		"date": "2023-05-07",
//		"time": "07:07:07",
//	}
//
//	u2 := map[string]string{
//		"date": "2023-05-07",
//		"time": "07:07:07",
//	}
//
// uu = append(uu, u1, u2)
// ms["dev1"] = uu
// ms["dev2"] = uu
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	var ms = make(map[string]*pb.Info)
	dd := []*pb.SubInfo{}
	d := &pb.SubInfo{
		D2: map[string]string{
			"devId": fmt.Sprintf("dev_%d", 1),
			"date":  "2023-05-05",
			"time":  "01:01:01",
		},
	}
	dd = append(dd, d)

	for i := 0; i < 3; i++ {
		d := &pb.SubInfo{
			D2: map[string]string{
				"devId": fmt.Sprintf("dev_%d", i),
				"date":  "2023-05-05",
				"time":  "01:01:01",
			},
		}
		dd = append(dd, d)
	}
	uu := &pb.Info{
		D1: dd,
	}
	ms["data"] = uu

	res := &pb.HelloResponse{
		Code:    200,
		Message: "成功",
		List:    ms,
	}
	return res, nil
}

func main() {
	// 要监听的协议和端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化gRPC server结构体
	s := grpc.NewServer()
	// 服务注册
	pb.RegisterHelloHttpServer(s, &server{})
	log.Println("开始监听，等待远程调用...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
