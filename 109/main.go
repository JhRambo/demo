package main

import (
	"fmt"
	"reflect"
)

type CurUser struct {
	Token    string `abc:"Token"`
	Account  string `abc:"Account"`
	Uid      int32  `abc:"UID"`
	Eid      int32  `abc:"EID"`
	Addr     string `abc:"Addr"`
	FilePath string `abc:"FilePath"`
	License  string `abc:"License"`
	State    string `abc:"State"`
}

func main() {
	user := CurUser{
		Token:    "abc123",
		Account:  "john",
		Uid:      12345,
		Eid:      67890,
		Addr:     "123 Main St",
		FilePath: "/path/to/file",
		License:  "1234-5678-9012",
		State:    "active",
	}

	// 获取结构体的类型信息
	t := reflect.TypeOf(user)
	// 如果 t 是指针类型，则获取其元素类型 &CurUser
	// 值类型：CurUser，以下代码也不会报错。如果是引用类型，没有以下代码，则会报错
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 遍历结构体的字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("abc") // 获取标签值

		fmt.Printf("Field: %s, Tag: %s\n", field.Name, tag)
	}
}
