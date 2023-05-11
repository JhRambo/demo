package main

import "fmt"

// 自定义结构体内嵌结构体
type NodeRoot struct {
	Node1
	Node2
}

type Node1 struct {
	Name string
	Age  int
}

type Node2 struct {
	Name  string
	Age   int
	Class string
}

func main() {
	r := GetNode()
	r.Node1.Name = "老六"
	fmt.Println(r.Node1)
}

func GetNode() *NodeRoot {
	resp := &NodeRoot{}
	return resp
}
