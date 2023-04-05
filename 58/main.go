package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func fn1(a int) {
	println(a) //1
	a = 2
}

func fn2(a *int) {
	*a = 2 //值类型的int，这里必须用*?
}

func fn3(s *Person) {
	s.Name = "老六"        //值类型的结构体，这里不能用*?
	fmt.Printf("%+v", s) //&{Name:老六 Age:33}
}
func main() {
	a := 1
	fn1(a)
	fmt.Println(a) //1
	fn2(&a)
	fmt.Println(a) //2
	p := Person{
		Name: "张三",
		Age:  33,
	}
	fn3(&p)
	fmt.Printf("%+v", p) //{Name:老六 Age:33}
}
