package consul

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
)

const (
	consulConfig     = "mysvr-config"
	consulIp         = "192.168.10.103"
	consulPort       = 38500
	consulToken      = "123456"
	serverIp         = "127.0.0.1"
	serverPort       = 9099
	serverName       = "mysvr-grpc" //Services服务名称
	ServerCancelTime = 5
)

var client *api.Client

func init() {
	// 创建Consul客户端连接
	cl, err := api.NewClient(&api.Config{
		Address: fmt.Sprintf("http://%v:%v", consulIp, consulPort),
		Token:   consulToken,
	})
	if err != nil {
		log.Fatalf("client 创建失败，退出:%v\n", err)
	}
	client = cl
}

// 生成随机字符串
func RandomStr(len int) string {
	nUid := uuid.NewV4().String()
	str := strings.Replace(nUid, "-", "", -1)
	if len < 0 || len >= 32 {
		return str
	}
	return str[:len]
}
