package main

import (
	"fmt"
)

// 切片copy复制
func main() {
	var arr = []int{1, 2, 3}
	var arr2 = make([]int, 3, 3)
	// copy(arr2, arr) //copy函数复制，修改一个切片的值【不会】影响另外一个切片的值
	arr2 = arr	//等号赋值，修改一个切片的值【会】影响另外一个切片的值

	fmt.Printf("arr地址：%p，arr2地址：%p\n", arr, arr2)

	fmt.Println(arr, len(arr))
	fmt.Println(arr2, len(arr2))

	arr2[0] = 4 //修改arr2的值
	fmt.Println(arr, len(arr))
	fmt.Println(arr2, len(arr2))
}
