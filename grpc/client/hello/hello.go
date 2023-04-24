package main

import (
	"context"
	"fmt"

	"demo/grpc/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// openssl证书
	// credentials.NewClientTLSFromFile()
	// 1.建立连接
	// 这里是通过tcp协议调用
	grpcClient, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	// 2.注册客户端
	client := hello.NewHelloHttpClient(grpcClient)
	ctx := context.Background()
	//客户端自定义metadata数据
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("token", "1234567890"))
	res, err := client.SayHello(ctx, &hello.HelloRequest{Name: "张三疯"})
	fmt.Printf("%#v", res)
}
