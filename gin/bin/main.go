package main

import (
	"demo/gin/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

/*
	监控文件内容变化

并自动生成代码
*/
func main() {
	watchDir := "D:/code/demo/gin/utils/proto" //proto生成的.go文件所在的目录
	readDir := "D:/code/demo/gin/proto"        //proto文件所在的目录

	// 定义需要热更新的函数
	restartFunc := func() {
		if err := generateCode(readDir); err != nil {
			fmt.Println("Error generating code:", err)
			return
		}
		fmt.Println("Code successfully generated...")
	}

	if err := generateCode(readDir); err != nil {
		fmt.Println("Error generating code:", err)
		return
	}

	// 监听指定文件夹下的所有文件的变化
	if err := watchFiles(watchDir, restartFunc); err != nil {
		fmt.Println("Error watching files:", err)
		return
	}
}

// 遍历proto文件所在的目录
func generateCode(readDir string) error {
	// 读取文件夹下所有文件
	files, err := ioutil.ReadDir(readDir)
	if err != nil {
		return err
	}

	// 遍历文件，生成代码
	for _, file := range files {
		if !file.IsDir() {
			fmt.Printf("Generating code from %s...\n", file.Name())
			// 根据文件生成所需代码
			utils.CreateCode()
		}
	}

	return nil
}

// 监控proto生成的.go文件所在的目录
func watchFiles(watchDir string, restartFunc func()) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
					restartFunc()
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	// 监听指定文件夹下的所有文件
	err = filepath.Walk(watchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if err := watcher.Add(path); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	<-done
	return nil
}
