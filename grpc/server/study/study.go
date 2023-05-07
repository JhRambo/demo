package study

import (
	"context"
	pb "demo/grpc/proto/study"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

// 定义结构体，在调用注册api的时候作为入参，
// 该结构体会带上SayStudy方法，里面是业务代码
// 这样远程调用时就执行了业务代码了
type server struct {
	// pb.go中自动生成的，是个空结构体
	pb.UnimplementedStudyHttpServer
}

// 实际处理业务逻辑的地方
// 业务代码在这里写，客户端远程tcp协议调用SayStudy
// 会执行这里的代码
func (s *server) SayStudy(ctx context.Context, in *pb.StudyRequest) (*pb.StudyResponse, error) {
	var infos []*pb.Info
	var subs []*pb.SubInfo
	sub1 := &pb.SubInfo{
		Date: "2023-05-05",
		Time: "01:02:03",
	}
	sub2 := &pb.SubInfo{
		Date: "2023-05-05",
		Time: "01:02:03",
	}
	subs = append(subs, sub1, sub2)
	info := &pb.Info{
		DevId: subs,
	}
	infos = append(infos, info)
	res := &pb.StudyResponse{
		Code:    200,
		Message: "成功",
		List:    infos,
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
	pb.RegisterStudyHttpServer(s, &server{})
	log.Println("开始监听，等待远程调用...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
