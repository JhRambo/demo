package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_gateway_requests_total",
			Help: "Total number of requests processed",
		},
		[]string{"method", "status_code"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_gateway_request_duration_seconds",
			Help:    "Request duration in seconds",
			Buckets: []float64{0.1, 0.5, 1, 2, 5},
		},
		[]string{"method", "status_code"},
	)
)

func main() {
	// 注册自定义指标
	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":8088", nil))
	}()

	// Wait indefinitely
	select {}

	// API 网关的其他业务逻辑
	// ...

	// 示例：记录请求的指标
	start := time.Now()
	processRequest() // 处理请求的函数
	duration := time.Since(start)

	// 更新请求计数和持续时间指标
	requestCount.WithLabelValues("GET", "200").Inc()
	requestDuration.WithLabelValues("GET", "200").Observe(duration.Seconds())
}

func processRequest() {
	// 解析请求参数
	// ...

	// 执行业务逻辑
	// ...

	// 构造响应
	// ...

	// 发送响应给客户端
	// ...
}
