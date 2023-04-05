package main

import "fmt"

//类型断言
func main() {
	var a interface{}
	a = 111
	v1, err1 := a.(int)
	v2, err2 := a.(string)
	fmt.Println(v1, err1) //111 true
	fmt.Println(v2, err2) // false
}
