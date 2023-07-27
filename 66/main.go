package main

import (
	"errors"
	"fmt"
)

func main() {
	//自定义error内容
	e := errors.New("SigningMethodError")
	fmt.Println(e) //SigningMethodErro
}
