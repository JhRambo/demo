package main

import (
	"fmt"
)

func main() {
	//切片是引用数据类型，需要用append扩容
	var arr1 = []int{1, 2, 3}
	arr1 = append(arr1, 5)
	fmt.Println(arr1)                     //[1 2 3 5]
	arr2 := append(arr1[:1], arr1[2:]...) //左闭右开	1,3,5
	fmt.Println(arr2)                     //[1 3 5]

	//或者用make先分配内存空间
	arr3 := make([]int, 4)
	arr3[2] = 1
	arr3 = append(arr3, 4)
	fmt.Println(arr3) //[0 0 1 0 4]

	arr4 := []interface{}{1, 3, 4, 5, "zs"}
	fmt.Println(arr4) //[0 0 1 0 4]
}
