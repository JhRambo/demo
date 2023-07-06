package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

type NewCreateSpaceResHttpRequest struct {
	Token    string `json:"token"`
	ConfigId int32  `json:"configId"`
	SpaceId  int32  `json:"spaceId"`
	Eid      int32  `json:"eid"`
	Data     string `json:"data"`
}

type MsgpackHttpRequest struct {
	Key  string `json:"key"`
	Data []byte `json:"val"`
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
	url := "ws://192.168.10.103:38902/sync/ws/space/get?spaceId=1&configId=1" // 替换为实际的 WebSocket URL

	err := connectWebSocket(url)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}
	defer conn.Close()

	req := NewCreateSpaceResHttpRequest{
		Token:    "111",
		ConfigId: 1,
		SpaceId:  1,
		Eid:      1,
		Data:     "{\"nodeList\": [],\"baseData\": {}}",
	}
	data, err := msgpack.Marshal(req)
	if err != nil {
		log.Fatal("Failed to encode message:", err)
	}

	//解析
	d := &NewCreateSpaceResHttpRequest{}
	msgpack.Unmarshal(data, d)
	fmt.Printf("%#v\n", d)

	message := MsgpackHttpRequest{
		Key:  "/v2/space/res/create",
		Data: data,
	}

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
