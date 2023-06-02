package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// 过滤文件的注释代码
func FilterFile(sourcePath string) string {
	// 目标文件的路径
	targetPath := "D:/code/demo/gin/doc/p.proto"

	// 确保目标目录存在
	dir := filepath.Dir(targetPath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 读取原文件内容
	originalData, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		panic(err)
	}

	// 正则表达式
	multiLineCommentRegex := regexp.MustCompile(`(?s)/\*.*?\*/`)
	singleCommentRegex := regexp.MustCompile(`//.*`)

	// 过滤注释
	filteredData := multiLineCommentRegex.ReplaceAllString(string(originalData), "")
	filteredData = singleCommentRegex.ReplaceAllString(filteredData, "")

	// 写入目标文件
	err = ioutil.WriteFile(targetPath, []byte(filteredData), os.ModePerm)
	if err != nil {
		panic(err)
	}

	return targetPath
}
