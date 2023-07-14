package main

import (
	"fmt"
	"log"
	"strconv"

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
		log.Fatalln(err)
	}

	for i := 0; i < 3; i++ {
		// 构建Service注册信息
		port, _ := strconv.Atoi(fmt.Sprintf("808%d", i))
		service := &api.AgentServiceRegistration{
			ID:      fmt.Sprintf("my-service-%d", i),
			Name:    "my-service",
			Address: "192.168.10.103",
			Port:    port,
			Check: &api.AgentServiceCheck{
				HTTP:     fmt.Sprintf("http://192.168.10.103:%d/check/health", port), //http,tcp,grpc等
				Interval: "10s",
				Timeout:  "30s",
			},
		}

		// 服务注册
		err = client.Agent().ServiceRegister(service)
		if err != nil {
			// 错误处理
			log.Fatalln(err)
		}
	}

	// 构建Service注册信息 GRPC
	service := &api.AgentServiceRegistration{
		ID:      "my-service-4",
		Name:    "my-service",
		Address: "192.168.10.103",
		Port:    8084,
		Check: &api.AgentServiceCheck{
			GRPC:     "192.168.10.103:8084/check", //http,tcp,grpc等
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	// 服务注册
	err = client.Agent().ServiceRegister(service)
	if err != nil {
		// 错误处理
		log.Fatalln(err)
	}

	fmt.Println("success")
}
