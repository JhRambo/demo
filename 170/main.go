package main

import (
	"demo/logs"
	"fmt"
)

func main() {
	A()
}

func A() {
	defer func() {
		if r := recover(); r != nil {
			logs.Error(logs.GetErrorLocation(r))
		}
	}()

	a := 1
	b := 0
	fmt.Println(a / b)
}
