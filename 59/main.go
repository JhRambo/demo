package main

import "fmt"

//for-range与for的区别：范围遍历(for-range)会在遍历之前，先拷贝一份被遍历的数据，然后遍历拷贝的数据
func main() {
	// 1.for-range
	s := []int{0, 1}
	for _, v := range s {
		s = append(s, v)
	}
	fmt.Printf("s=%v\n", s) //s=[0 1 0 1]

	// 2.for
	// s := []int{0, 1}
	// for i := 0; i < len(s); i++ {
	// 	s = append(s, s[i])
	// }
	// fmt.Printf("s=%v\n", s) //没有输出内容，无限死循环，因为没有拷贝一份数据，是在原有的数据上做追加，所以一直在遍历追加
}
