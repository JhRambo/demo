package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Hobby []string
	Work  [2]string
	card  map[string]string //私有属性
	AA    Addr1             //结构体嵌套 别名
	Addr2                   //结构体嵌套 匿名形式
}

type Addr1 struct {
	Name   string
	City   string
	Detail string
}

type Addr2 struct {
	Name   string
	City   string
	Detail string
}

func main() {
	var p Person
	p.Name = "张三"
	p.Age = 20
	p.Hobby = make([]string, 2)
	p.Hobby[0] = "吃饭"
	p.Work[1] = "写bug"
	p.AA.Name = "京城"
	p.AA.Detail = "1000号胡同"
	p.Addr2.Name = "北京"
	p.card = make(map[string]string)
	p.card["card1"] = "建行"
	p.card["card2"] = "工行"
	fmt.Printf("%#v\n", p)
}
