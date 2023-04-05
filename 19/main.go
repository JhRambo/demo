package main

import (
	"fmt"
	"math"
)

// 匿名函数
/* 声明函数变量 */
var getSquareRoot = func(x float64) float64 {
	return math.Sqrt(x)
}

// 函数作为另外一个函数的实参
func main() {
	/* 使用函数 函数做实参*/
	fmt.Println(getSquareRoot(9))
}
