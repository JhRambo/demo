package utils

import (
	"fmt"
	"os"
)

func CreateDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create %v directory: %v", dirPath, err)
			return
		}
		fmt.Println("success")
	}
}
