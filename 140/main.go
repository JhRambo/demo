package main

import "fmt"

func main() {
	s := []string{"1", "2", "3", "4", "5", "6"}
	for i := 0; i < 50; i++ {
		for _, v := range s {
			fmt.Println(v)
		}
	}
}
