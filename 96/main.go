package main

import (
	"demo/96/utils"
	"fmt"
	"log"
)

func main() {
	// // 创建
	// resp, err := utils.Create()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段1 操作值类型或object类型 新增字段 ok
	// path := []string{"basedata.light.x"}
	// data := []interface{}{111.1}
	// eid := int64(1000)
	// action := []string{""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段2 操作值类型或object类型 新增字段 ok
	// path := []string{"basedata.light.y", "basedata.desc"}
	// data := []interface{}{"222.2", "basedata描述内容"}
	// eid := int64(1000)
	// action := []string{"", ""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段3 操作值类型或object类型 修改值内容 ok
	// path := []string{"basedata.light.y", "basedata.desc"}
	// data := []interface{}{"[11,00]", "啦啦啦啦啦啦啦啦啦"}
	// eid := int64(1000)
	// action := []string{"", ""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段4 操作值类型或object类型 删除字段 ok
	// path := []string{"basedata.light.y"}
	// data := []interface{}{""}
	// eid := int64(1000)
	// action := []string{"d"}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段5 操作数组 修改值内容 ok
	// path := []string{"nodeList.1.slices.0.a", "nodeList.0.baseInfo.description"}
	// data := []interface{}{"111111", "baseinfo的描述内容===================="}
	// action := []string{"", ""}
	// eid := int64(1000)
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段6 操作数组 新增字段 ok
	// path := []string{"nodeList.1.slices.0.c", "nodeList.1.slices.1.c"}
	// data := []interface{}{"x", "y"}
	// eid := int64(1000)
	// action := []string{"", ""}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段7 操作数组 删除字段 ok
	// path := []string{"nodeList.1.slices.0.c", "nodeList.1.slices.1.c", "nodeList.1.slices.2.c"}
	// data := []interface{}{"", "", ""}
	// eid := int64(1000)
	// action := []string{"d", "d", "d"}
	// configId := int64(1)
	// spaceId := int64(1)
	// resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	//更新字段8 操作数组或值类型或object类型 不同类型同时更新字段 ok
	path := []string{"nodeList.1.slices.0.a", "basedata.light.y"}
	data := []interface{}{"aaaaaaaaaa", "2222222222"}
	eid := int64(1000)
	action := []string{"", ""}
	configId := int64(1)
	spaceId := int64(1)
	resp, err := utils.Update(path, data, action, configId, spaceId, eid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", resp)
}
