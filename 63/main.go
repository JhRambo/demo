package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

// 需要手动实现ServeHTTP接口
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Handler!")
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello HandleFunc!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8088",
	}
	helloHandler := HelloHandler{}
	http.Handle("/hello1", helloHandler)
	http.HandleFunc("/hello2", helloHandleFunc)
	server.ListenAndServe()
}
