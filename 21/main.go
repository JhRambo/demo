package main

import "fmt"

// 接口
type Usb interface {
	start()
	end()
	Writer //接口嵌套
}

type Writer interface {
	php()
	golang()
}

//空结构体
type Computer struct {
}

type Phone struct {
	Name string
}

func (c Computer) work(usb Usb, name string) {
	fmt.Println(name + "隐性的实现了Usb接口")
}

func (p Phone) start() {
	fmt.Println(p.Name + "开机")
}

func (p Phone) end() {
	fmt.Println(p.Name + "关机")
}

//指针类型，实例化phone结构体必须用地址符&
func (p *Phone) php() {
	fmt.Println(p.Name + "php")
}

func (p Phone) golang() {
	fmt.Println(p.Name + "golang")
}

func main() {
	//1、显性实现接口
	var usb Usb //定义Usb接口类型
	//实例化Phone结构体
	phone := &Phone{
		Name: "华为",
	}
	usb = phone  //Phone结构体实现Usb接口
	usb.start()  //华为开机
	usb.end()    //华为关机
	usb.php()    //华为php
	usb.golang() //华为golang

	//2、非显性的实现接口
	var computer Computer            //定义Computer结构体类型
	computer.work(phone, phone.Name) //因为这里的Phone结构体已经实现了User接口的所有方法，所以可以用这种非显性的实现接口
}
