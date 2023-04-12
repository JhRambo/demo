package main

import (
	"fmt"
	"net/http"
	"strings"
)

// 空结构体
type HelloHandler struct{}

// 需要手动实现ServeHTTP接口
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI // /hello1?id=1&name=xxxx
	index := strings.Index(uri, "hello1")
	fmt.Println("index=", index)
	if index > 0 {
		fmt.Fprintf(w, "Hello Handler!")
	} else {
		fmt.Fprintf(w, "error Handler!")
	}
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello HandleFunc!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8088",
	}
	helloHandler := &HelloHandler{}
	http.Handle("/hello1", helloHandler)
	http.Handle("/hello3", helloHandler)
	http.HandleFunc("/hello2", helloHandleFunc)
	server.ListenAndServe()
}
