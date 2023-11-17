package main

import "fmt"

func main() {
	// fmt.Println(f1())
	fmt.Println(f2())
}

// 发生错误了
// 1
func f1() int {
	var err error = fmt.Errorf("eeeeeeeeeeee")
	// 这里注册defer可被执行到
	defer func() {
		if err != nil {
			fmt.Println("发生错误了")
		}
	}()
	return 1
}

func f2() int {
	var err error = fmt.Errorf("eeeeeeeeeeee")
	if err != nil {
		return 2
	}
	// 这里注册defer执行不到
	defer func() {
		if err != nil {
			fmt.Println("发生错误了")
		}
	}()
	return 0
}
