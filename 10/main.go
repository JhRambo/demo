package main

import (
	"fmt"
)

// 引用数据类型赋值前必须先分配内存空间
func main() {
	// 错误写法
	// var user map[string]string
	// user["name"] = "zhangsan"
	// fmt.Println(user)

	//正确写法1
	// var user = map[string]string{
	// 	"name": "zhangsan",
	// }
	// fmt.Println(user)

	// 正确写法2
	var user = make(map[string]string)
	user["name"] = "zhangsan"
	fmt.Println(user)
}
