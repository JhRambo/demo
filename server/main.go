package main

//开启http服务方式1
import (
	"demo/server/controllers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http start----1")
	http.HandleFunc("/user/add", controllers.User{}.Add)
	http.HandleFunc("/user/del", controllers.User{}.Del)
	http.HandleFunc("/user/upd", controllers.User{}.Upd)
	http.HandleFunc("/user/list", controllers.User{}.List)
	http.HandleFunc("/user/getById", controllers.User{}.GetById)
	if err := http.ListenAndServe("localhost:8088", nil); err != nil {
		fmt.Println("启动服务失败", err)
		return
	}
}

// //开启http服务方式2
// import (
// 	"fmt"
// 	"net/http"
// )

// type httpServer struct {
// }

// // 实现ServeHTTP接口
// func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(r.URL.Path))
// }

// func main() {
// 	var server httpServer
// 	http.Handle("/", server)
// 	fmt.Println("http start----2")
// 	if err := http.ListenAndServe("localhost:8088", nil); err != nil {
// 		fmt.Println("启动服务失败", err)
// 		return
// 	}
// }
