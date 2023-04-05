package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func fm(m map[string]string) {
	m["name"] = "李四"
	fmt.Println(m)
}

func fc(c chan int) {
	b := <-c
	println(b)
}

func main() {
	// 1.值类型 int string float bool struct 数组
	// 1.1 int
	// var i int = 10
	// var i_cp = i
	// i_cp = 20
	// fmt.Println(i, i_cp)
	// 1.2 struct
	// var s = Person{
	// 	Name: "张三",
	// 	Age:  20,
	// }
	// var s_cp = s
	// s.Name = "李四"
	// fmt.Println(s, s_cp)
	// 1.3 数组
	// var a = [...]int{1, 2, 3}
	// var a_cp = a
	// a[0] = 10
	// fmt.Println(a, a_cp)
	var a = [3]int{1, 2, 3}
	b := a[:]
	b = append(b, 4)
	fmt.Println(b)
	var a_cp = a
	a[0] = 10
	fmt.Println(a, a_cp)
	// 2.引用类型 map 切片 channel
	// 2.1 map
	// var m = map[string]string{
	// 	"name": "张三",
	// 	"age":  "20",
	// }
	// fm(m)
	// fmt.Println(m)
	// 2.2 切片
	// var s = []string{"php", "golang"}
	// var s_cp = s
	// s_cp[0] = "py"
	// fmt.Println(s, s_cp)
	// 2.3 channel
	// var c = make(chan int, 2)
	// c <- 11
	// c <- 22
	// fc(c)
	// println(<-c)
}
