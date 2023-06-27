package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{"nodeList": [],"baseData": {}}`
	bytes, _ := json.Marshal(jsonData)
	jsonString := string(bytes)
	fmt.Println(jsonString)
}
