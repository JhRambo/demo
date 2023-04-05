package main

import (
	"fmt"
)

// 数组
func main() {
	var arr = [2]int{1}
	arr[1] = 11	//修改数组的值
	fmt.Println(arr)
}