package main

import "fmt"

func main() {
	var i interface{} = []interface{}{"apple", "banana", "cherry"}

	// 类型断言和类型转换
	if interfaceSlice, ok := i.([]interface{}); ok {
		strSlice := make([]string, len(interfaceSlice))
		for i, v := range interfaceSlice {
			if str, ok := v.(string); ok {
				// 类型断言，将v转换为string类型
				strSlice[i] = str
			} else {
				// 类型断言失败，v不是string类型
				fmt.Printf("Invalid element at index %d\n", i)
			}
		}
		fmt.Println(strSlice) // Output: [apple banana cherry]
	} else {
		// 类型断言失败，i不是[]interface{}类型
		fmt.Println("Cannot convert to []interface{}")
	}
}
