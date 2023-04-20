package main

import "fmt"

//接口
type Animaler interface {
	GetName() string
}

//结构体
type Dog struct {
	Name string
}

func (d *Dog) GetName() string {
	return d.Name
}

func main() {
	var a Animaler //定义接口类型的变量
	//实例化结构体
	d := &Dog{
		Name: "大狗",
	}
	a = d //结构体实现Animaler接口
	fmt.Println(a.GetName())
}
