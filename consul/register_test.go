package consul

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/consul/api"
)

// 设置key/Value
func TestSetKV(t *testing.T) {
	kv := client.KV()
	p := &api.KVPair{
		Key: consulConfig,
		Value: []byte(`{
			"name":"golang",
			"age":"99"
		}`),
		Flags: 32,
	}
	_, err := kv.Put(p, nil)
	if err != nil {
		log.Fatalln(err)
	}
	pair, _, err := kv.Get(consulConfig, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}

// 测试服务注册
func TestServerRegister(t *testing.T) {
	ServerRegister(serverName, serverIp, serverPort)
	// for true {
	// 	// 制定一个定时器，模拟30s后注销服务
	// 	ticker := time.NewTicker(ServerCancelTime * time.Second)
	// 	select {
	// 	case tt := <-ticker.C:
	// 		ServerCancel(svrID) //服务注销
	// 		log.Printf("服务注销，tt:%v\n", tt)
	// 		return
	// 	}
	// }
}
