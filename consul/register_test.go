package consul

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/consul/api"
)

func TestClient(t *testing.T) {
	//设置key/value
	kv := client.KV()
	p := &api.KVPair{Key: consulConfig, Value: []byte("111111111111"), Flags: 32}
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
	svrID := ServerRegister(serverName, serverIp, serverPort)
	log.Printf("服务注册成功，服务ID：%v\n", svrID)
	for true {
		// 制定一个定时器，模拟30s后注销服务
		t := time.NewTicker(ServerCancelTime * time.Second)
		select {
		case tt := <-t.C:
			ServerCancel(svrID) //服务注销
			log.Printf("服务注销，tt:%v\n", tt)
			return
		}
	}
}
