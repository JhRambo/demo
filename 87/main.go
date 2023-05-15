package main

import (
	"fmt"
	"io/ioutil"

	"github.com/vmihailenco/msgpack/v5"
)

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
}

func writeMsg(filename string) (err error) {
	var persons []*Person
	for i := 0; i < 20; i++ {
		p := &Person{
			Id:   i + 1,
			Name: fmt.Sprintf("msgpack_%d", i),
			Age:  i,
			Sex:  "male",
		}
		persons = append(persons, p)
	}
	data, err := msgpack.Marshal(persons)
	fmt.Printf("%#v\n", data)
	if err != nil {
		fmt.Printf("Marshal failed err:%v\n", err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Printf("WriteFile failed err:%v\n", err)
	}
	return
}

func readMsg(filename string) (err error) {
	var persons []*Person
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("ReadFile failed err:%v\n", err)
		return
	}
	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		fmt.Printf("Unmarshal failed err:%v\n", err)
		return
	}
	for _, v := range persons {
		fmt.Printf("%#v\n", v)
	}
	return
}

func main() {
	/*
	   二进制json协议
	   优点：
	       性能更快
	       更省空间
	   缺点：
	       可读性差
	   用于API通信
	*/
	filename := "D:/code/demo/doc/msgpack.txt"
	writeMsg(filename)
	readMsg(filename)
}
