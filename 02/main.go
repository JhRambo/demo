package main

import "fmt"

// 切片 map
func main() {
	// 1.切片：
	var s1 = []string{"zhangsan", "lishi"}
	fmt.Printf("%#v---%T\n", s1, s1) //[]string{"zhangsan", "lishi"}---[]string

	//2.切片索引只能是int型：
	s := []string{
		1: "aaaa",
		2: "bbbb",
	}
	fmt.Println(s) //[ aaaa bbbb]

	// 3.定义map类型的切片
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

	// 4.定义切片[]类型的map：
	var ms1 = make(map[string][]string)
	ms1["hobby"] = []string{
		"吃饭",
		"睡觉",
		"写bug",
	}
	ms1["work"] = []string{
		"php",
		"golang",
		"vue",
	}
	fmt.Printf("%v---%T\n", ms1, ms1) //map[hobby:[吃饭 睡觉 写bug] work:[php golang vue]]---map[string][]string

	// 5.定义切片[]类型的map：
	var ms2 = make(map[string][]map[string]string)
	uu := []map[string]string{}
	u1 := map[string]string{
		"date": "2023-05-07",
		"time": "07:07:07",
	}
	u2 := map[string]string{
		"date": "2023-05-07",
		"time": "07:07:07",
	}
	uu = append(uu, u1, u2)
	ms2["dev1"] = uu
	ms2["dev2"] = uu
	fmt.Println(ms2) //map[dev1:[map[date:2023-05-07 time:07:07:07] map[date:2023-05-07 time:07:07:07]] dev2:[map[date:2023-05-07 time:07:07:07] map[date:2023-05-07 time:07:07:07]]]
}
