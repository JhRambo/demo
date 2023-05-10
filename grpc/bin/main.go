package main

import (
	"demo/grpc/utils"
	"log"
)

func main() {
	err := utils.CreateTable()
	if err != nil {
		//创建失败可以增加告警通知=========================================== TODO
		log.Fatalln(err)
	}
}
