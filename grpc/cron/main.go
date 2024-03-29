package main

import (
	"demo/grpc/utils"
	"log"
	"time"
)

// 定时器任务每月执行一次
func main() {
	ticker := time.NewTicker(time.Second * 10)
	// ticker := time.NewTicker(time.Hour * 24 * 31)
	for range ticker.C {
		table, err := utils.CreateTable()
		if err != nil {
			//创建失败可以增加告警通知或者记录错误日志表===========================================TODO
			log.Println(table, "创建失败：", err)
		} else {
			log.Println(table, "创建成功")
		}
	}
}
