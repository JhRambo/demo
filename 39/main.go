package main

import "fmt"

// 别名与自定义类型的区别
type myint int    //自定义类型
type myint2 = int //别名

func main() {

	var a myint = 10
	var b int = 10
	var c myint2 = 10

	fmt.Println(a + myint(b)) //自定义类型需要强制转换
	fmt.Println(int(a) + b)   //自定义类型需要强制转换

	if a == myint(b) { //强制转换
		fmt.Println("hello")
	}

	fmt.Println(c + b) //别名不用强制转换
	if c == b {        //别名不用强制转换
		fmt.Println("hello")
	}
}
