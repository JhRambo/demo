package main

import (
	"fmt"
)

// 指针，分配内存空间
func main() {
	//错误写法
	// var a *int
	// *a = 10
	// fmt.Println(*a)

	// 正确写法1
	// var a = new(int) //指针
	// fmt.Printf("%v---%T---%p---%v\n", a, a, &a, *a)
	// *a = 10
	// fmt.Printf("%v---%T---%p---%v\n", a, a, &a, *a)

	// 正确写法2
	var a int
	b := &a
	fmt.Printf("%v --- %T --- %p\n", a, a, &a)  //0 int a的地址
	fmt.Printf("%v --- %T --- %p\n", b, b, &b)  //a的地址 *int b的地址
	fmt.Printf("%v --- %T --- %p\n", *b, b, &b) //0 *int b的地址
	// *b = 30
	// fmt.Printf("%v --- %T --- %p\n", a, a, &a)	//30 		*int 		a的地址
	// fmt.Printf("%v --- %T --- %p\n", b, b, &b)	//a的地址 	*int 		b的地址
	// fmt.Printf("%v --- %T --- %p\n", *b, b, &b)	//30 		*int 		b的地址
	// a = 40
	// fmt.Printf("%v --- %T --- %p\n", a, a, &a)	//40 		*int 		a的地址
	// fmt.Printf("%v --- %T --- %p\n", b, b, &b)	//a的地址 	*int 		b的地址
	// fmt.Printf("%v --- %T --- %p\n", *b, b, &b) //40 		*int 		b的地址

	// 正确写法3
	// var str = new(string)
	// *str = "你好golang"
	// fmt.Printf("值：%v，地址：%p\n", *str, &str)

	// 正确写法4
	// var a *int //指针类型
	// a = new(int)
	// *a = 1
	// fmt.Printf("值：%v----类型：%T", *a, a)

	//new 与 make 的区别
	// var a = new(string)
	// fmt.Printf("%v--%T\n", a, a) //0xc000042050--*string
	// var b = make([]string, 1)
	// fmt.Printf("%v--%T", b, b) //[]--[]string
}
