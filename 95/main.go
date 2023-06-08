package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	// 创建结构体实例
	p := Person{"Bob", 30}

	// // 获取结构体的反射类型和值
	// pType := reflect.TypeOf(p)
	// pValue := reflect.ValueOf(p)

	// 获取结构体的反射类型和值
	pType := reflect.TypeOf(p)
	pValue := reflect.ValueOf(&p).Elem()

	// // 修改结构体字段的值
	// for i := 0; i < pType.NumField(); i++ {
	// 	// 获取结构体字段的名称和值
	// 	fieldName := pType.Field(i).Name
	// 	fieldValue := pValue.Field(i)
	// 	// 如果是Name字段，则设置新的值
	// 	if fieldName == "Name" {
	// 		fieldValue.SetString("Alice")
	// 	}
	// }

	// 获取结构体的字段列表
	for i := 0; i < pType.NumField(); i++ {
		// 获取结构体字段的名称和值
		fieldName := pType.Field(i).Name
		fieldValue := pValue.Field(i)
		// 打印结果
		fmt.Printf("%v: %v\n", fieldName, fieldValue.Interface())
	}
}
