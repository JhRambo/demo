package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chdir("../../Starverse/com.ghs.gateway/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Successfully changed directory!")
}
