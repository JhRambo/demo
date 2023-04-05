package main

import (
	"context"
	"demo/grpc/proto/user"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// 定义User结构体
type User struct {
}

// 实现获取数据接口
func (this User) GetUser(c context.Context, g *user.GetUserData) (*user.Res, error) {
	fmt.Printf("%#v", g)
	return &user.Res{
		Success: true,
		Message: "获取成功",
	}, nil
}

// 实现添加数据接口
func (this User) AddUser(c context.Context, a *user.AddUserData) (*user.Res, error) {
	fmt.Printf("%#v", a)
	return &user.Res{
		Success: true,
		Message: "添加成功",
	}, nil
}

func main() {
	// openssl证书	TODO
	// credentials.NewServerTLSFromFile()
	// 1.初始化grpc对象
	grpcServer := grpc.NewServer()
	// 2.注册grpc服务
	user.RegisterUserServer(grpcServer, &User{})
	// 3.监听，指定IP，端口
	listener, err := net.Listen("tcp", ":8088") //默认127.0.0.1
	if err != nil {
		fmt.Println(err)
	}
	// 4.关闭服务
	defer listener.Close()
	// 5.启动服务
	fmt.Println("服务启动中...")
	grpcServer.Serve(listener)
}
