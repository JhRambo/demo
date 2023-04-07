package main

import (
	"context"
	"fmt"

	"demo/grpc/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// openssl证书 TODO
	// credentials.NewClientTLSFromFile()
	// 1.建立连接
	grpcClient, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	// 2.注册客户端
	client := hello.NewGreeterClient(grpcClient)
	res, err := client.SayHello(context.Background(), &hello.HelloRequest{Name: "张三"})
	fmt.Printf("%#v", res)
}
