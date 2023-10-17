package main

import "fmt"

func main() {
	var json interface{}
	fmt.Println(json == nil)
	fmt.Println(json == "")
}
