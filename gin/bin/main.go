package main

import (
	"demo/gin/utils"
	"log"
)

/*
	实现

自动注册router路由
自动注册grpc客户端
*/
func main() {
	utils.InitRouters()
	utils.InitHandlers()
	log.Println("File created and updated successfully.")
}
