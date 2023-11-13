package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := []string{"{\"id\":\"1111111111111\"}", "{\"id\":\"2222222222222\"}"}
	var arr []map[string]interface{}
	var item map[string]interface{}
	for i := 0; i < len(str); i++ {
		json.Unmarshal([]byte(str[i]), &item)
		arr = append(arr, item)
	}
	bys, _ := json.Marshal(arr)
	fmt.Println(string(bys))
}
