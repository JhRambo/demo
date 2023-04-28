package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int
}

func main() {
	// s := User{
	// 	Name: "张三",
	// 	Age:  10,
	// }

	// s := []string{"a", "b", "c"}

	s := map[string]interface{}{
		"name": "张三",
		"age":  10,
	}

	v := StringData(s)
	fmt.Println(v)
}

func StringData(s interface{}) string {
	u, _ := json.Marshal(s)
	return string(u)
}
