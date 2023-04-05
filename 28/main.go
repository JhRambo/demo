package main

import "fmt"

// 类型转换
func main() {
	var a int
	b := 123.33
	a = int(b)
	fmt.Println(a)	//123
}
