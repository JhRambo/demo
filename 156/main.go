package main

import (
	"fmt"
	"reflect"
)

type Table struct {
	Fields TableInfo `json:"fields"`
}

type TableInfo struct {
	// 定义结构体字段
}

func main() {
	t := Table{}
	tv := reflect.ValueOf(t)
	fieldName := reflect.TypeOf(t).Field(0).Name // 获取字段名
	fieldValue := tv.Field(0).Interface()        // 获取字段值
	fmt.Println(fieldName)                       // 输出字段名
	fmt.Println(fieldValue)                      // 输出字段值
}
