package main

import (
	"fmt"
	"net/http"
	"time"
)

// 空结构体
type myHandler1 struct{}

func (my *myHandler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello  myHandler1  ")
}

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// http.Error(w, http.StatusText(403), 403)
		tStart := time.Now()
		next.ServeHTTP(w, r)
		tEnd := time.Since(tStart)
		fmt.Println("middleware1:", tEnd)
	})
}

func main() {
	mux := http.NewServeMux()
	myh := &myHandler1{}
	mux.Handle("/h1", middleware1(myh))
	if err := http.ListenAndServe(":8088", mux); err != nil {
		panic(err)
	}
}
