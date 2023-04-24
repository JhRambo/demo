package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------------")
	fmt.Println(r.Proto)      //协议	HTTP/1.1
	fmt.Println(r.TLS)        //<nil>
	fmt.Println(r.Host)       //127.0.0.1:8088
	fmt.Println(r.RequestURI) // /index?id=1&name=xxxx
	url := GetURL(r)
	fmt.Println(url) //http://127.0.0.1:8088/index?id=1&name=xxxx
}

func GetURL(r *http.Request) (Url string) {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, r.Host, r.RequestURI}, "")
}

func main() {
	fmt.Println("服务启动中...")
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
