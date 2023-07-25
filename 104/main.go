package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// 存储自定义类型的数据
	type MyStruct struct {
		Field1 string
		Field2 int
	}
	value := MyStruct{
		Field1: "hello",
		Field2: 123,
	}
	m.Store("key3", value)

	// 加载 key3 对应的值
	if val, ok := m.Load("key3"); ok {
		// 修改 Field1 字段的值
		if v, ok := val.(MyStruct); ok {
			v.Field1 = "" // 将 Field1 设置为空字符串
			// 存储修改后的值回 sync.Map
			m.Store("key3", v)
		}
	}

	// 通过 Load 方法获取修改后的值
	if val, ok := m.Load("key3"); ok {
		fmt.Printf("%#v\n", val)
	}
}
