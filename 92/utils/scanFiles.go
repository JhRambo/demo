package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// 扫描.proto文件，并提取文件名
func ScanFiles() []string {
	// 指定目录路径
	dir := "D:/code/demo/gin/proto"
	fileNames := []string{} // 用于存储文件名的切片

	// 遍历目录下的所有文件和目录
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("读取目录失败:", err)
		return fileNames
	}
	// 遍历文件列表，找出指定后缀名为 ".proto" 的文件
	for _, fi := range fileInfos {
		// 判断是否是普通文件，并且后缀名为 ".proto"
		if !fi.IsDir() && filepath.Ext(fi.Name()) == ".proto" {
			// 打印文件名
			fileNames = append(fileNames, strings.Split(fi.Name(), ".")[0])
		}
	}
	return fileNames
}
