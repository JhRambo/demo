package utils

import (
	"log"
)

/*
	实现

自动注册router路由
自动注册grpc客户端
*/
func CreateCode() {
	InitRouters()
	InitHandlers()
	log.Println("File created and updated successfully.")
}
