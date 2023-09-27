package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	err := doSomething()
	if err != nil {
		log.Printf("Error occurred: %v\n", err)
		printStackTrace()
	}
}

func doSomething() error {
	return fmt.Errorf("something went wrong")
}

func printStackTrace() {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	fmt.Println("Stack Trace:")
	fmt.Printf("%s\n", buf[:n])
}
