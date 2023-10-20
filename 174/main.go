package grpc

import (
	"net"
	"net/http"
	"time"

	"com.ghs.utils/logs"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus" // 引入 grpc-prometheus 库
	"google.golang.org/grpc/keepalive"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

const maxSize = 100 * 1024 * 1024

var opts = []grpc.ServerOption{
	grpc_middleware.WithUnaryServerChain(
		RecoveryInterceptor,
	),
	grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute, // 这个连接最大的空闲时间，超过就释放，解决proxy等到网络问题（不通知 grpc 的 client 和 server）
	}),
	grpc.MaxRecvMsgSize(maxSize),
	grpc.MaxSendMsgSize(maxSize),
}

var grpcServer = grpc.NewServer(opts...)

func GetGrpc() *grpc.Server {
	return grpcServer
}

func Run(addr string, errc chan error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}

	defer func() {
		if err1 := recover(); err1 != nil {
			logs.Debug(err1)
		}
	}()

	grpc_prometheus.Register(grpcServer) // 注册 Prometheus 指标收集器到 gRPC 服务器

	http.Handle("/metrics", promhttp.Handler()) // 添加 Prometheus HTTP handler

	go func() {
		err := http.ListenAndServe(":9090", nil) // 启动 HTTP 服务器
		if err != nil {
			logs.Errorf("HTTP server failed: %v", err)
		}
	}()

	errc <- grpcServer.Serve(lis)
}
