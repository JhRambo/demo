package main

import (
	"context"
	"log"

	pb "demo/grpc/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const CTXINFO = "ctx_info"

var MapNode = map[string]string{}

// 设置自定义metadata
func SetMetaData() metadata.MD {
	md := metadata.Pairs(CTXINFO, MapNode[CTXINFO])
	return md
}

func main() {
	// 1.建立连接
	grpcClient, err := grpc.Dial(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	// 2.注册grpc客户端
	client := pb.NewHelloHttpClient(grpcClient)
	ctx := context.Background()
	//客户端自定义metadata数据
	ctx = metadata.NewOutgoingContext(ctx, SetMetaData())
	res, err := client.SayHello(ctx, &pb.HelloHttpRequest{
		Name: "张三李四王五老六",
		Age:  100,
	})
	log.Println(res)
}
