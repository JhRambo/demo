package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	// 注册信号处理程序，以便捕获控制台中断信号
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	shutdownChannel := make(chan struct{})

	err := make(chan error)

	go func() {
		select {
		case sig := <-signalChannel:
			fmt.Println("接收到信号:", sig)
			close(shutdownChannel) // 关闭shutdownChannel
		case <-shutdownChannel:
			fmt.Println("接收到关闭信号")
		}
		fmt.Println(">>>接收到中断信号>>>>>>>>>>>>>>>>>>>>")
		// 推送错误信息至通道
		err <- fmt.Errorf("服务器关闭")
	}()

	// 监测任务进程是否已关闭的goroutine
	go func() {
		for {
			// 检查shutdownChannel是否关闭
			if _, ok := <-shutdownChannel; !ok {
				log.Println("任务进程已关闭")
				// 在这里执行监测到任务进程关闭后的逻辑
				break
			}
			time.Sleep(time.Second) // 每秒钟检查一次
		}
	}()

	// 主 goroutine 中从通道中接收错误信息
	receivedErr := <-err
	if receivedErr != nil {
		fmt.Println("关闭信息：", receivedErr)
	} else {
		fmt.Println("没有接收到关闭信号")
	}

	CreateFile("../doc/error.log", receivedErr.Error())
}

// 创建文件
func CreateFile(filePath string, content string) {
	// 获取文件所在目录
	dirPath := filepath.Dir(filePath)
	// 创建目录，如果存在则不创建
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatalln(err)
		return
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create new file
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			log.Fatalln(err)
			return
		}
	} else {
		// Replace contents of existing file
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalln(err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
