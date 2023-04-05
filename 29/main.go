package main

import "fmt"

//接口
type Animaler interface {
	SetName(string)
	GetName() string
}

//结构体
type Dog struct {
	Name string
}

func (d Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

func main() {
	var a Animaler
	d := Dog{
		Name: "大狗",
	}
	a = d //表示实现Animaler接口
	a.SetName(d.Name)
	fmt.Println(a.GetName())
}
