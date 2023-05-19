package main

import (
	"encoding/json"
	"log"
)

const CTX = "ctx"

var MapNode = map[string]string{}

// 自定义结构体内嵌结构体
type NodeRoot struct {
	N1 Node1
	N2 Node2
}

type Node1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Node2 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class string `json:"class"`
}

// 获取node
func GetNodes() *NodeRoot {
	return &NodeRoot{}
}

// 设置node key value的值
func MapNodeData(x *NodeRoot) map[string]string {
	u, _ := json.Marshal(x)
	MapNode[CTX] = string(u)
	return MapNode
}

// 设置node
func setNodes() {
	nodes := GetNodes()
	nodes.N1.Name = "老六"
	nodes.N1.Age = 20
	nodes.N2.Name = "小张"
	nodes.N2.Age = 30
	nodes.N2.Class = "3年2班"
	MapNodeData(nodes)
}

func main() {
	setNodes()
	nodes := GetNodes() //返回空结构体
	strinfo := MapNode[CTX]
	json.Unmarshal([]byte(strinfo), nodes)
	log.Printf("%#v\n", nodes)
}
