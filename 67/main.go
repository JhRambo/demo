package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Duration(125 * time.Second)) //2m5s
	fmt.Println(125 * time.Second)                //2m5s
}
