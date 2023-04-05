package main

import "fmt"

func main() {
	var m = make([]map[string]interface{}, 3)
	m[0] = map[string]interface{}{
		"name": "张三",
		"age":  20,
		"sex":  "男",
	}
	m[1] = map[string]interface{}{
		"name": "李四",
		"age":  30,
		"sex":  "男",
	}
	m[2] = map[string]interface{}{
		"name": "王五",
		"age":  40,
		"sex":  "男",
	}
	var slice []string
	for _, v := range m {
		if vv, ok := v["name"].(string); ok {
			slice = append(slice, vv)
		}
	}
	fmt.Println(slice)
}
