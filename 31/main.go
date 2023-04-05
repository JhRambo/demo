package main

import "fmt"

func main() {
	a := 'a'
	fmt.Printf("%c---%v---%T\n", a, a, a) //a---97---int32

	b := "golang"
	fmt.Printf("%c---%v---%T\n", b[0], b, b) //g---golang---string

	c := '狗'
	fmt.Printf("%c---%v---%T\n", c, c, c) //狗---29399---int32
}
