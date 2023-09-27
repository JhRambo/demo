package main

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Panic occurred:", r)
	// 		debug.PrintStack()
	// 	}
	// }()

	panic("Something went wrong")
}
