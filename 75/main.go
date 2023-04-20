package main

import "log"

func main() {
	var m map[int]int     //因为未初始化的map值为nil，
	log.Println(m == nil) //输出true
	log.Println(m)        //map[]
}
