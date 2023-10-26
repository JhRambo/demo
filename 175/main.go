package main

import "strings"

func main() {
	a := CutPath("111.2.3333")
	println(a)
}

func CutPath(path string) string {
	lastIdx := strings.LastIndex(path, ".")
	println(lastIdx)
	return path[:lastIdx]
}
