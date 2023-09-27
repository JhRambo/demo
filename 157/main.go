package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
)

const CONSUL_CONFIG_KEY = "alarm-robot"

var (
	clientConnCache sync.Map
)

func main() {
	// 创建Consul客户端
	config := api.DefaultConfig()
	config.Address = "http://192.168.10.103:38500" // 设置Consul服务器地址和端口
	config.Token = "123456"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// 启动定期查询键值对的协程
	go watchKeyChanges(client)

	// 阻塞主线程以保持程序运行
	select {}
}

// 定期查询键值对的函数
func watchKeyChanges(client *api.Client) {
	for {
		// 查询键值对
		key, _, err := client.KV().Get(CONSUL_CONFIG_KEY, nil)
		if err != nil {
			log.Println(err)
			continue
		}
		// 处理键值对的变化
		handleKeyChange(key)
		// 休眠一段时间后再次查询
		time.Sleep(5 * time.Second)
	}
}

// 处理键值对变化的函数
func handleKeyChange(key *api.KVPair) {
	if key != nil {
		v, ok := clientConnCache.Load(CONSUL_CONFIG_KEY)
		if ok {
			if v != string(key.Value) {
				clientConnCache.Store(CONSUL_CONFIG_KEY, string(key.Value))
				log.Println("数据不一致")
				ExitService()
			}
			return
		}
		clientConnCache.Store(CONSUL_CONFIG_KEY, string(key.Value))
	}
}

// 重启服务
func RestartService() {
	// 重启Docker服务命令
	cmd := exec.Command("docker", "restart", "com.ghs.alarm")
	// 执行命令并等待完成
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	// 检查命令的退出状态
	if cmd.ProcessState.ExitCode() == 0 {
		log.Println("alarm服务重启成功")
	} else {
		log.Println("alarm服务重启失败")
	}
}

// docker容器运行的服务，直接退出即可，创建容器的时候设置了restart
func ExitService() {
	os.Exit(0)
}
