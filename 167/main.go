package main

import (
	"encoding/json"
	"fmt"
)

// 示例结构体
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  25,
	}
	// 对象转字符串方式1
	str := fmt.Sprintf("%v", person)
	fmt.Printf("%s-%T\n", str, str)

	// 对象转字符串方式2
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	jsonString := string(jsonData)
	fmt.Printf("%s-%T\n", jsonString, jsonString)
}
