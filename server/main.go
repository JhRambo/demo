package main

//开启http服务方式1
import (
	"demo/server/controllers"
	"fmt"
	"log"
	"net/http"
)

const port = 8088

func main() {
	c := &controllers.User{}
	http.HandleFunc("/user/add", c.Add)
	http.HandleFunc("/user/del", c.Del)
	http.HandleFunc("/user/upd", c.Upd)
	http.HandleFunc("/user/list", c.List)
	http.HandleFunc("/user/getById", c.GetById)
	log.Println(fmt.Sprintf("server启动中，监听%d端口...", port))
	if err := http.ListenAndServe(":"+fmt.Sprintf("%d", port), nil); err != nil {
		log.Println("启动服务失败", err)
		return
	}
}
