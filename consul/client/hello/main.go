package main

import (
	"context"
	"demo/consul/common"
	"demo/consul/config"
	proto "demo/consul/proto/hello"
	"fmt"
	"log"
	"time"
)

func main() {
	var srvClient proto.GreeterClient
	// 服务
	if conn := common.InitSrvConn(config.ServiceName); conn != nil {
		srvClient = proto.NewGreeterClient(conn)
	}
	req := &proto.HelloRequest{
		Name: "",
	}

	for i := 0; i < 1000; i++ {
		req.Name = fmt.Sprintf("james%d", i)
		resp, err := srvClient.SayHello(context.Background(), req)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(resp.Message)
		time.Sleep(time.Second)
	}
}
