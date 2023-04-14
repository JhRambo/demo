package main

import (
	"fmt"
	"net/http"
)

// 空结构体
type HelloHandler struct{}

// Handle需要手动实现ServeHTTP接口
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Handler!")
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello HandleFunc!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8088",
	}
	helloHandler := &HelloHandler{}
	//hello1和hello3都是调用HelloHandler.ServeHTTP这个方法
	http.Handle("/hello1", helloHandler)
	http.Handle("/hello3", helloHandler)
	http.HandleFunc("/hello2", helloHandleFunc)
	server.ListenAndServe()
}
