package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "how do you do  "
	var arr = strings.Split(str, " ")     //字符串转切片
	fmt.Printf("值：%#v，类型：%T\n", arr, arr) //值：[]string{"how", "do", "you", "do", "", ""}，类型：[]string

	var sss = make(map[string]int)
	for _, v := range arr {
		sss[v]++ //统计切片中每个元素的个数
	}
	fmt.Println(sss)
}
