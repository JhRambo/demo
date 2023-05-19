package main

import "fmt"

// 可变参数，接口类型的切片
func fn1(x ...interface{}) {
	fmt.Printf("%v----%T", x, x) //[1 2 3 4]----[]interface {}
}

func main() {
	fn1(1, 2, 3, 4)
}
