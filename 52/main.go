package main

import "fmt"

// 切片与nil比较
func main() {
	var slice []int
	fmt.Println(slice == nil) //true
	slice = make([]int, 0)
	fmt.Println(slice == nil) //false
	slice = []int{}
	fmt.Println(slice == nil) //false
}
