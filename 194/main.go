package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// path := "a.b[1].c"
	path := "b[1]"
	fmt.Println(ParsePath1(path))
	fmt.Println(ParsePath2(path))
	fmt.Println(ParsePath3(path))

	paths := []string{"a", "b[0]", "c", "d"}
	fmt.Println(strings.Join(paths[1+1:], "."))
}

// 解析路径，获取各级字段名
func ParsePath1(path string) []string {
	return strings.Split(path, ".")
}

// 解析path，获取各级字段名
func ParsePath2(path string) []string {
	// 假设path的格式为：nodeList[0].fileInfo[1].name
	// 先按点号分割字段名
	fieldNames := strings.Split(path, ".")

	// 逐个处理字段名
	var result []string
	for _, fieldName := range fieldNames {
		// 判断是否包含索引
		if idx := strings.Index(fieldName, "["); idx != -1 {
			fieldName = fieldName[:idx] // 截取方括号之前的部分
		}
		result = append(result, fieldName)
	}

	return result
}

func ParsePath3(path string) []string {
	var result []string

	// 利用正则表达式匹配路径中的各级字段名和数组索引
	re := regexp.MustCompile(`([^\[\].]+)|(\d+)`)
	matches := re.FindAllStringSubmatch(path, -1)

	// 遍历匹配结果，组装各级字段名和数组索引
	for _, match := range matches {
		if match[1] != "" {
			result = append(result, match[1])
		} else {
			result = append(result, match[2])
		}
	}

	return result
}
