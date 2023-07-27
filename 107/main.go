package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// p := Person{Name: "Alice", Age: 25}
	p := Person{}

	// 使用反射获取结构体字段名称
	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name)
	}
	// fmt.Println(p.Name)
	// fmt.Println(p.Age)
}
