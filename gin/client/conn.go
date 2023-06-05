package client

import (
	"fmt"
	"log"
	"time"

	"demo/gin/config"

	"google.golang.org/grpc"
)

var ClientConnMap = make(map[string]*grpc.ClientConn)

// 初始化grpc-client连接
func InitGRPCClients() error {
	config := config.GRPCserver
	for name, addr := range config {
		_, ok := ClientConnMap[name]
		if ok {
			continue
		}
		conn, err := grpc.Dial(
			addr,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithTimeout(60*time.Second),
		)
		if err != nil {
			log.Printf("Failed to connect to gRPC service [%s] on [%s]: %v", name, addr, err)
			panic(err)
		}

		ClientConnMap[name] = conn
		log.Printf("Connected to gRPC service [%s] on [%s]", name, addr)
	}
	return nil
}

func GetGRPCClient(name string) (*grpc.ClientConn, error) {
	cc, ok := ClientConnMap[name]
	if !ok {
		return nil, fmt.Errorf("not found grpc client [%s]", name)
	}
	return cc, nil
}
