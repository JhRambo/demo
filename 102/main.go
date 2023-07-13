package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	// 创建Consul客户端配置
	config := api.DefaultConfig()
	config.Address = "http://192.168.10.103:38500" // 设置Consul服务器地址和端口
	config.Token = "123456"
	// 创建Consul客户端
	client, err := api.NewClient(config)
	if err != nil {
		// 错误处理
		fmt.Println(err)
	}

	// 构建Service注册信息
	service := &api.AgentServiceRegistration{
		ID:      "my-service-1",
		Name:    "my-service",
		Address: "192.168.10.103",
		Port:    38902,
		Check: &api.AgentServiceCheck{
			TCP:      "192.168.10.103:38902", //http,tcp,grpc等
			Interval: "30s",
			Timeout:  "60s",
		},
	}

	// 服务注册
	err = client.Agent().ServiceRegister(service)
	if err != nil {
		// 错误处理
		fmt.Println(err)
	}
}
