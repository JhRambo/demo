package main

import "fmt"

func main() {
	fmt.Println("A")
	goto B           //执行节点B的所有内容，直到遇到新的节点
	fmt.Println("C") //直接跳过不会被执行
B:
	fmt.Println("B")
	goto C
	fmt.Println("X")
C:
}
