package main

import (
	"fmt"
	"sort"
	"time"
)

type Item struct {
	NodeType  int    `json:"nodeType"`
	EditType  int    `json:"editType"`
	NodeID    string `json:"nodeId"`
	Editor    string `json:"editor"`
	Data      string `json:"data"`
	CreateAt  string `json:"createAt"`
	VersionID string `json:"versionId"`
}

func main() {
	// 模拟给定的二维数组
	list := []Item{
		{
			NodeType:  1,
			EditType:  2,
			NodeID:    "baseInfo",
			Editor:    "GENHAOSAN03",
			Data:      `{"path":"baseData.baseInfo.name","data":"baseData.baseInfo节点基本信息名称","id":"baseInfo","typeId":1,"dataType":1,"desc":"更新内容的描述信息"}`,
			CreateAt:  "2023-07-03 10:25:00",
			VersionID: "64a2590a08980992c982b1c8",
		},
		{
			NodeType:  4,
			EditType:  2,
			NodeID:    "461762418213654528",
			Editor:    "GENHAOSAN03",
			Data:      `{"path":"nodeList.baseInfo","data":"{\"name\":\"节点461762418213654528\",\"describe\":\"描述信息1111111111111111111111111\"}","id":"461762418213654528","typeId":4,"dataType":2}`,
			CreateAt:  "2023-07-03 13:14:28",
			VersionID: "",
		},
		{
			NodeType:  0,
			EditType:  0,
			NodeID:    "",
			Editor:    "GENHAOSAN03",
			Data:      "",
			CreateAt:  "2023-07-03 13:49:54",
			VersionID: "64a2593508980992c982b1c9",
		},
	}

	// 使用自定义的排序函数
	sort.Slice(list, func(i, j int) bool {
		t1, _ := time.Parse("2006-01-02 15:04:05", list[i].CreateAt)
		t2, _ := time.Parse("2006-01-02 15:04:05", list[j].CreateAt)
		return t1.Before(t2) //正序
		return t1.After(t2)  //倒叙
	})

	// 打印排序结果
	for _, item := range list {
		fmt.Println(item.CreateAt)
	}
}
