package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// jsonData := `{"others":{"x":1,"y":2}}`

	jsonData := `{
		"id": "uuid3",
		"type": 1,
		"level": 1,
		"baseInfo": {
			"name": "基础信息3",
			"description": "详细描述信息3"
		},
		"transformInfo": {
			"scale": {
				"x": 1.1,
				"y": 1.1,
				"z": 1.1
			},
			"position": {
				"x": 2.2,
				"y": 2.2,
				"z": 2.2
			},
			"rotation": {
				"x": 3.3,
				"y": 3.3,
				"z": 3.3
			}
		},
		"fileInfo": {

		}
	}`

	bytes, _ := json.Marshal(jsonData)

	jsonString := string(bytes)
	fmt.Println(jsonString)
}
