package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type MsgpackRequest struct {
	MessageId string      `json:"messageId"`
	Path      string      `json:"path"`
	Data      interface{} `json:"data"`
}

type MsgpackResponse struct {
	MessageId string      `json:"messageId"`
	Path      string      `json:"path"`
	Data      interface{} `json:"data"`
}

// 更新空间资源通用结构体
type UpdateSpaceResData struct {
	Path     string `json:"path"`
	Data     string `json:"data"`
	Action   string `json:"action"`
	Id       string `json:"id"`
	TypeId   int32  `json:"typeId"`
	DataType int32  `json:"dataType"` //0.默认string 1.string 2.object或slice 3.数值类型（int float）
	Desc     string `json:"desc"`     //更新内容描述信息
}

var conn *websocket.Conn

// 连接conn
func connectWebSocket(url string) error {
	var err error
	conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	return err
}

// 发送数据到服务端
func sendMsgpackData(data []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, data)
	return err
}

// 接收服务端返回的数据
func receiveMsgpackData() {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to receive data:", err)
			return
		}
		// 解码为 MsgpackRequest 结构体
		var receivedMessage MsgpackResponse
		// err = msgpack.Unmarshal(data, &receivedMessage)
		err = json.Unmarshal(data, &receivedMessage)
		if err != nil {
			log.Println("Failed to decode message:", err)
			continue
		}
		fmt.Println("Received message:", receivedMessage)
	}
}

// 模拟ws客户端连接服务端
func main() {
	url := "ws://192.168.10.103:38902/sync/ws/space/res/get?spaceId=1&configId=1&token=111" // 替换为实际的 WebSocket URL

	err := connectWebSocket(url)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}

	// //update
	// var data = []*UpdateSpaceResData{}
	// data = append(data, &UpdateSpaceResData{
	// 	Path:     "nodeList",
	// 	Data:     "{\"id\": \"uuid2\",\"type\": 4,\"level\": 1,\"baseInfo\": {\"name\": \"基础信息\",\"description\": \"详细描述信息3\"},\"transformInfo\": {\"scale\": {\"x\": 1.1,\"y\": 1.1,\"z\": 1.1},\"position\": {\"x\": 2.2,\"y\": 2.2,\"z\": 2.2},\"rotation\": {\"x\": 3.3,\"y\": 3.3,\"z\": 3.3}},\"fileInfoList\": {},\"modelInfo\":{\"isEdit\":true,\"isTransform\":false}}",
	// 	Action:   "",
	// 	Id:       "uuid1",
	// 	TypeId:   4,
	// 	DataType: 2,
	// 	Desc:     "新增nodeList节点",
	// }, &UpdateSpaceResData{
	// 	Path:     "nodeList.baseInfo.name",
	// 	Data:     "node2节点基本信息名称",
	// 	Action:   "",
	// 	Id:       "uuid2",
	// 	TypeId:   4,
	// 	DataType: 1,
	// 	Desc:     "更新nodeList id:uuid1 node节点 string值类型",
	// })

	// //update
	// var data = []*UpdateSpaceResData{}
	// data = append(data, &UpdateSpaceResData{
	// 	Path:     "nodeList.modelInfo.isEdit",
	// 	Data:     "false",
	// 	Action:   "",
	// 	Id:       "uuid1",
	// 	TypeId:   4,
	// 	DataType: 4,
	// 	Desc:     "更新nodeList id:uuid1 node节点 bool值类型",
	// })
	// message := &MsgpackRequest{
	// 	Path:      "update",
	// 	MessageId: "111111111",
	// 	Data:      data,
	// }

	//lock
	message := &MsgpackRequest{
		Path:      "lock",
		MessageId: "111111",
		Data:      []string{"uuid1", "uuid2"},
	}

	// //unlock
	// message := &MsgpackRequest{
	// 	Path:      "unlock",
	// 	MessageId: "111111",
	// 	Data:      []string{"uuid1", "uuid2"},
	// }

	// //rollback
	// message := &MsgpackRequest{
	// 	Path:      "rollback",
	// 	MessageId: "111111",
	// 	Data:      "{\"nodeList\": [],\"baseData\": {\"light\":{\"x\":11,\"y\":22,\"Z\":33}}}",
	// }

	// //transform
	// message := &MsgpackRequest{
	// 	Path:      "transform",
	// 	MessageId: "111111",
	// 	Data:      "位置移动",
	// }

	// 转成msgpack bytes
	// msgpackData, err := msgpack.Marshal(message)
	// 转成json bytes
	msgpackData, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Failed to encode message:", err)
	}

	err = sendMsgpackData(msgpackData)
	if err != nil {
		log.Fatal("Failed to send data:", err)
	}

	// 保持长连接
	receiveMsgpackData()
}
