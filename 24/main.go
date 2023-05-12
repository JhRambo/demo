package main

import (
	"encoding/json"
	"fmt"
)

// 结构体，json序列化，反序列化
type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"` //私有属性，不会转换成json格式
}

type Class struct {
	Title   string    `json:"title"`
	Student []Student `json:"student"` //结构体嵌套
}

func main() {
	var c = &Class{
		Title:   "001班",
		Student: make([]Student, 0),
	}
	for i := 1; i <= 3; i++ {
		s := Student{
			Id:   i,
			Name: fmt.Sprintf("stu_%v", i),
			Sex:  "男",
		}
		c.Student = append(c.Student, s)
	}
	v, _ := json.Marshal(c) //json序列化
	str := string(v)
	fmt.Printf("%#v\n", str)

	var cc = &Class{} //返回空结构体
	fmt.Printf("%#v\n", cc)
	json.Unmarshal([]byte(str), cc) //json反序列化
	fmt.Printf("%#v\n", cc)
}
