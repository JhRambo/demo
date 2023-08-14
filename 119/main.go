package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s = []string{"1", "2"}
	// var s = "[\"1\", \"2\"]"
	// ss := fmt.Sprintf("%v", s)
	bys, _ := json.Marshal(s)
	ss := string(bys)
	fmt.Printf("%T-%v", ss, ss)
}
