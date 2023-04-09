package main

import (
	"context"
	proto "demo/consul/proto/hello"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"demo/consul/config"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct{}

var (
	Port *int
	IP   *string
)

func (s Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: fmt.Sprintf("hello %s from %s:%d", request.Name, *IP, *Port)}, nil
}

func main() {
	// 这里进行服务注册
	// 设置本地服务启动的ip和端口号
	// 默认ip如果改成127.0.0.1会导致consul健康检查失败
	IP = flag.String("ip", "127.0.0.1", "ip地址")
	Port = flag.Int("port", 8081, "端口号")

	flag.Parse()

	g := grpc.NewServer()

	proto.RegisterGreeterServer(g, &Server{})
	tcpAddr := fmt.Sprintf("%s:%d", *IP, *Port)
	log.Printf("service listen:%s", tcpAddr)
	lis, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		log.Panicln("failed to listen:" + err.Error())
	}
	// 注册健康检查
	grpc_health_v1.RegisterHealthServer(g, health.NewServer())
	// 将当前grpc服务注册到consul
	serviceId := uuid.NewV4().String()
	client, err := RegisterGRPCService(*IP, config.ServiceName, serviceId, *Port, nil)
	if err != nil {
		log.Panicln("grpc服务注册失败：", err)
	}

	// 将启动服务的部分放到协程里面，使得后面监听终止信号的部分可以被执行
	go func() {
		err = g.Serve(lis)
		if err != nil {
			log.Panicln("failed to start grpc:" + err.Error())
		}
	}()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceId); err != nil {
		log.Println("服务注销失败")
	} else {
		log.Println("服务注销成功")
	}
}

// 函数功能：将grpc服务注册到consul
// 参数说明
// address：待注册的服务的ip
// name：服务名称
// id：服务id
// port：服务端口
// tags：服务标签
func RegisterGRPCService(address, name, id string, port int, tags []string) (*api.Client, error) {
	cfg := api.DefaultConfig()
	// 设置consul服务运行所在的ip和端口
	//cfg.Address的ip可以是127.0.0.1
	cfg.Address = fmt.Sprintf("%s:%d", config.ConsulIp, config.ConsulPort)

	client, err := api.NewClient(cfg)
	if err != nil {
		log.Panic(err)
	}

	// 生成健康检查对象
	check := &api.AgentServiceCheck{
		// 这里的ip不可以是127.0.0.1
		GRPC:                           fmt.Sprintf("%s:%d", address, port), // 服务的运行地址
		Timeout:                        "5s",                                // 超过此时间说明服务状态不健康
		Interval:                       "5s",                                // 每5s检查一次
		DeregisterCriticalServiceAfter: "30s",                               // 失败多久后注销服务
	}

	// 生成注册对象
	registration := &api.AgentServiceRegistration{
		Name:    name,
		ID:      id,
		Address: address,
		Port:    port,
		Tags:    tags,
		Check:   check,
	}

	// 注册服务
	return client, client.Agent().ServiceRegister(registration)
}
