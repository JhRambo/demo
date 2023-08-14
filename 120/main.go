package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 指定要删除的文件夹路径和目标后缀名
	folderPath := `D:\code\Starverse\com.ghs.utils\proto`
	targetExt := ".gw.go"

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 判断是否是文件，并且后缀名是指定的目标后缀
		if !info.IsDir() && strings.HasSuffix(path, targetExt) {
			fmt.Printf("Deleting file: %s\n", path)
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Error deleting file: %s\n", err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %s\n", err)
	}
}
