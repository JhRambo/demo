package main

import "fmt"

// 闭包
// 该函数返回一个函数
func getSequence() func() int {
	i := 0 //1.局部变量 2.常驻内存
	return func() int {
		i += 1   //更改了i的值
		return i //返回1，2，3
		// return i + 1 //没有更改i的值，返回全部返回1
	}
}

func main() {
	/* nextNumber 为一个函数 */
	nextNumber := getSequence() //需要赋值个一个变量，才能沿用上个函数返回的值继续操作
	fmt.Println(nextNumber())   //1
	fmt.Println(nextNumber())   //2
	fmt.Println(nextNumber())   //3

	nextNumber1 := getSequence() //不同的变量，从0开始
	fmt.Println(nextNumber1())   //1
	fmt.Println(nextNumber1())   //2
	fmt.Println(nextNumber1())   //3

	// 没有赋值给变量
	fmt.Println(getSequence()()) //1
	fmt.Println(getSequence()()) //1
	fmt.Println(getSequence()()) //1
}
