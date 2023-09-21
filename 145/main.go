package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.RemoveAll("nonexistent.txt")
	if err != nil {
		fmt.Printf("删除文件出错：%v\n", err)
	} else {
		fmt.Println("文件删除成功！")
	}
}
