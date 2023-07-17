package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthImpl struct{}

// Check 实现健康检查接口
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch 实现监视接口，不需要特定逻辑则直接返回 nil
func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

func main() {
	// 创建 gRPC 服务器
	server := grpc.NewServer()

	// 创建健康检查服务实现对象
	healthImpl := &HealthImpl{}

	// 注册健康检查服务实现对象到 gRPC 服务器
	grpc_health_v1.RegisterHealthServer(server, healthImpl)

	// 启动 gRPC 服务器
	address := ":8084" // 设置监听地址
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on %s...", address)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
