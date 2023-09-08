package main

import (
	"fmt"
	"strings"
)

func main() {
	k := "example:key:value"

	// 用于返回字符串 k 中从右向左搜索第一个出现的冒号字符 : 的索引位置。
	index := strings.LastIndex(k, ":")
	fmt.Println(index) // 输出 11	索引从0开始
}
