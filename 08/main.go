package main

import (
	"fmt"
)

// 指针
func main() {
	var a int                                   //默认0
	b := &a                                     //b的类型是指针类型
	fmt.Printf("%v --- %T --- %p\n", a, a, &a)  //0，int，a的地址
	fmt.Printf("%v --- %T --- %p\n", b, b, &b)  //a的地址，*int，b的地址
	fmt.Printf("%v --- %T --- %p\n", *b, b, &b) //0,*int,b的地址
	fmt.Println("------------------------")
	*b = 30
	fmt.Printf("%v --- %T --- %p\n", a, a, &a)  //30，int，a的地址
	fmt.Printf("%v --- %T --- %p\n", *b, b, &b) //30，*int，b的地址
	fmt.Println("------------------------")
	a = 40
	fmt.Printf("%v --- %T --- %p\n", a, a, &a)  //40，int，a的地址
	fmt.Printf("%v --- %T --- %p\n", *b, b, &b) //40，*int，b的地址
}
