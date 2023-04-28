package main

import "fmt"

func main() {
	fmt.Println("A")
	goto B
	fmt.Println("C") //直接跳过不会执行
B:
	fmt.Println("B")
}
