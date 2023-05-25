package main

import (
	"log"

	"github.com/golang/protobuf/proto"
)

type Person struct {
	Name    string `protobuf:"bytes,1,opt,name=name"`
	Age     int32  `protobuf:"varint,2,opt,name=age"`
	Address string `protobuf:"bytes,3,opt,name=address"`
}

// Person结构体实现protoiface.MessageV1接口
func (p *Person) Reset()         { *p = Person{} }
func (p *Person) String() string { return proto.CompactTextString(p) }
func (*Person) ProtoMessage()    {}

func main() {
	p1 := &Person{
		Name:    "Alice",
		Age:     22,
		Address: "1 Main Street",
	}
	//1.将结构体序列化为二进制格式
	data, err := proto.Marshal(p1)
	log.Println("data=========", data)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	//2.将二进制格式反序列化为结构体
	var p2 Person
	proto.Unmarshal(data, &p2)
	log.Println("person=========", &p2)
}
