package main

import (
	"demo/92/utils"
	"fmt"
)

func main() {
	utils.InitRouters()
	utils.InitRouterGroup()
	fmt.Println("File created and updated successfully.")
}
