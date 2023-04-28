package main

import "fmt"

func main() {
	a := 1
	// a++ //正确写法1
	a += 1 //正确写法2
	// ++a // 错误写法	golang没有这种写法
	fmt.Println(a)
}
