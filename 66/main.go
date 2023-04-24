package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("SigningMethodError")
	fmt.Println(e) //SigningMethodErro
}
