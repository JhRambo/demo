package main

// 客户端
//1.go rpc通信	默认
//2.jsonrpc通信 跨语言
import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

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

func main() {
	//1.建立连接
	// conn, err1 := rpc.Dial("tcp", "127.0.0.1:8866")	//1.rpc
	conn, err1 := net.Dial("tcp", "127.0.0.1:8866") //2.jsonrpc
	if err1 != nil {
		fmt.Println(err1)
	}
	//2.延迟关闭连接
	defer conn.Close()
	// 3.调用远程函数
	// =====添加数据
	// var addData = AddUserData{
	// 	Name: "张三",
	// 	Age:  20,
	// 	Hobby: []string{
	// 		"吃饭", "睡觉",
	// 	},
	// 	Sex:  "男",
	// 	addr: "深圳龙岗",
	// }
	// =====获取数据
	var getData = GetUserData{
		Id: 10,
	}
	// =====字符串
	// var res string
	// err2 := conn.Call("user.SayHello", "我是客户端", &res)
	// ======字符串
	// 响应值
	var res Res
	// err2 := conn.Call("user.AddUserData", addData, &res)
	// err2 := conn.Call("user.GetUserData", getData, &res)	//1.rpc
	// =======2.jsonrpc========
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err2 := client.Call("user.GetUserData", getData, &res)
	// =======2.jsonrpc========
	if err2 != nil {
		fmt.Println(err2)
	}
	//4.获取服务端返回的值
	fmt.Printf("%#v", res)
}
