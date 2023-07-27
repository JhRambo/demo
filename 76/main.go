package main

import (
	"log"
	"time"
)

func g1() {
	log.Println("========g1")
}
func g2() {
	log.Println("g2========")
}

func main() {
	go g1()
	go g2()
	log.Println("========================================main")
	time.Sleep(time.Second * 10) //10秒之后退出main函数
}
