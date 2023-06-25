package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nodes interface{}
	nodes = "123456"
	nodes = 123456

	v := reflect.ValueOf(nodes)

	fmt.Printf("%T,%s\n", v, v)
	fmt.Println("reflect...")
	fmt.Printf("%T,%s\n", v, v)
}
