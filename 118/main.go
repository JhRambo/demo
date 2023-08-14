package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := "[\"uuid1\",\"uuid2\"]"

	var strSlice []string
	err := json.Unmarshal([]byte(jsonData), &strSlice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(strSlice)
}
