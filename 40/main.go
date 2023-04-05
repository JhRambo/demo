package main

import "fmt"

type calc func(int, int) int    //自定义函数类型
type calc2 = func(int, int) int //自定义函数别名

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	var c calc
	c = add
	c = sub
	fmt.Printf("类型：%T，值：%d\n", c, c(10, 20)) //类型：main.calc，值：-10

	var c2 calc2
	c2 = add
	c2 = sub
	fmt.Printf("类型：%T，值：%d\n", c2, c2(10, 20)) //类型：func(int, int) int，值：-10

	d := add
	d = sub
	fmt.Printf("类型：%T，值：%d\n", d, d(10, 20)) //类型：func(int, int) int，值：-10
}
