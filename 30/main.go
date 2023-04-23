package main

import "fmt"

type User struct {
	Name string
	Age  int
}

//类型断言
func main() {
	var a interface{}
	a = 123
	v1, err1 := a.(int)
	v2, err2 := a.(string)
	fmt.Println(v1, err1) //123 true
	fmt.Println(v2, err2) // false

	user := User{}
	u := f1(user)
	v3, err3 := u.(User)
	fmt.Println(v3, err3)
}

func f1(zdy interface{}) interface{} {
	return zdy
}
