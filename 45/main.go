package main

import (
	"fmt"
)

func checkSlice(s int, slice []int) (flag bool) { //bool类型默认false
	for _, v := range slice {
		fmt.Println(v, s)
		if v == s {
			flag = true
			return
		}
	}
	return
}

// 切片去重 没有内置函数吗？
func main() {
	var strSlice = []int{1, 2, 3, 3, 2, 1, 4}
	var strSliceUnique = []int{}
	for i := 0; i < len((strSlice)); i++ {
		if checkSlice(strSlice[i], strSliceUnique) == false {
			strSliceUnique = append(strSliceUnique, strSlice[i])
		}
	}
	fmt.Println(strSliceUnique)
}
