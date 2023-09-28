package main

import (
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
)

func main() {
	// 创建Consul客户端
	config := api.DefaultConfig()
	config.Address = "http://192.168.10.103:38500" // 设置Consul服务器地址和端口
	config.Token = "123456"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	var index uint64 = 0
	// 创建阻塞查询参数
	params := &api.QueryOptions{
		WaitIndex: index,
		WaitTime:  time.Duration(5) * time.Second,
	}

	// 阻塞查询 key/value 变化
	for {
		keyPair, meta, err := client.KV().Get("alarm-robot", params)
		if err != nil {
			fmt.Println(err)
			return
		}
		if keyPair != nil && meta.LastIndex != index {
			// 处理变化
			fmt.Printf("value: %s\n", keyPair.Value)
			index = meta.LastIndex
			params.WaitIndex = index
		}
	}
}
