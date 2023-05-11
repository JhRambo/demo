package main

import (
	"encoding/json"
	"fmt"
)

// 结构体，json序列化，反序列化
type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	sex  string `json:"sex"` //私有属性，不可以转换成json格式
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
			sex:  "男",
		}
		c.Student = append(c.Student, s)
	}
	v, _ := json.Marshal(c) //json序列化
	str := string(v)
	fmt.Println(str)

	err := json.Unmarshal([]byte(str), &c) //json反序列化
	if err == nil {
		fmt.Printf("%#v\n", c)
	}
}
