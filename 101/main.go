package main

import (
	"encoding/json"
	"fmt"
)

type WsRequest struct {
	Data string `json:"data"`
}

func main() {
	jsonData := `[
			{
				"path": "nodeList",
				"data": "{\"id\": \"uuid1\",\"type\": 4,\"level\": 1,\"baseInfo\": {\"name\": \"基础信息\",\"description\": \"详细描述信息3\"},\"transformInfo\": {\"scale\": {\"x\": 1.1,\"y\": 1.1,\"z\": 1.1},\"position\": {\"x\": 2.2,\"y\": 2.2,\"z\": 2.2},\"rotation\": {\"x\": 3.3,\"y\": 3.3,\"z\": 3.3}},\"fileInfo\": {}}",
				"action": "",
				"id": "uuid1",
				"typeId": 4,
				"dataType": 2,
				"desc": "新增nodeList节点"
			},
			{
				"path": "nodeList.baseInfo.name",
				"data": "node1节点基本信息名称",
				"action": "",
				"id": "uuid1",
				"typeId": 4,
				"dataType": 1,
				"desc": "更新nodeList id:uuid1 node节点 string值类型"
			}
		]`
	jsonBytes, _ := json.Marshal(jsonData)
	jsonStr := string(jsonBytes)
	fmt.Println(jsonStr)
}
