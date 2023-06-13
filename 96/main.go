package main

import (
	"demo/96/utils"
	"fmt"
)

func main() {
	// // 创建
	// utils.Create()

	// //更新字段
	// path := "data.node1.baseInfo.type"
	// val := "模型2222222"
	// eid := int64(1000)
	// action := ""
	// configId := int64(1)
	// spaceId := int64(1)
	// resp := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	// fmt.Printf("%#v", resp)

	// //新增字段
	// path := "data.node1.transformInfo.aaa.y"
	// val := "200.1"
	// eid := int64(1000)
	// action := ""
	// configId := int64(1)
	// spaceId := int64(1)
	// resp := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	// fmt.Printf("%#v", resp)

	//删除字段
	path := "data.node1.transformInfo.aaa"
	val := ""
	eid := int64(1000)
	action := "d"
	configId := int64(1)
	spaceId := int64(1)
	resp := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	fmt.Printf("%#v", resp)
}
