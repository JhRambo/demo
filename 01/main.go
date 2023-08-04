package main

import (
	"fmt"
	"sort"
)

// map
func main() {
	// 1.创建map类型第一种方式：make函数创建map类型
	// var user = make(map[string]string)
	// user["name"] = "张三"
	// user["age"] = "20"
	// user["sex"] = "男"

	// 2.创建map类型第二种方式：
	// var user = map[string]string{
	// 	"name": "张三",
	// 	"age": "20",
	// 	"sex": "男",
	// }

	// 3.创建map类型第三种方式：
	user := map[string]string{
		"name": "张三",
		"age":  "20",
		"sex":  "男",
	}
	fmt.Printf("%v--%p\n", user, user) //map[age:20 name:张三 sex:男]--0xc0000221b0

	//1.无序输出
	// for k, v := range user {
	// 	// age 20
	// 	// sex 男
	// 	// name 张三
	// 	fmt.Println(k, v)
	// }
	//2.有序输出
	var keys []string
	for k, _ := range user {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)

	// v, ok := user["age"]
	// fmt.Println(v, ok) //20 true

	// delete(user, "age") //删除属性
	// fmt.Println(user)   //map[name:张三 sex:男]
}
