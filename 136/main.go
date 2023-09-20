package main

import (
	"fmt"
	"sort"
)

// // 判断值是否在切片中的方案1
// func IsStringInSlice(str string, slice []string) bool {
// 	for _, v := range slice {
// 		if v == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// 判断值是否在切片中的方案2
func IsStringInSlice(str string, slice []string) bool {
	sort.Strings(slice)
	index := sort.SearchStrings(slice, str)
	return index < len(slice) && slice[index] == str
}

func main() {
	str := "apple"
	slice := []string{"banana", "orange", "apple"}

	if IsStringInSlice(str, slice) {
		fmt.Println("目标字符串存在于字符串切片中")
	} else {
		fmt.Println("目标字符串不存在于字符串切片中")
	}
}
