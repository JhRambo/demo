package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// 设置 MQTT 服务器地址
	server := "tcp://127.0.0.1:1883"
	// 设置客户端ID
	clientID := "subscriber-client"
	// 设置要订阅的主题
	topic := "mytopic"

	// 创建 MQTT 客户端连接配置
	opts := mqtt.NewClientOptions().AddBroker(server).SetClientID(clientID)

	// 创建 MQTT 客户端
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 定义消息回调处理函数
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message on topic: %s\n", msg.Topic())
		fmt.Printf("Message: %s\n", msg.Payload())
	}

	// 订阅主题，并指定消息回调函数
	token := client.Subscribe(topic, 1, messageHandler)
	token.Wait()
	if token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 等待接收消息
	time.Sleep(30 * time.Second)

	// 取消订阅主题
	token = client.Unsubscribe(topic)
	token.Wait()
	if token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 断开 MQTT 客户端连接
	client.Disconnect(250)
	fmt.Println("Disconnected")
}
