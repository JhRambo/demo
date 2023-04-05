package main

import "fmt"

//定义一个函数返回指针类型
func getzz() *int {
	var cnumber = 12
	c := &cnumber
	fmt.Println(*c) //12
	return c        //cnumber的地址
}

func main() {
	fmt.Println(getzz())
	// fmt.Println(*getzz())
}
