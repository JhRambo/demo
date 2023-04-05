package main

import "fmt"

// 函数内append之后，切片指向的是同一个地址
func main() {
	var i = make([]int, 2, 8)
	i[0] = 1
	fmt.Printf("容量：%v，值：%v，地址：%p\n", cap(i), i, &i) //容量：8，值：[1 0]，地址：0xc000008078
	i[1] = 2
	i = append(i, 3)
	fmt.Printf("容量：%v，值：%v，地址：%p\n", cap(i), i, &i) //容量：8，值：[1 2 3]，地址：0xc000008078
	f1 := f1(i)
	fmt.Printf("容量：%v，值：%v，地址：%p\n", cap(f1), f1, &f1) //容量：8，值：[1 2 3 999]，地址：0xc0000080c0
}

//函数外append之后，切片指向的不是同一个地址
func f1(i []int) []int {
	i = append(i, 999)
	return i
	// fmt.Printf("容量：%v，值：%v，地址：%p\n", cap(i), i, &i) //容量：8，值：[1 2 3 999]，地址：0xc0000080c0
}
