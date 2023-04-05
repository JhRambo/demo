package main

import (
	"context"
	"demo/grpc/proto/user"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// openssl证书 TODO
	// credentials.NewClientTLSFromFile()
	// 1.建立连接
	grpcClient, err := grpc.Dial("127.0.0.1:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	// 2.注册客户端
	client := user.NewUserClient(grpcClient)
	res, err := client.AddUser(context.Background(), &user.AddUserData{Name: "张三", Age: 20, Addr: "厦门", Sex: "男", Hobby: []string{"吃饭", "睡觉"}}) //添加接口
	// res, err := client.GetUser(context.Background(), &user.GetUserData{Id: 10}) //查询接口
	fmt.Printf("%#v", res)
}
