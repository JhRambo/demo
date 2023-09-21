package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ""
	if !strings.Contains(str, "1234") {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}
}
