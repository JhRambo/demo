package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type handle struct {
	host string
	port string
}

// 服务种类
type Service struct {
	auth *handle
	user *handle
}

// 反向代理，路由转发
func (this *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var remote *url.URL
	if strings.Contains(r.RequestURI, "api/auth") {
		remote, _ = url.Parse("http://" + this.auth.host + ":" + this.auth.port)
	} else if strings.Contains(r.RequestURI, "api/user") {
		remote, _ = url.Parse("http://" + this.user.host + ":" + this.user.port)
	} else {
		fmt.Fprintf(w, "404 Not Found")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}
func startServer() {
	log.Println("服务启动中...")
	// 注册被代理的服务器 (host， port)
	service := &Service{
		auth: &handle{host: "127.0.0.1", port: "8081"},
		user: &handle{host: "127.0.0.1", port: "8082"},
	}
	err := http.ListenAndServe(":8088", service)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
func main() {
	startServer()
}
