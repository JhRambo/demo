package main

import (
	"log"

	"github.com/vmihailenco/msgpack/v5"
)

type Item struct {
	Foo string
}

// msgpack协议
func main() {
	b, err := msgpack.Marshal(&Item{Foo: "bar"}) //将结构体转化为二进制流
	if err != nil {
		panic(err)
	}
	log.Printf("%#v\n", b)

	var item Item
	err = msgpack.Unmarshal(b, &item) //将二进制流转化回结构体
	if err != nil {
		panic(err)
	}
	log.Printf("%#v\n", item)
}
