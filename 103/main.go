package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/hashicorp/consul/api"
)

func main() {
	// 创建 Consul 客户端配置
	config := api.DefaultConfig()
	config.Address = "http://192.168.10.103:38500" // 设置Consul服务器地址和端口
	config.Token = "123456"
	// 创建 Consul 客户端
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 要查询的服务名称
	serviceName := "my-service"

	// 查询健康的服务
	checks, _, err := client.Health().Service(serviceName, "", true, &api.QueryOptions{})
	if err != nil {
		// 处理错误
		log.Fatal(err)
	}
	// 筛选健康的服务实例
	var healthyServices []*api.ServiceEntry
	for _, check := range checks {
		if len(check.Checks) > 0 {
			if check.Checks.AggregatedStatus() == api.HealthPassing {
				healthyServices = append(healthyServices, check)
			}
		}
	}

	// 随机选择一个健康的服务
	if len(healthyServices) > 0 {
		service := healthyServices[rand.Intn(len(healthyServices))]
		// 在这里执行业务操作，使用选中的健康服务
		fmt.Printf("Selected service: %s, Address: %s:%d\n", service.Service.ID, service.Service.Address, service.Service.Port)
	} else {
		fmt.Println("No healthy services found for the specified service")
	}
}
