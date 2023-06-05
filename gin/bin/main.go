package main

import (
	"demo/gin/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

/*
	监听文件内容变化

并自动生成gin网关代码
*/
func main() {
	watchDir := "D:/code/demo/gin/utils/proto" //proto生成的.go文件所在的目录
	readDir := "D:/code/demo/gin/proto"        //proto文件所在的目录

	// 初始化proto文件自动生成代码
	utils.InitProto(readDir)

	// 定义需要热更新的函数
	restartFunc := func() {
		// 自动生成网关所需的代码
		utils.CreateCode()
	}

	restartFunc()

	// 监听指定文件夹下的所有文件的变化
	if err := watchFiles(watchDir, restartFunc); err != nil {
		fmt.Println("Error watching files:", err)
		return
	}
}

// 监听proto生成的.go文件所在的目录
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
