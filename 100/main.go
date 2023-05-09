package main

import "fmt"

func main() {
	selectx(100)
}

func selectx(x int64) {
	switch x {
	case 100:
	case 200:
		fmt.Println("case 200")
	case 300:
		fmt.Println("case 300")
	}
}
