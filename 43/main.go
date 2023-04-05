package main

import "fmt"

// 可变参数，接口类型的切片
func fn1(a ...interface{}) {
	fmt.Printf("%v----%T", a, a) //[1 2 3 4]----[]interface {}
}
func main() {
	fn1(1, 2, 3, 4)
}
