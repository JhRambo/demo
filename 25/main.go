package main

import "fmt"

//递归函数 阶乘
func fn1(x int) int {
	if x == 1 {
		return 1
	} else {
		// x=5：5*fn1(4)
		// x=4：fn1(4) = 4*fn1(3)
		// x=3: fn1(3) = 3*fn1(2)
		// x=2: fn1(2) = 2*fn1(1)
		// x=1: fn1(1) = 1
		return x * fn1(x-1)
	}
}

func main() {
	res := fn1(5)
	fmt.Println(res)
}
