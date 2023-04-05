package main

import (
	"fmt"
)

// 管道 引用数据类型
func main() {
	var a = make(chan int, 2) //必须先分配内存空间
	fmt.Println(a)            //0xc00001e180
	a <- 1
	b := <-a
	fmt.Println(a)                         //0xc00001e180
	fmt.Println(b)                         //1
	fmt.Printf("%v---%T---%p\n", a, a, &a) //0xc00001e180---chan int---0xc00000a028
	c := a
	fmt.Printf("%v---%T---%p\n", c, c, &c) //0xc00001e180---chan int---0xc00000a038
}
