package main

import "fmt"

func fn1() (a int) { //有名返回值，defer可以修改return的返回值
	defer func() {
		a++
	}()
	return a //输出：1
}

func fn2() int { //无名返回值，defer不可以修改return的返回值
	var a = 0
	defer func() {
		a++
	}()
	return a //输出：0
}

func fn3(x, y int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err：", err)
			// panic("抛出一个异常")
		}
	}()
	return x / y
}

// 开始
// 结束
// 1
// 0
// err： runtime error: integer divide by zero
// 0
// 555
// 444
// 333
// 222
// 111
func main() {
	fmt.Println("开始")
	// defer 延迟执行，逆序执行	先后输出 555 444 333 222 111
	defer fmt.Println("111")
	defer fmt.Println("222")
	defer fmt.Println("333")
	defer func() {
		defer fmt.Println("444")
		defer fmt.Println("555")
	}()
	fmt.Println("结束")
	fmt.Println(fn1())      //1
	fmt.Println(fn2())      //0
	fmt.Println(fn3(10, 0)) //err： runtime error: integer divide by zero
}
