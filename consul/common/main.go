package common

import (
	"demo/consul/config"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 同时完成了服务发现和负载均衡算法（轮询）
func InitSrvConn(srvName string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?healthy=true&wait=14s",
			config.ConsulIp, config.ConsulPort, srvName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc目前似乎只支持轮询的负载均衡算法
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Println("service层服务发现出错: ", err)
		return nil
	}
	return conn
}
