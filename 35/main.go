package main

//反射reflect.TypeOf
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

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("类型：%v 类型名称：%v 类型种类：%v\n", v, v.Name(), v.Kind())
}

// 类型：int 类型名称：int 类型种类：int
// 类型：float64 类型名称：float64 类型种类：float64
// 类型：bool 类型名称：bool 类型种类：bool
// 类型：string 类型名称：string 类型种类：string
// 类型：main.myInt 类型名称：myInt 类型种类：int
// 类型：string 类型名称：string 类型种类：string
// 类型：main.Student 类型名称：Student 类型种类：struct
// 类型：*int 类型名称： 类型种类：ptr
// 类型：[3]int 类型名称： 类型种类：array
// 类型：[]int 类型名称： 类型种类：slice
func main() {
	a := 10
	b := 12.3
	c := true
	d := "你好golang"
	var e myInt = 20
	var f myType = "hello golang"
	g := Student{
		Name: "张三",
		Age:  20,
	}
	h := 30
	i := [3]int{11, 22, 33}   //数组
	j := []int{111, 222, 333} //切片
	reflectType(a)
	reflectType(b)
	reflectType(c)
	reflectType(d)
	reflectType(e)
	reflectType(f)
	reflectType(g)
	reflectType(&h)
	reflectType(i)
	reflectType(j)
}
