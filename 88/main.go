package main

import (
	"fmt"
	"io/ioutil"

	"github.com/vmihailenco/msgpack/v5"
)

// type User struct {
// 	TypeId         int32
// 	PageSize       int32
// 	PageIndex      int32
// 	EnterpriseName string
// }

type User struct {
	Name string
	Age  int32
}

func writeMsg(filename string) (err error) {
	// var user = &User{
	// 	TypeId:         0,
	// 	PageSize:       10,
	// 	PageIndex:      1,
	// 	EnterpriseName: "GGGGGGG",
	// }
	var user = &User{
		Name: "张三李四王五老六",
		Age:  int32(10),
	}
	data, err := msgpack.Marshal(user)
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
	var persons User
	data, err := ioutil.ReadFile(filename)
	fmt.Println(data)
	if err != nil {
		fmt.Printf("ReadFile failed err:%v\n", err)
		return
	}
	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		fmt.Printf("Unmarshal failed err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", persons)
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
	filename := "D:/code/demo/doc/hello-user.dat"
	writeMsg(filename)
	readMsg(filename)
}
