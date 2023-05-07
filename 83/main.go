package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	mp := make([]map[string]string, 0)
	m1 := map[string]string{
		"devId": "111",
		"date":  "2022-01-01",
	}
	m2 := map[string]string{
		"devId": "111",
		"date":  "2022-01-02",
	}
	m3 := map[string]string{
		"devId": "222",
		"date":  "2022-02-02",
	}
	mp = append(mp, m1, m2, m3)
	ss := make(map[string][]map[string]string)
	for _, v := range mp {
		ss[v["devId"]] = append(ss[v["devId"]], v) //同一个key，append
	}
	v, _ := json.Marshal(ss)
	fmt.Println(string(v))
}
