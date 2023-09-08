package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// 创建文件
func CreateFile(filePath string, content string) {
	// 获取文件所在目录
	dirPath := filepath.Dir(filePath)
	CreateDir(dirPath)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create new file
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// Replace contents of existing file
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
