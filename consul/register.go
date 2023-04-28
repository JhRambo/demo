package consul

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/hashicorp/consul/api"
)

// ServerRegister 服务注册	候提供一个健康检查地址，支持http，grpc
func ServerRegister(svrName string, ipAddr string, port int) string {
	svrID := RandomStr(32)
	// 设置Consul对服务健康检查的参数
	check := api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%v:%v/consul/health/?id=%v", consulIp, consulPort, svrID), // 健康检查地址，自定义ip和端口
		Interval:                       "3s",                                                                          // 健康检查频率
		Timeout:                        "2s",                                                                          // 健康检查超时
		Notes:                          "Consul 代码健康检查",
		DeregisterCriticalServiceAfter: "5s", // 健康检查失败30s后 consul自动将注册服务删除
		Name:                           "代码自定义检查svr1",
		// GRPC:                           fmt.Sprintf("%v:%v/%v", serverIp, serverPort, svrName),
	}
	// 设置微服务向Consul注册信息
	reg := &api.AgentServiceRegistration{
		ID:      svrID,                      // 服务节点的ID
		Name:    svrName,                    // 服务名称
		Address: ipAddr,                     // 服务IP
		Port:    port,                       // 服务端口
		Tags:    []string{"v1.1", "backup"}, // 标签，可在服务发现时筛选服务，类似v1.1
		Check:   &check,                     // 健康检查
	}
	// 执行注册
	if err := client.Agent().ServiceRegister(reg); err != nil {
		log.Fatalln(err)
	}
	return svrID
}

// ServerCancel 服务注销
func ServerCancel(svrID string) {
	// 执行服务注销
	err := client.Agent().ServiceDeregister(svrID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("注销服务成功==================", err)
}

//========如果使用grpc接口实现健康检查，则需要实现HealthServer 接口，服务启动时候注册这个pb==========

// HealthImpl 健康检查实现
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch 让HealthImpl实现RegisterHealthServer内部的interface接口
func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}
