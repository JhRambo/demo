package main

import "fmt"

func c() *int {
	var i int //初始值0
	defer func() {
		i++
		fmt.Println("defer2:", i) //defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) //defer1: 1
	}()
	return &i
}

func main() {
	fmt.Println("return:", *(c())) //return: 2
}
