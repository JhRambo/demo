package main

import "fmt"

// 接口类型的默认值及比较
type A interface{}
type B interface{}

func main() {
	var a1 A
	var b1 B
	fmt.Println(a1 == b1) //true

	var a2 A = 10
	var b2 B = 10
	fmt.Println(a2 == b2) //true

	var a7 A = "aaa"
	var b7 B = "aaa"
	fmt.Println(a7 == b7) //true

	var a3 A
	fmt.Println(a3 == nil) //默认值是nil true

	var a4 A
	fmt.Println(a4 == 0) //false

	var a5 A
	fmt.Println(a5 == "") //false

	var a6 A
	fmt.Println(a6 == false) //false
}
