package main

//开启http服务方式1
import (
	"demo/server/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/add", controllers.User{}.Add)
	http.HandleFunc("/user/del", controllers.User{}.Del)
	http.HandleFunc("/user/upd", controllers.User{}.Upd)
	http.HandleFunc("/user/list", controllers.User{}.List)
	http.HandleFunc("/user/getById", controllers.User{}.GetById)
	if err := http.ListenAndServe("localhost:8088", nil); err != nil {
		log.Println("启动服务失败", err)
		return
	}
}
