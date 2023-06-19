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

	//更新字段1 操作object类型，更新单个值 ok
	// jsons := `[{
	// 	"path": "basedata.light.yaw",
	// 	"data": "111111111",
	// 	"action": "",
	// 	"id":"light",
	// 	"typeId": "light"
	// }]`

	// //更新字段2 操作object类型，更新整个结构体 ok
	// jsons := `[{
	// 	"path": "basedata.light",
	// 	"data": {
	// 		"yaw": 1111111,
	// 		"pitch": 222222
	// 	},
	// 	"action": "",
	// 	"id":"light",
	// 	"typeId": "light"
	// }]`

	// //更新字段3 操作object类型，新增结构体 ok
	// jsons := `[{
	// 	"path": "basedata.space",
	// 	"data": {
	// 		"x": 1111111,
	// 		"y": 2222222,
	// 		"z": 3333333
	// 	},
	// 	"action": "",
	// 	"id":"space",
	// 	"typeId": "space"
	// }]`

	// //更新字段4 操作object类型，删除结构体 ok
	// jsons := `[{
	// 	"path": "basedata.space",
	// 	"data": "",
	// 	"action": "d",
	// 	"id":"space",
	// 	"typeId": "space"
	// }]`

	// //更新字段5 操作object类型，删除结构体某个字段 ok
	// jsons := `[{
	// 	"path": "basedata.light.yaw",
	// 	"data": "",
	// 	"action": "d",
	// 	"id":"light",
	// 	"typeId": "light"
	// }]`

	//更新字段6 操作数组类型，更新单个值 ok
	jsons := `[{
		"path": "nodeList.baseInfo.name",
		"data": "我是哪个节点22222222222222222",
		"action": "",
		"id":"uuid2",
		"typeId": "1"
	}]`

	var dd []map[string]interface{}
	json.Unmarshal([]byte(jsons), &dd)
	data := make([]map[string]interface{}, 0)
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
