package main

import (
	"fmt"
)

// 指针 引用传值
func swap1(x, y *int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
}

// 值传值
func swap2(x, y int) (int, int) {
	var tmp int
	tmp = x
	x = y
	y = tmp
	return x, y
}

func main() {
	var x = 100
	var y = 200
	// 1.引用传值
	fmt.Printf("交换前x的值是：%v\n", x)
	fmt.Printf("交换前y的值是：%v\n", y)
	swap1(&x, &y)
	fmt.Printf("引用传值，交换后x的值是：%d\n", x)
	fmt.Printf("引用传值，交换后y的值是：%d\n", y)
	// 2.值传值
	xx, yy := swap2(x, y)
	fmt.Printf("值传值，交换后x的值是：%d\n", xx)
	fmt.Printf("值传值，交换后y的值是：%d\n", yy)
}
