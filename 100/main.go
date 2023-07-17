package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

type MsgpackRequest struct {
	MessageId string `msgpack:"messageId"`
	Path      string `msgpack:"path"`
	Data      string `msgpack:"data"`
}

type MsgpackResponse struct {
	MessageId string `msgpack:"messageId"`
	Path      string `msgpack:"path"`
	Data      string `msgpack:"data"`
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
	err := conn.WriteMessage(websocket.BinaryMessage, data)
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
		err = msgpack.Unmarshal(data, &receivedMessage)
		if err != nil {
			log.Println("Failed to decode message:", err)
			continue
		}
		fmt.Println("Received message:", receivedMessage)
	}
}

// 模拟ws客户端连接服务端
func main() {
	url := "ws://192.168.10.103:38902/sync/ws/space/res/get?spaceId=1&configId=1&token=222" // 替换为实际的 WebSocket URL

	err := connectWebSocket(url)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}

	// //update
	// message := &MsgpackRequest{
	// 	Path:      "update",
	// 	MessageId: "111111111",
	// 	Data:      "[{\"path\": \"nodeList\",\"data\": \"{\\\"id\\\": \\\"uuid1\\\",\\\"type\\\": 4,\\\"level\\\": 1,\\\"baseInfo\\\": {\\\"name\\\": \\\"基础信息\\\",\\\"description\\\": \\\"详细描述信息3\\\"},\\\"transformInfo\\\": {\\\"scale\\\": {\\\"x\\\": 1.1,\\\"y\\\": 1.1,\\\"z\\\": 1.1},\\\"position\\\": {\\\"x\\\": 2.2,\\\"y\\\": 2.2,\\\"z\\\": 2.2},\\\"rotation\\\": {\\\"x\\\": 3.3,\\\"y\\\": 3.3,\\\"z\\\": 3.3}},\\\"fileInfo\\\": {}}\",\"action\": \"\",\"id\": \"uuid1\",\"typeId\": 4,\"dataType\": 2,\"desc\": \"新增nodeList节点\"},{\"path\": \"nodeList.baseInfo.name\",\"data\": \"node1节点基本信息名称\",\"action\": \"\",\"id\": \"uuid1\",\"typeId\": 4,\"dataType\": 1,\"desc\": \"更新nodeList id:uuid1 node节点 string值类型\"}]",
	// }

	// //lock
	// message := &MsgpackRequest{
	// 	Path:      "lock",
	// 	MessageId: "111111",
	// 	Data:      "[\"uuid1\",\"uuid2\"]",
	// }

	//unlock
	message := &MsgpackRequest{
		Path:      "unlock",
		MessageId: "111111",
		Data:      "[\"uuid3\",\"uuid4\"]",
	}

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

	// 转成msgpack
	msgpackData, err := msgpack.Marshal(message)
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
