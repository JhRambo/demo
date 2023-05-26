package main

import (
	"fmt"
	"net/http"
	"time"
)

// 空结构体
type myHandler struct{}

func (my *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello  myHandler  ")
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tStart := time.Now()
		next.ServeHTTP(w, r)
		tEnd := time.Since(tStart)
		fmt.Println("middleware:", tEnd)
	})
}

func main() {
	mux := http.NewServeMux() //路由器
	myh := &myHandler{}
	mux.Handle("/h1", middleware(myh))
	if err := http.ListenAndServe(":8089", mux); err != nil {
		panic(err)
	}
}
