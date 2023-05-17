package main

import (
	"context"
	"log"

	pb "demo/grpc/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// openssl证书
	// credentials.NewClientTLSFromFile()
	// 1.建立连接
	// 这里是通过tcp协议调用
	grpcClient, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	// 2.注册客户端
	client := pb.NewHelloHttpClient(grpcClient)
	ctx := context.Background()
	//客户端自定义metadata数据
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("token", "1234567890"))
	res, err := client.SayHello(ctx, &pb.HelloHttpRequest{
		Name: "张三李四王五老六",
		Age:  100,
	})
	log.Println(res)
}
