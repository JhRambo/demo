package main

import "fmt"

func main() {
	a := 1
	b := []string{"1", "2"}
	c := fmt.Sprint(a, b)
	fmt.Printf("%T,%T,%T,%v", a, b, c, c)
}
