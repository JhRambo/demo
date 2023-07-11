package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

type MsgpackRequest struct {
	MessageId string `json:"messageId"`
	Path      string `json:"path"`
	Data      string `json:"data"`
}

var conn *websocket.Conn

func connectWebSocket(url string) error {
	var err error
	conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	return err
}

func sendMsgpackData(data []byte) error {
	err := conn.WriteMessage(websocket.BinaryMessage, data)
	return err
}

func main() {
	url := "ws://192.168.10.103:38902/sync/ws/space/res/get?spaceId=1&configId=1&token=111" // 替换为实际的 WebSocket URL

	err := connectWebSocket(url)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}

	message := &MsgpackRequest{
		Path:      "update",
		MessageId: "111111111",
		Data:      "[{\"path\": \"nodeList\",\"data\": \"{\\\"id\\\": \\\"uuid1\\\",\\\"type\\\": 4,\\\"level\\\": 1,\\\"baseInfo\\\": {\\\"name\\\": \\\"基础信息\\\",\\\"description\\\": \\\"详细描述信息3\\\"},\\\"transformInfo\\\": {\\\"scale\\\": {\\\"x\\\": 1.1,\\\"y\\\": 1.1,\\\"z\\\": 1.1},\\\"position\\\": {\\\"x\\\": 2.2,\\\"y\\\": 2.2,\\\"z\\\": 2.2},\\\"rotation\\\": {\\\"x\\\": 3.3,\\\"y\\\": 3.3,\\\"z\\\": 3.3}},\\\"fileInfo\\\": {}}\",\"action\": \"\",\"id\": \"uuid1\",\"typeId\": 4,\"dataType\": 2,\"desc\": \"新增nodeList节点\"},{\"path\": \"nodeList.baseInfo.name\",\"data\": \"node1节点基本信息名称\",\"action\": \"\",\"id\": \"uuid1\",\"typeId\": 4,\"dataType\": 1,\"desc\": \"更新nodeList id:uuid1 node节点 string值类型\"}]",
	}

	// 转成msgpack
	msgpackData, err := msgpack.Marshal(message)
	if err != nil {
		log.Fatal("Failed to encode message:", err)
	}

	err = sendMsgpackData(msgpackData)
	if err != nil {
		log.Fatal("Failed to send data:", err)
	}

	fmt.Println("Msgpack data sent successfully.")
}
