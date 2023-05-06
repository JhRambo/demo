package main

import (
	"fmt"
	"strings"
)

// strings
func main() {
	a := "php-golang-vue"
	arr := strings.Split(a, "-")      //字符串转切片
	str := strings.Join(arr, "-")     //切片转字符串
	fmt.Printf("%v---%T\n", arr, arr) //[php golang vue]---[]string  切片
	fmt.Printf("%v---%T\n", str, str) //php-golang-vue---string	字符

	s := []string{"1qaz", "2wsx", "3edc"}
	st := strings.Join(s, "','")
	fmt.Println(st)
}
