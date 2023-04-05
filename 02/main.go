package main

import "fmt"

// 切片
func main() {
	// 1.定义切片类型第一种方式：
	// var user = []string{"zhangsan", "lishi"}
	// fmt.Printf("%#v---%T\n", user, user) //[]string{"zhangsan", "lishi"}---[]string

	// 2.定义切片[]结合map方式1：
	var user = make([]map[string]string, 3)
	if user[0] == nil {
		user[0] = make(map[string]string)
		user[0]["name"] = "张三"
		user[0]["age"] = "20"
		user[0]["sex"] = "男"
	}
	user[2] = make(map[string]string)
	user[2]["name"] = "李四"
	user[2]["age"] = "30"
	user[2]["sex"] = "女"
	fmt.Printf("%v---%T---%v\n", user, user, len(user)) //[map[age:20 name:张三 sex:男] map[] map[age:30 name:李四 sex:女]]---[]map[string]string---3
	// for _, v := range user {
	// 	for kk, vv := range v {
	// 		fmt.Println(kk, vv)
	// 	}
	// }

	// 3.定义切片[]结合map方式2：
	// var user = make(map[string][]string)
	// user["hobby"] = []string{
	// 	"吃饭",
	// 	"睡觉",
	// 	"写bug",
	// }
	// user["work"] = []string{
	// 	"php",
	// 	"golang",
	// 	"vue",
	// }
	// fmt.Printf("%v---%T\n", user, user) //map[hobby:[吃饭 睡觉 写bug] work:[php golang vue]]---map[string][]string
	// for _, v := range user {
	// 	for kk, vv := range v {
	// 		fmt.Println(kk, vv)
	// 	}
	// }
}
