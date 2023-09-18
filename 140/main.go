package main

import (
	"fmt"
	"log"
	"sync"
	"time"

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

	// 创建 KV API
	kv := client.KV()

	// 设置要监控的键名
	key := "alarm-robot"

	// 创建一个 WaitGroup
	var wg sync.WaitGroup

	wg.Add(1)
	// 启动协程进行监控
	go func() {
		// 在任务结束时调用 Done 方法
		defer wg.Done()

		for {
			// 获取键/值
			pair, _, err := kv.Get(key, nil)
			if err != nil {
				log.Printf("Failed to get key: %s\n", err)
				continue
			}

			if pair != nil {
				fmt.Printf("Key: %s, Value: %s\n", pair.Key, pair.Value)
			} else {
				fmt.Printf("Key not found: %s\n", key)
			}

			// 等待一段时间后再次查询
			time.Sleep(5 * time.Second)
		}
	}()

	// 阻塞主线程
	// select {}
	// for {
	// }
	wg.Wait()
}
