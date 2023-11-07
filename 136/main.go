package main

import (
	"fmt"
	"sort"
)

// // 判断值是否在切片中的方案1  string
// func IsStringInSlice(str string, slice []string) bool {
// 	for _, v := range slice {
// 		if v == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// 判断值是否在切片中的方案2  string
func IsStringInSlice(str string, slice []string) bool {
	sort.Strings(slice)
	index := sort.SearchStrings(slice, str)
	return index < len(slice) && slice[index] == str
}

// 判断值是否在切片中的方案3  int
func IsIntInSlice(i int, slice []int) bool {
	sort.Ints(slice)
	index := sort.SearchInts(slice, i)
	return index < len(slice) && slice[index] == i
}

func main() {
	// str := "apple"
	// slice := []string{"banana", "orange", "apple"}

	// if IsStringInSlice(str, slice) {
	// 	fmt.Println("目标字符串存在于字符串切片中")
	// } else {
	// 	fmt.Println("目标字符串不存在于字符串切片中")
	// }

	str := 1
	slice := []int{1, 2, 3}

	if IsIntInSlice(str, slice) {
		fmt.Println("目标字符串存在于字符串切片中")
	} else {
		fmt.Println("目标字符串不存在于字符串切片中")
	}
}
