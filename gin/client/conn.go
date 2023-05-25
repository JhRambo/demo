package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"demo/gin/config"

	"google.golang.org/grpc"
)

// ClientConnMap keep connections
var ClientConnMap = make(map[string]*grpc.ClientConn)

// InitGRPCClients read grpc configuration and init clients
func InitGRPCClients() error {
	config := config.Grpcserver
	for name, addr := range config {
		_, ok := ClientConnMap[name]
		if ok {
			continue
		}
		conn, err := grpc.Dial(
			addr,
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithTimeout(3*time.Second),
		)
		if err != nil {
			log.Printf("Failed to connect to gRPC service [%s] on [%s]: %v", name, addr, err)
			continue
		}

		ClientConnMap[name] = conn
		log.Printf("Connected to gRPC service [%s] on [%s]", name, addr)
	}
	return nil
}

// GetGRPCClient get grpc client by name
func GetGRPCClient(ctx context.Context, name string) (*grpc.ClientConn, error) {
	cc, ok := ClientConnMap[name]
	if !ok {
		return nil, fmt.Errorf("not found grpc client [%s]", name)
	}
	return cc, nil
}
