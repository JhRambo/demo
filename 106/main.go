package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	// 创建一个示例的 map
	myMap := map[string]int{
		"one":   1,
		"three": 3,
		"awo":   2,
	}

	// 提取 map 的键到一个切片中
	keys := make([]string, 0, len(myMap))
	for key := range myMap {
		log.Println(key)
		keys = append(keys, key)
	}

	// 对键进行排序
	sort.Strings(keys)

	// 按排序后的键的顺序输出键值对
	for _, key := range keys {
		fmt.Println(key, myMap[key])
	}
}
