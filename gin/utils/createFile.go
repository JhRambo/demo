package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// 创建文件
func CreateFile(filePath string, content string) {
	// 获取文件所在目录
	dirPath := filepath.Dir(filePath)
	// 创建目录，如果存在则不创建
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create new file
		err := ioutil.WriteFile(filePath, []byte(content), 0644)
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

	// Format the file
	cmd := exec.Command("go", "fmt", filePath)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
