package main

//反射reflect.ValueOf
import (
	"fmt"
	"reflect"
)

type myInt int
type myType interface{}
type Student struct {
	Name string
	Age  int
}

// switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	// fmt.Printf("值：%v 值类型种类：%v\n", v, v.Kind())
	kind := v.Kind() //值类型种类
	fmt.Println(kind)
	switch kind {
	case reflect.Int:
		fmt.Printf("int类型原始值：%v，计算后的值：%v\n", v.Int(), v.Int()+99)
	case reflect.Ptr:
		fmt.Println("指针类型")
	default:
		fmt.Println("不参与计算")
	}
}

func main() {
	a := 10
	b := 12.3
	c := true
	d := "你好golang"
	var e myInt = 20
	var f myType = 30
	g := Student{
		Name: "张三",
		Age:  20,
	}
	h := 30
	i := [3]int{11, 22, 33}   //数组
	j := []int{111, 222, 333} //切片
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
	reflectValue(d)
	reflectValue(e)
	reflectValue(f)
	reflectValue(g)
	reflectValue(&h) //指针类型
	reflectValue(i)
	reflectValue(j)
}
