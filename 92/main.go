package main

import (
	"demo/92/utils"
	"log"
)

func main() {
	utils.InitRouters()
	utils.InitRouterGroup()
	utils.InitHandlers()
	log.Println("File created and updated successfully.")
}
