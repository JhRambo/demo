package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{"nodeList": [],"basedata": {}}`
	bytes, _ := json.Marshal(jsonData)
	jsonString := string(bytes)
	fmt.Println(jsonString)
}
