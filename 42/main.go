package main

import (
	"fmt"
)

// 结构体
type Person struct {
	Name string
}

func main() {
	//方式1
	var p1 Person
	p1.Name = "老六"
	fmt.Printf("值：%v 类型：%T\n", p1, p1) //值：{老六} 类型：main.Person

	//方式2
	p2 := struct {
		Name string
		Age  int
	}{"老六", 10}
	fmt.Printf("值：%v 类型：%T\n", p2, p2) //值：{老六 10} 类型：struct { Name string; Age int }

	//方式3
	p3 := struct{}{}
	fmt.Printf("值：%v 类型：%T\n", p3, p3) //值：{} 类型：struct {}

	//方式4
	p4 := make(map[string]struct{})
	p4["name"] = struct{}{}
	p4["age"] = struct{}{}
	_, ok := p4["name"]
	fmt.Println(ok) //true
	_, ok = p4["sex"]
	fmt.Println(ok)        //false
	fmt.Printf("%v\n", p4) //map[age:{} name:{}]
}
