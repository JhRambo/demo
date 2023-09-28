package main

import "fmt"

func main() {
	fmt.Println(getValue())
}

func init() {
	c = "1111111111111111111"
	fmt.Printf("%v,%T", c, c)
}

var (
	c string
)

func getValue() string {
	return c
}
