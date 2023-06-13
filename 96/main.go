package main

import (
	"demo/96/utils"
	"fmt"
)

func main() {
	// // 创建
	// utils.Create()

	// //更新字段1
	// path := "data.node1.baseInfo.type"
	// val := "模型333333"
	// eid := int64(1000)
	// action := ""
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	// if err != nil {
	// 	fmt.Printf("%#v", err)
	//  return
	// }
	// fmt.Printf("%#v", resp)

	//更新字段2
	path := "data.node1.baseInfo.type"
	val := "[type1,type2]"
	eid := int64(1000)
	action := ""
	configId := int64(1)
	spaceId := int64(1)
	resp, err := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	fmt.Printf("%#v", resp)

	// //更新字段3 操作数组
	// path := "data.node2.slices.0.a" //更新指定索引下标的值
	// val := "11"
	// eid := int64(1000)
	// action := ""
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	// if err != nil {
	// 	fmt.Printf("%#v", err)
	//  return
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段4 操作数组
	// path := "data.node2.slices.2.a.1.name" //更新指定索引下标的值
	// val := "张三李四王五老六"
	// eid := int64(1000)
	// action := ""
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	// if err != nil {
	// 	fmt.Printf("%#v", err)
	// 	return
	// }
	// fmt.Printf("%#v", resp)

	// //新增字段
	// path := "data.node1.transformInfo.aaa.y"
	// val := "200.1"
	// eid := int64(1000)
	// action := ""
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	// if err != nil {
	// 	fmt.Printf("%#v", err)
	//  return
	// }
	// fmt.Printf("%#v", resp)

	// //删除字段
	// path := "data.node2.slices[0]"
	// val := ""
	// eid := int64(1000)
	// action := "d"
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.UpdateSpace(path, val, action, configId, spaceId, eid)
	// if err != nil {
	// 	fmt.Printf("%#v", err)
	//  return
	// }
	// fmt.Printf("%#v", resp)
}
