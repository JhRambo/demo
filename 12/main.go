package main

import "fmt"

// 指针，分配内存空间
func main() {
	//错误写法
	// var a *int
	// *a = 10
	// // 报错:panic: runtime error: invalid memory address or nil pointer dereference
	// // [signal 0xc0000005 code=0x1 addr=0x0 pc=0x9fdd36]
	// // 原因是 go 初始化指针的时候会为指针 a 的值赋为 nil ，但a的值代表的是 *a的地址， nil的话系统还并没有给 *a分配地址，所以这时给 *a赋值肯定会出错
	// fmt.Println(*a)

	// 正确写法1
	// var a = new(int) //指针
	// fmt.Printf("%v---%T---%p---%v\n", a, a, &a, *a)
	// *a = 10
	// fmt.Printf("%v---%T---%p---%v\n", a, a, &a, *a)

	// 正确写法2
	// var a int
	// b := &a
	// fmt.Printf("%v --- %T --- %p\n", a, a, &a)  //0 int a的地址
	// fmt.Printf("%v --- %T --- %p\n", b, b, &b)  //a的地址 *int b的地址
	// fmt.Printf("%v --- %T --- %p\n", *b, b, &b) //0 *int b的地址
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
	var a = new(string)
	fmt.Printf("%v--%T\n", a, a) //0xc000042050--*string
	var b = make([]string, 0)
	fmt.Printf("%v--%T", b, b) //[]--[]string
}
