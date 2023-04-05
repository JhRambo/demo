package main

import "fmt"

func main() {
	a := []string{"php", "golang"}
	fmt.Printf("%#v---%T\n", a, a)	//[]string{"php", "golang"}---[]string

	b := make([]string, 5)
	b[1] = "php"
	fmt.Printf("%#v---%T\n", b, b)	//[]string{"", "php", "", "", ""}---[]string

	c := [5]string{"php"}
	c[3] = "golang"
	fmt.Printf("%#v---%T\n", c, c)	//[5]string{"php", "", "", "golang", ""}---[5]string
}
