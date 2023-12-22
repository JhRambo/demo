package main

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTT 发布者示例
func main() {
	// 设置 MQTT 服务器地址
	server := "tcp://127.0.0.1:1883"
	// 设置要发布的主题
	topic := "mytopic"
	// 设置客户端ID
	clientID := "publisher-client"

	// 创建 MQTT 客户端连接配置
	opts := mqtt.NewClientOptions().AddBroker(server).SetClientID(clientID)

	// 创建 MQTT 客户端
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 生成要发布的消息
	messagePayload := "Hello, MQTT!!!!!!12345679"

	// 发布消息到指定主题
	token := client.Publish(topic, 1, true, messagePayload)
	token.Wait()
	if token.Error() != nil {
		log.Fatal(token.Error())
	}

	fmt.Println("Published message:", messagePayload)

	// 断开 MQTT 客户端连接
	client.Disconnect(250)
	fmt.Println("Disconnected")
}
