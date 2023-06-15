package main

import (
	"demo/96/utils"
	"encoding/json"
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
	// jsons := `[{
	// 	"path": "basedata.light.x",
	// 	"value": "111.1",
	// 	"action": "",
	// 	"nodeId": "light"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段2 操作值类型或object类型 新增字段 ok
	// jsons := `[{
	// 	"path": "basedata.light.y",
	// 	"value": "222.2",
	// 	"action": "",
	// 	"nodeId": "light"
	// }, {
	// 	"path": "basedata.desc",
	// 	"value": "basedata描述内容",
	// 	"action": "",
	// 	"nodeId": "desc"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段3 操作值类型或object类型 新增字段 ok
	// jsons := `[{
	// 	"path": "basedata.light.y",
	// 	"value": "[11,00]",
	// 	"action": "",
	// 	"nodeId": "light"
	// }, {
	// 	"path": "basedata.desc",
	// 	"value": "啦啦啦啦啦啦啦啦啦",
	// 	"action": "",
	// 	"nodeId": "desc"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段4 操作值类型或object类型 删除字段 ok
	// jsons := `[{
	// 	"path": "basedata.light.y",
	// 	"value": "",
	// 	"action": "d",
	// 	"nodeId": "light"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段5 操作数组 修改值内容 ok
	// jsons := `[{
	// 	"path": "nodeList.1.slices.0.a",
	// 	"value": "111111",
	// 	"action": "",
	// 	"nodeId" : "1"
	// },{
	// 	"path": "nodeList.0.baseInfo.description",
	// 	"value": "baseinfo的描述内容====================",
	// 	"action": "",
	// 	"nodeId" : "0"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段6 操作数组 新增字段 ok
	// jsons := `[{
	// 	"path": "nodeList.1.slices.0.c",
	// 	"value": "x",
	// 	"action": "",
	// 	"nodeId": "1"
	// },{
	// 	"path": "nodeList.1.slices.1.c",
	// 	"value": "y",
	// 	"action": "",
	// 	"nodeId": "1"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段7 操作数组 删除字段 ok
	// jsons := `[{
	// 	"path": "nodeList.1.slices.0.c",
	// 	"value": "",
	// 	"action": "d",
	// 	"nodeId": "1"
	// },{
	// 	"path": "nodeList.1.slices.1.c",
	// 	"value": "",
	// 	"action": "d",
	// 	"nodeId": "1"
	// },{
	// 	"path": "nodeList.1.slices.2.c",
	// 	"value": "",
	// 	"action": "d",
	// 	"nodeId": "1"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段8 操作数组或值类型或object类型 不同类型同时更新字段 ok
	// jsons := `[{
	// 	"path": "basedata.light.yaw",
	// 	"value": "11",
	// 	"action": "",
	// 	"nodeId": "light"
	// }, {
	// 	"path": "basedata.light.x",
	// 	"value": "22",
	// 	"action": "",
	// 	"nodeId": "light"
	// }, {
	// 	"path": "nodeList.1.level",
	// 	"value": "",
	// 	"action": "d",
	// 	"nodeId": "1"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段9 操作数组 更新&删除字段混合操作 ok
	// jsons := `[{
	// 	"path": "nodeList.1.slices.0.c",
	// 	"value": "11",
	// 	"action": "",
	// 	"nodeId": "1"
	// }, {
	// 	"path": "nodeList.1.slices.1.c",
	// 	"value": "22",
	// 	"action": "",
	// 	"nodeId": "1"
	// }, {
	// 	"path": "nodeList.1.slices.0.x",
	// 	"value": "",
	// 	"action": "d",
	// 	"nodeId": "1"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段10 操作值类型或object类型 新增字段（值为结构体）TODO 批量更新处理
	// jsons := `[{
	// 	"path": "basedata.others.x",
	// 	"value": "11",
	// 	"action": "",
	// 	"nodeId": "others"
	// }, {
	// 	"path": "basedata.others.y",
	// 	"value": "22",
	// 	"action": "",
	// 	"nodeId": "others"
	// }, {
	// 	"path": "basedata.others.z",
	// 	"value": "33",
	// 	"action": "",
	// 	"nodeId": "others"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	// //更新字段11 操作数组类型 批量更新处理
	// jsons := `[{
	// 	"path": "nodeList.1.slices.0.c",
	// 	"value": "111",
	// 	"action": "",
	// 	"nodeId": "11"
	// }, {
	// 	"path": "nodeList.1.slices.1.c",
	// 	"value": "222",
	// 	"action": "",
	// 	"nodeId": "22"
	// }]`
	// var dd []map[string]string
	// json.Unmarshal([]byte(jsons), &dd)
	// data := make([]map[string]string, 0)
	// for _, v := range dd {
	// 	data = append(data, v)
	// }
	// eid := int64(1000)
	// configId := int64(1)
	// resp, err := utils.Update(configId, data, eid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", resp)

	//更新字段12 操作数组类型 调整顺序
	jsons := `[{
		"path": "nodeList.1.slices.3",
		"value": "111",
		"action": "d",
		"nodeId": "n"
	}]`
	var dd []map[string]string
	json.Unmarshal([]byte(jsons), &dd)
	data := make([]map[string]string, 0)
	for _, v := range dd {
		data = append(data, v)
	}
	eid := int64(1000)
	configId := int64(1)
	resp, err := utils.Update(configId, data, eid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", resp)
}
