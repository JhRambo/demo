package main

import (
	"fmt"
)

func Add(x, y int) int {
	sum := x + y
	return sum
}

// 函数
func main() {
	res := Add(1, 2)
	fmt.Println(res)
}
