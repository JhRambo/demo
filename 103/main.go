package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/hashicorp/consul/api"
)

func main() {
	// 启动定时器
	runTimer()
}

func runTimer() {
	count := 0
	// 创建定时器，每隔指定时间检查一次服务健康状态
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// 监听定时器触发事件
	for range ticker.C {
		// 检查服务健康状态
		isHealthy := checkHealth()
		if !isHealthy {
			count++
			fmt.Println("监测服务不健康次数：", count)
			if count >= 3 {
				fmt.Println("Server is not healthy. Restarting...")
				// 重新启动服务器
				if err := startServer(); err != nil {
					fmt.Println(err)
				}
				count = 0
			}
		}
	}
}

func checkHealth() bool {
	// 创建 Consul 客户端配置
	config := api.DefaultConfig()
	config.Address = "http://192.168.10.103:38500" // 设置Consul服务器地址和端口
	config.Token = "123456"
	// 创建 Consul 客户端
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	// 要查询的服务名称
	serviceName := "gateway"
	// 查询健康的服务
	checks, _, err := client.Health().Service(serviceName, "", true, &api.QueryOptions{})
	if err != nil {
		// 处理错误
		panic(err)
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
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 随机选择一个健康的服务
	if len(healthyServices) > 0 {
		service := healthyServices[rand.Intn(len(healthyServices))]
		// 在这里执行业务操作，使用选中的健康服务
		fmt.Printf("Selected service: %s, Address: %s:%d\n", service.Service.ID, service.Service.Address, service.Service.Port)
		return true
	} else {
		fmt.Println("No healthy services found for the specified service")
		return false
	}
}

func startServer() error {
	err := os.Chdir("D:/code/Starverse/")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// 执行 docker-compose -f docker-compose-gateway-win.yml up -d --build
	cmd := exec.Command("docker-compose", "-f", "docker-compose-gateway-win.yml", "up", "-d", "--build")
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf("Output: %s\n", output)
	return nil
}
