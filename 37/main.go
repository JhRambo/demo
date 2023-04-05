package main

//反射reflect.Elem().set
import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Elem().Kind() //值类型种类
	if kind == reflect.Int {
		v.Elem().SetInt(20)
	} else if kind == reflect.String {
		v.Elem().SetString("hello world")
	}
}

func main() {
	a := 10
	b := "你好golang"
	reflectValue(&a)
	reflectValue(&b)

	fmt.Println(a)
	fmt.Println(b)
}
