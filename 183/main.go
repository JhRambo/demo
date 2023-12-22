package main

import (
	"fmt"
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源的WebSocket连接
	},
}

func main() {
	// 创建一个WebSocket处理函数
	http.HandleFunc("/ws", handleWebSocket)

	// 启动HTTP服务器
	go func() {
		log.Fatal(http.ListenAndServe(":8888", nil))
	}()

	// 连接到MQTT代理
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 订阅MQTT主题
	topic := "data"
	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 阻塞主goroutine
	select {}
}

// 处理WebSocket连接
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 在新的goroutine中处理WebSocket消息
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			// 将WebSocket消息发布到MQTT主题
			topic := "data"
			client := mqtt.NewClient(nil)
			if token := client.Connect(); token.Wait() && token.Error() != nil {
				log.Println(token.Error())
				return
			}
			token := client.Publish(topic, 0, false, message)
			token.Wait()
			if token.Error() != nil {
				log.Println(token.Error())
				return
			}
		}
	}()

	// 在新的goroutine中处理MQTT消息并发送给WebSocket客户端
	go func() {
		client := mqtt.NewClient(nil)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			return
		}

		client.Subscribe("data", 0, func(client mqtt.Client, msg mqtt.Message) {
			err := conn.WriteMessage(websocket.TextMessage, msg.Payload())
			if err != nil {
				log.Println(err)
			}
		})
	}()
}

// 处理收到的MQTT消息
func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
