package main

import (
	"log"
	"net/http"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "my_exporter"
	subsystem = "consul_health"
)

var (
	serviceHealthMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "service_health",
			Help:      "Health status of services in Consul",
		},
		[]string{"service"},
	)
)

func init() {
	prometheus.MustRegister(serviceHealthMetric)
}

func main() {
	// 创建Consul客户端配置
	config := consulapi.DefaultConfig()
	config.Address = "http://192.168.10.103:38500" // 设置Consul服务器地址和端口
	config.Token = "123456"
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("Failed to create Consul client:", err)
	}

	// 创建HTTP路由和处理器
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	// 循环获取服务健康状态并更新指标
	for {
		checks, _, err := client.Health().State("any", &consulapi.QueryOptions{})
		if err != nil {
			log.Println("Failed to retrieve service health:", err)
			continue
		}

		for _, check := range checks {
			serviceHealthMetric.With(prometheus.Labels{"service": check.ServiceID}).Set(toFloat64(check.Status))
		}

		time.Sleep(30 * time.Second) // 间隔时间可以根据需要进行调整
	}
}

func toFloat64(status string) float64 {
	if status == consulapi.HealthPassing {
		return 1.0
	}
	return 0.0
}
