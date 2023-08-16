package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	broker := "tcp://localhost:1883" // MQTT 代理服务器地址
	topic := "topic"                 // 消息主题

	// 创建 MQTT 客户端选项
	opts := MQTT.NewClientOptions().AddBroker(broker).SetClientID("mqtt-golang-example")

	// 创建 MQTT 客户端实例
	client := MQTT.NewClient(opts)

	// 定义消息接收处理函数
	messageHandler := func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("收到消息：%s\n", msg.Payload())
	}

	// 设置消息接收处理函数
	opts.SetDefaultPublishHandler(messageHandler)

	// 连接到 MQTT 代理服务器
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 订阅主题
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	// 发布消息
	text := "Hello, MQTT!"
	token := client.Publish(topic, 0, false, text)
	token.Wait()

	fmt.Printf("已发布消息：%s\n", text)

	// 等待中断信号（Ctrl+C）
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// 断开连接
	client.Disconnect(250)
}
