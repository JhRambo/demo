package main

// 服务端
//1.go rpc通信	默认，不能跨语言
//2.jsonrpc通信 可跨语言通信
import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 定义User结构体
type User struct {
}

// 获取用户结构体参数
type GetUserData struct {
	Id int
}

// 添加用户结构体参数
type AddUserData struct {
	Name  string
	Age   int
	Hobby []string
	Sex   string
	addr  string
}

// 接口响应结构体参数
type Res struct {
	Success bool
	Message string
}

// 定义User结构体对应的函数
func (this User) SayHello(req string, res *string) error {
	fmt.Printf("%#v", req)
	*res = "你好 " + req
	return nil
}

// 获取数据
func (this User) GetUserData(req GetUserData, res *Res) error {
	fmt.Printf("%#v", req)
	*res = Res{
		Success: true,
		Message: "获取成功",
	}
	return nil
}

// 添加数据
func (this User) AddUserData(req AddUserData, res *Res) error {
	fmt.Printf("%#v", req)
	*res = Res{
		Success: true,
		Message: "添加成功",
	}
	return nil
}

func main() {
	// 1、注册服务
	// err1 := rpc.RegisterName("user", new(User))	//1 rpc
	err1 := rpc.RegisterName("user", new(User)) //2 jsonrpc
	if err1 != nil {
		fmt.Println(err1)
	}
	// 2、监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8866")
	if err2 != nil {
		fmt.Println(err2)
	}
	//3、应用退出时关闭监听
	defer listener.Close()

	for {
		// 4、建立连接
		fmt.Println("建立连接中...")
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		// 5、绑定服务
		// rpc.ServeConn(conn)	//1.rpc 默认
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) //2.jsonrpc
	}
}
