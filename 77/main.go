package main

import (
	"fmt"
)

var g int

func main() {
	var a = 1
	var b = 2
	g = a + b
	fmt.Println("main g==========================", g)
	f1()
	//保持不退出
	for {

	}
}

func f1() {
	fmt.Println("f1 g=============================", g)
}
