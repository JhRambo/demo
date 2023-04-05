package main

import "fmt"

// 多个 defer 调用顺序是 LIFO（后入先出），defer后的操作可以理解为压入栈中
// defer，return，return value（函数返回值） 执行顺序：首先return，其次return value，最后defer。
// defer可以修改函数最终返回值，修改时机：有名返回值或者函数返回指针

//1.有名返回值
func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) //defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) //defer1: 1
	}()
	return i //或者直接写成return
}

//2.函数返回指针
func c() *int {
	var i int //初始值：0
	defer func() {
		i++
		fmt.Println("defer2:", i) //defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) //defer1: 1
	}()
	return &i
}

func main() {
	fmt.Println("return:", b())  //return: 2
	fmt.Println("return:", *c()) //return: 2
}
