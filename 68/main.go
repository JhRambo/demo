package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 每次请求，Handler会创建一个Goroutine为其提供服务；
// 如果连续请求3次，request的地址也是不同的：
func SayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(&request)
	go func() {
		for range time.Tick(time.Second) { //每隔1s打印
			//每次打印前检查context是否已经结束，如果结束则退出循环，即结束生命周期
			// context包可以提供一个请求从API请求边界到各goroutine的请求域数据传递、取消信号及截至时间等能力。
			select {
			case <-request.Context().Done():
				fmt.Println("request is outgoing")
				return
			default:
				fmt.Println("Current request is in progress") ////如果没有引入context，这里在2s后还会一直打印
			}
		}
	}()

	time.Sleep(2 * time.Second) //2s后执行下面的代码
	writer.Write([]byte("Hi, New Request Comes"))
}

func main() {
	http.HandleFunc("/", SayHello) // 设置访问的路由
	log.Fatalln(http.ListenAndServe(":8088", nil))
}
