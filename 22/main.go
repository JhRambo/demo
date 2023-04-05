package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Hobby []string  //切片
	Work  [2]string //数组
	Addr  interface{}
	card  map[string]string //私有属性
}

func main() {
	p := Person{
		Name: "张三",
		Age:  20,
	}
	p.Hobby = make([]string, 2) //引用数据类型必须先分配内存空间
	p.Hobby[0] = "吃饭"
	p.Work[1] = "写bug"
	p.Addr = "北京"
	p.card = make(map[string]string) //引用数据类型必须先分配内存空间
	p.card["card1"] = "建行"
	p.card["card2"] = "工行"
	fmt.Printf("%#v", p)
}
